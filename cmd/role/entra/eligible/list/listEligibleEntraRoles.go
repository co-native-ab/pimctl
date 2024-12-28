package list

import (
	"context"
	"fmt"

	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/graph"
	"github.com/co-native-ab/pimctl/internal/tableprinter"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List eligible Entra roles you have access to",
	Long:  "List all the eligible Entra roles you have access to in PIM",
	RunE:  runListEligibleEntraRoles,
}

func init() {}

func runListEligibleEntraRoles(cmd *cobra.Command, _ []string) error {
	return listEligibleEntraRoles(cmd.Context())
}

func listEligibleEntraRoles(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleEligibleAssignments, err := graphClient.PIMEntraRoleEligibleAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get eligible entra roles: %w", err)
	}

	io := iostreams.System()
	err = printEligibleGroupsList(io, entraRoleEligibleAssignments)
	if err != nil {
		return fmt.Errorf("failed to print eligible groups: %w", err)
	}

	return nil
}

func printEligibleGroupsList(io *iostreams.IOStreams, entraRoleEligibleAssignments graph.EntraRoleEligibleAssignments) error {
	headers := []string{
		"ROLE",
		"SCOPE",
		"MEMBERSHIP",
		"END TIME",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, entraRoleEligibleAssignment := range entraRoleEligibleAssignments {
		table.AddField(entraRoleEligibleAssignment.RoleDefinition.DisplayName)
		table.AddField(entraRoleEligibleAssignment.Scope())
		table.AddField(entraRoleEligibleAssignment.MemberType)
		table.AddField(entraRoleEligibleAssignment.ScheduleInfo.EndTime())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
