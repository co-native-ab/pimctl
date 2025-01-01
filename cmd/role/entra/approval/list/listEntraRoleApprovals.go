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
	Short: "List Entra role PIM requests waiting for approval from current user",
	Long:  "List Entra role PIM request waiting for approval from current user",
	RunE:  runListEntraRoleApprovals,
}

func init() {}

func runListEntraRoleApprovals(cmd *cobra.Command, _ []string) error {
	return listEntraRoleApprovals(cmd.Context())
}

func listEntraRoleApprovals(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleApprovalRequests, err := graphClient.PIMEntraRoleApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get entra role approval requests: %w", err)
	}

	io := iostreams.System()
	err = printEntraRoleApprovalRequestList(io, time.Now(), entraRoleApprovalRequests)
	if err != nil {
		return fmt.Errorf("failed to print approval requests: %w", err)
	}

	return nil
}

func printEntraRoleApprovalRequestList(io *iostreams.IOStreams, now time.Time, entraRoleApprovalRequests graph.EntraRoleAssignmentRequests) error {
	headers := []string{
		"ROLE",
		"REQUESTOR",
		"REQUEST TIME",
		"RESOURCE TYPE",
		"REASON",
		"TICKET NUMBER",
		"TICKET SYSTEM",
		"START TIME",
		"END TIME",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, entraRoleAssignmentRequest := range entraRoleApprovalRequests {
		table.AddField(entraRoleAssignmentRequest.RoleDefinition.DisplayName) // "ROLE"
		table.AddField(entraRoleAssignmentRequest.Principal.DisplayName)      // "REQUESTOR"
		table.AddField(entraRoleAssignmentRequest.RequestTime())              // "REQUEST TIME"
		table.AddField(entraRoleAssignmentRequest.Scope())                    // "RESOURCE TYPE"
		table.AddField(entraRoleAssignmentRequest.Justification)              // "REASON"
		table.AddField(entraRoleAssignmentRequest.TicketInfo.TicketNumber)    // "TICKET NUMBER"
		table.AddField(entraRoleAssignmentRequest.TicketInfo.TicketSystem)    // "TICKET SYSTEM"
		scheduleInfo := entraRoleAssignmentRequest.ScheduleInfo
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
