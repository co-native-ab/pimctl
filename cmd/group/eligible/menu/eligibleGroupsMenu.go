package menu

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/graph"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu to manage eligible groups you have access to",
	Long:  "Menu to manage all the eligible groups you have access to in Azure PIM",
	RunE:  runEligibleGroupsMenu,
}

func init() {}

func runEligibleGroupsMenu(cmd *cobra.Command, _ []string) error {
	return eligibleGroupsMenu(cmd.Context())
}

func eligibleGroupsMenu(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	groupEligibleAssignments, err := graphClient.PIMGroupEligibleAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get eligible groups: %w", err)
	}

	selectedGroups, quit, err := chooseGroups(groupEligibleAssignments)
	if err != nil {
		return fmt.Errorf("failed to choose group: %w", err)
	}

	if quit {
		return nil
	}

	justificationResult, err := getJustification(selectedGroups)
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	justification := justificationResult.textInput.Value()
	statuses := []string{}
	for _, group := range selectedGroups {
		status, err := cmdhelper.PIMGroupAssignmentScheduleRequest(ctx, graphClient, 0, group.GroupID, justification)
		if err != nil {
			return fmt.Errorf("failed to schedule group assignment request for %s: %w", group.Group.DisplayName, err)
		}
		statuses = append(statuses, status)
	}

	renderOutput(selectedGroups, justification, statuses)

	return nil
}

func renderOutput(selectedGroups graph.GroupEligibleAssignments, justification string, statuses []string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	builder := &strings.Builder{}
	builder.WriteString("# PIM Group Assignment Request\n\n")
	builder.WriteString("## Request Status\n\n")
	builder.WriteString("| Group Name | Group ID | Status |\n")
	builder.WriteString("|---|---|---|\n")
	for i, group := range selectedGroups {
		builder.WriteString(fmt.Sprintf("| %s | %s | %s |\n", group.Group.DisplayName, group.GroupID, statuses[i]))
	}
	builder.WriteString("\n## Justification\n\n")
	builder.WriteString(fmt.Sprintf("  > %s\n", justification))

	out, err := r.Render(builder.String())
	if err != nil {
		return fmt.Errorf("failed to render output: %w", err)
	}

	fmt.Print(out)

	return nil
}
