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
	Short: "List all Entra role PIM requests current user can see",
	Long:  "List all Entra role PIM requests current user can see",
	RunE:  runListEntraRoleRequests,
}

func init() {}

func runListEntraRoleRequests(cmd *cobra.Command, _ []string) error {
	return listEntraRoleRequests(cmd.Context())
}

func listEntraRoleRequests(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleAssignmentRequests, err := graphClient.PIMEntraRoleAssignmentRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get entra role assignment requests: %w", err)
	}

	io := iostreams.System()
	err = printEntraRoleRequestList(io, time.Now(), entraRoleAssignmentRequests)
	if err != nil {
		return fmt.Errorf("failed to print entra role assignment requests: %w", err)
	}

	return nil
}

func printEntraRoleRequestList(io *iostreams.IOStreams, now time.Time, entraRoleAssignmentRequests graph.EntraRoleAssignmentRequests) error {
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
	for _, entraRoleAssignmentRequest := range entraRoleAssignmentRequests {
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
