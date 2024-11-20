package menu

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu to manage group approvals others have requested",
	Long:  "Menu to manage all the group approvals others have requested in Azure PIM",
	RunE:  runGroupsApprovalMenu,
}

func init() {}

func runGroupsApprovalMenu(cmd *cobra.Command, _ []string) error {
	return groupsApprovalMenu(cmd.Context())
}

func groupsApprovalMenu(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	groupApprovalRequests, err := graphClient.PIMGroupApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get approval requests: %w", err)
	}

	chooseGroupResult, err := chooseGroup(groupApprovalRequests, time.Now())
	if err != nil {
		return fmt.Errorf("failed to choose group: %w", err)
	}

	if chooseGroupResult.quit {
		return nil
	}

	principalDisplayName := chooseGroupResult.selectedGroup.Principal.DisplayName
	requestJustification := chooseGroupResult.selectedGroup.Justification
	groupDisplayName := chooseGroupResult.selectedGroup.Group.DisplayName
	groupID := chooseGroupResult.selectedGroup.Group.ID
	approvalID := chooseGroupResult.selectedGroup.ApprovalID
	reviewResult, err := getReviewResult(principalDisplayName, groupDisplayName, requestJustification)
	if err != nil {
		return fmt.Errorf("failed to get review result: %w", err)
	}

	if reviewResult.quit {
		return nil
	}

	reviewResultChoice := reviewResult.choice
	justificationResult, err := getJustification(principalDisplayName, groupDisplayName, reviewResultChoice)
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	justification := justificationResult.textInput.Value()
	err = graphClient.PIMGroupAssignmentApprovalByApprovalID(ctx, approvalID, justification, reviewResultChoice)
	if err != nil {
		return fmt.Errorf("failed to approve or deny group assignment request: %w", err)
	}

	renderOutput(principalDisplayName, requestJustification, groupID, approvalID, groupDisplayName, justification, reviewResultChoice.String())

	return nil
}

func renderOutput(requestor, requestJustification, selectedGroupID, approvalID, selectedGroupDisplayName, justification, reviewResult string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	in := fmt.Sprintf(`# PIM Group Approval

## Requested by

  > %s

## Request Justification

  > %q

## Approval Status

| Group Name | Group ID | Approval ID | Review Result |
|---|---|---|---|
| %s | %s | %s | %s |

## Justification

  > %s
`, requestor, requestJustification, selectedGroupDisplayName, selectedGroupID, approvalID, reviewResult, justification)

	out, err := r.Render(in)
	if err != nil {
		return fmt.Errorf("failed to render output: %w", err)
	}

	fmt.Print(out)

	return nil
}
