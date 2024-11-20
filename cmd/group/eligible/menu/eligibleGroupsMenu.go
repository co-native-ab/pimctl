package menu

import (
	"context"
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
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

	chooseGroupResult, err := chooseGroup(groupEligibleAssignments)
	if err != nil {
		return fmt.Errorf("failed to choose group: %w", err)
	}

	if chooseGroupResult.quit {
		return nil
	}

	justificationResult, err := getJustification()
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	selectedGroupID := chooseGroupResult.selectedGroupID
	selectedGroupDisplayName := chooseGroupResult.selectedGroupDisplayName
	justification := justificationResult.textInput.Value()
	status, err := cmdhelper.PIMGroupAssignmentScheduleRequest(ctx, graphClient, 0, selectedGroupID, justification)
	if err != nil {
		return fmt.Errorf("failed to schedule group assignment request: %w", err)
	}

	renderOutput(selectedGroupID, selectedGroupDisplayName, justification, status)

	return nil
}

func renderOutput(selectedGroupID, selectedGroupDisplayName, justification, status string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	in := fmt.Sprintf(`# PIM Group Assignment Request

## Request Status

| Group Name | Group ID | Status |
|---|---|---|
| %s | %s | %s |

## Justification

  > %s
`, selectedGroupDisplayName, selectedGroupID, status, justification)

	out, err := r.Render(in)
	if err != nil {
		return fmt.Errorf("failed to render output: %w", err)
	}

	fmt.Print(out)

	return nil
}
