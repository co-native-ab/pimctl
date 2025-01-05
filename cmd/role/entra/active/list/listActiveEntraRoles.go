package list

import (
	"context"
	"fmt"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/graph"
	"github.com/co-native-ab/pimctl/internal/tableprinter"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List active Entra roles you have access to",
	Long:  "List all the active Entra roles you have access to in PIM",
	RunE:  runListActiveEntraRoles,
}

func init() {}

func runListActiveEntraRoles(cmd *cobra.Command, _ []string) error {
	return listActiveEntraRoles(cmd.Context())
}

func listActiveEntraRoles(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleActiveAssignments, err := graphClient.PIMEntraRoleActiveAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active entra roles: %w", err)
	}

	io := iostreams.System()
	err = printActiveEntraRolesList(io, entraRoleActiveAssignments)
	if err != nil {
		return fmt.Errorf("failed to print active entra roles: %w", err)
	}

	return nil
}

func printActiveEntraRolesList(io *iostreams.IOStreams, entraRoleActiveAssignments graph.EntraRoleActiveAssignments) error {
	headers := []string{
		"ROLE",
		"SCOPE",
		"MEMBERSHIP",
		"STATE",
		"END TIME",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, entraRoleActiveAssignment := range entraRoleActiveAssignments {
		table.AddField(entraRoleActiveAssignment.RoleDefinition.DisplayName)
		table.AddField(entraRoleActiveAssignment.Scope())
		table.AddField(entraRoleActiveAssignment.MemberType)
		table.AddField(entraRoleActiveAssignment.AssignmentType)
		table.AddField(entraRoleActiveAssignment.EndTime())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
