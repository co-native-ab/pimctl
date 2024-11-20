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
	Short: "List active groups you have access to",
	Long:  "List all the active groups you have access to in Azure PIM",
	RunE:  runListActiveGroups,
}

func init() {}

func runListActiveGroups(cmd *cobra.Command, _ []string) error {
	return listActiveGroups(cmd.Context())
}

func listActiveGroups(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	groupActiveAssignments, err := graphClient.PIMGroupActiveAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get active groups: %w", err)
	}

	io := iostreams.System()
	err = printActiveGroupsList(io, groupActiveAssignments)
	if err != nil {
		return fmt.Errorf("failed to print eligible groups: %w", err)
	}

	return nil
}

func printActiveGroupsList(io *iostreams.IOStreams, groupActiveAssignments graph.GroupActiveAssignments) error {
	headers := []string{
		"ROLE",
		"GROUP",
		"GROUP TYPE",
		"MEMBERSHIP",
		"STATE",
		"END TIME",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, groupActiveAssignment := range groupActiveAssignments {
		table.AddField(tableprinter.Title(groupActiveAssignment.AcessID))
		table.AddField(groupActiveAssignment.Group.DisplayName)
		table.AddField(groupActiveAssignment.Group.GroupType())
		table.AddField(tableprinter.Title(groupActiveAssignment.MemberType))
		table.AddField(tableprinter.Title(groupActiveAssignment.AssignmentType))
		table.AddField(groupActiveAssignment.EndTime())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
