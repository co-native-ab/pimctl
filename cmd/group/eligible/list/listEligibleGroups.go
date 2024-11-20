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
	Short: "List eligible groups you have access to",
	Long:  "List all the eligible groups you have access to in Azure PIM",
	RunE:  runListEligibleGroups,
}

func init() {}

func runListEligibleGroups(cmd *cobra.Command, _ []string) error {
	return listEligibleGroups(cmd.Context())
}

func listEligibleGroups(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	groupEligibleAssignments, err := graphClient.PIMGroupEligibleAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get eligible groups: %w", err)
	}

	io := iostreams.System()
	err = printEligibleGroupsList(io, groupEligibleAssignments)
	if err != nil {
		return fmt.Errorf("failed to print eligible groups: %w", err)
	}

	return nil
}

func printEligibleGroupsList(io *iostreams.IOStreams, groupEligibleAssignments graph.GroupEligibleAssignments) error {
	headers := []string{
		"ROLE",
		"GROUP",
		"GROUP TYPE",
		"MEMBERSHIP",
		"END TIME",
	}

	table := tableprinter.New(io, tableprinter.WithHeader(headers...))
	for _, groupEligibleAssignment := range groupEligibleAssignments {
		table.AddField(tableprinter.Title(groupEligibleAssignment.AccessID))
		table.AddField(groupEligibleAssignment.Group.DisplayName)
		table.AddField(groupEligibleAssignment.Group.GroupType())
		table.AddField(tableprinter.Title(groupEligibleAssignment.MemberType))
		table.AddField(groupEligibleAssignment.ScheduleInfo.EndTime())
		table.EndRow()
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}

	return nil
}
