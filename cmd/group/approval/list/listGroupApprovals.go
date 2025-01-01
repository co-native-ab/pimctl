package list

import (
	"context"
	"fmt"
	"time"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/graph"
	"github.com/co-native-ab/pimctl/internal/tableprinter"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List Azure PIM group request waiting for approval from current user",
	Long:  "List Azure PIM group request waiting for approval from current user",
	RunE:  runListGroupApprovals,
}

func init() {}

func runListGroupApprovals(cmd *cobra.Command, _ []string) error {
	return listGroupApprovals(cmd.Context())
}

func listGroupApprovals(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	groupApprovalRequests, err := graphClient.PIMGroupApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get approval requests: %w", err)
	}

	io := iostreams.System()
	err = printGroupApprovalRequestList(io, time.Now(), groupApprovalRequests)
	if err != nil {
		return fmt.Errorf("failed to print approval requests: %w", err)
	}

	return nil
}

func printGroupApprovalRequestList(io *iostreams.IOStreams, now time.Time, groupApprovalRequests graph.GroupAssignmentRequests) error {
	headers := []string{
		"ROLE",
		"REQUESTOR",
		"REQUEST TIME",
		"RESOURCE",
		"RESOURCE TYPE",
		"REASON",
		"TICKET NUMBER",
		"TICKET SYSTEM",
		"START TIME",
		"END TIME",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, groupAssignmentRequest := range groupApprovalRequests {
		table.AddField(tableprinter.Title(groupAssignmentRequest.AccessID)) // "ROLE"
		table.AddField(groupAssignmentRequest.Principal.DisplayName)        // "REQUESTOR"
		table.AddField(groupAssignmentRequest.RequestTime())                // "REQUEST TIME"
		table.AddField(groupAssignmentRequest.Group.DisplayName)            // "RESOURCE"
		table.AddField(groupAssignmentRequest.Group.GroupType())            // "RESOURCE TYPE"
		table.AddField(groupAssignmentRequest.Justification)                // "REASON"
		table.AddField(groupAssignmentRequest.TicketInfo.TicketNumber)      // "TICKET NUMBER"
		table.AddField(groupAssignmentRequest.TicketInfo.TicketSystem)      // "TICKET SYSTEM"
		scheduleInfo := groupAssignmentRequest.ScheduleInfo
		scheduleInfo.StartDateTime = now
		table.AddField(scheduleInfo.StartTime()) // "START TIME"
		table.AddField(scheduleInfo.EndTime())   // "END TIME"
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
