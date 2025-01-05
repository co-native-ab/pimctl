package menu

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/glamour"
	"github.com/co-native-ab/pimctl/internal/azurerm"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/tuihelper"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu to manage Azure role approvals others have requested",
	Long:  "Menu to manage all the Azure role approvals others have requested in PIM",
	RunE:  runAzureRolesApprovalMenu,
}

func init() {}

func runAzureRolesApprovalMenu(cmd *cobra.Command, _ []string) error {
	return azureRolesApprovalMenu(cmd.Context())
}

func azureRolesApprovalMenu(ctx context.Context) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	azureRoleApprovalRequests, err := azurermClient.PIMAzureRoleApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get approval requests: %w", err)
	}

	chooseAzureRolesResult, quit, err := chooseAzureRoles(azureRoleApprovalRequests)
	if err != nil {
		return fmt.Errorf("failed to choose group: %w", err)
	}

	if quit {
		return nil
	}

	reviewResult, err := getReviewResult(chooseAzureRolesResult)
	if err != nil {
		return fmt.Errorf("failed to get review result: %w", err)
	}

	if reviewResult.quit {
		return nil
	}

	reviewResultChoice := reviewResult.choice
	justificationResult, err := getJustification(chooseAzureRolesResult, reviewResultChoice)
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	justification := justificationResult.textInput.Value()
	for _, azureRoleApprovalRequest := range chooseAzureRolesResult {
		err = azurermClient.PIMAzureRoleAssignmentApprovalByApprovalID(ctx, *azureRoleApprovalRequest.Properties.ApprovalID, justification, reviewResultChoice)
		if err != nil {
			return fmt.Errorf("failed to approve or deny group assignment request: %w", err)
		}
	}

	renderOutput(chooseAzureRolesResult, justification, reviewResultChoice.String())

	return nil
}

func chooseAzureRoles(azureRoleApprovalRequests azurerm.AzureRoleAssignmentRequests) (azurerm.AzureRoleAssignmentRequests, bool, error) {
	columns := []table.Column{
		{Title: "Role", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.Role() }, 10, azureRoleApprovalRequests)},
		{Title: "Resource", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.Resource() }, 10, azureRoleApprovalRequests)},
		{Title: "Resource Type", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.ResourceType() }, 10, azureRoleApprovalRequests)},
		{Title: "Requestor", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.Requestor() }, 10, azureRoleApprovalRequests)},
		{Title: "Request Time", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.RequestTime() }, 10, azureRoleApprovalRequests)},
		{Title: "Reason", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.Reason() }, 10, azureRoleApprovalRequests)},
		{Title: "Condition", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.Condition() }, 10, azureRoleApprovalRequests)},
		{Title: "Request Type", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.RequestType() }, 10, azureRoleApprovalRequests)},
		{Title: "Start Time", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.StartTime() }, 10, azureRoleApprovalRequests)},
		{Title: "End Time", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.EndTime() }, 10, azureRoleApprovalRequests)},
		{Title: "Scope", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleAssignmentRequest) string { return v.Scope() }, 10, azureRoleApprovalRequests)},
	}

	rows := []table.Row{}
	for _, azureRoleApprovalRequest := range azureRoleApprovalRequests {
		rows = append(rows, table.Row{
			azureRoleApprovalRequest.Role(),
			azureRoleApprovalRequest.Resource(),
			azureRoleApprovalRequest.ResourceType(),
			azureRoleApprovalRequest.Requestor(),
			azureRoleApprovalRequest.RequestTime(),
			azureRoleApprovalRequest.Reason(),
			azureRoleApprovalRequest.Condition(),
			azureRoleApprovalRequest.RequestType(),
			azureRoleApprovalRequest.StartTime(),
			azureRoleApprovalRequest.EndTime(),
			azureRoleApprovalRequest.Scope(),
		})
	}

	return tuihelper.TableMultiChooser(azureRoleApprovalRequests, columns, rows)
}

func renderOutput(selectedAzureRoles azurerm.AzureRoleAssignmentRequests, justification, reviewResult string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	s := &strings.Builder{}
	s.WriteString("# PIM Azure Role Approval\n\n")
	s.WriteString("## Request Status\n\n")
	fmt.Fprintf(s, "Review Result: %s\n\n", reviewResult)
	s.WriteString("| Requested By | Role | Justification | Scope |\n")
	s.WriteString("|---|---|---|---|\n")
	for _, azureRoleApprovalRequest := range selectedAzureRoles {
		fmt.Fprintf(s, "| %s | %s | %s | %s |\n", azureRoleApprovalRequest.Requestor(), azureRoleApprovalRequest.Role(), azureRoleApprovalRequest.Justification(), azureRoleApprovalRequest.Scope())
	}
	s.WriteString("\n## Justification\n\n")
	fmt.Fprintf(s, "  > %s\n", justification)

	out, err := r.Render(s.String())
	if err != nil {
		return fmt.Errorf("failed to render output: %w", err)
	}

	fmt.Print(out)

	return nil
}
