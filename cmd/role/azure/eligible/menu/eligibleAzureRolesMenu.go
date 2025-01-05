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
	Short: "Menu to manage eligible Azure roles you have access to",
	Long:  "Menu to manage all the eligible Azure roles you have access to in PIM",
	RunE:  runEligibleAzureRolesMenu,
}

func init() {}

func runEligibleAzureRolesMenu(cmd *cobra.Command, _ []string) error {
	return eligibleAzureRolesMenu(cmd.Context())
}

func eligibleAzureRolesMenu(ctx context.Context) error {
	azurermClient, err := cmdhelper.NewAzurermClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	azureRoleEligibleAssignments, err := azurermClient.PIMAzureRoleEligibleAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get eligible Azure roles: %w", err)
	}

	selectedAzureRoles, quit, err := chooseAzureRoles(azureRoleEligibleAssignments)
	if err != nil {
		return fmt.Errorf("failed to choose Azure role: %w", err)
	}

	if quit {
		return nil
	}

	justificationResult, err := getJustification(selectedAzureRoles)
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	justification := justificationResult.textInput.Value()
	statuses := []string{}
	for _, role := range selectedAzureRoles {
		status, err := cmdhelper.PIMAzureRoleAssignmentScheduleRequest(ctx, azurermClient, role, 0, justification)
		if err != nil {
			return fmt.Errorf("failed to schedule role assignment request for %s (%s): %w", role.Role(), role.Scope(), err)
		}
		statuses = append(statuses, status)
	}

	renderOutput(selectedAzureRoles, justification, statuses)

	return nil
}

func chooseAzureRoles(azureRoleEligibleAssignments azurerm.AzureRoleEligibleAssignments) (azurerm.AzureRoleEligibleAssignments, bool, error) {
	columns := []table.Column{
		{Title: "Role", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleEligibleAssignment) string { return v.Role() }, 10, azureRoleEligibleAssignments)},
		{Title: "Resource", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleEligibleAssignment) string { return v.Resource() }, 10, azureRoleEligibleAssignments)},
		{Title: "Resource Type", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleEligibleAssignment) string { return v.ResourceType() }, 10, azureRoleEligibleAssignments)},
		{Title: "Membership", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleEligibleAssignment) string { return v.Membership() }, 10, azureRoleEligibleAssignments)},
		{Title: "Condition", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleEligibleAssignment) string { return v.Condition() }, 10, azureRoleEligibleAssignments)},
		{Title: "End Time", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleEligibleAssignment) string { return v.EndTime() }, 10, azureRoleEligibleAssignments)},
		{Title: "Scope", Width: tuihelper.LongestStringLengthFn(func(v azurerm.AzureRoleEligibleAssignment) string { return v.Scope() }, 10, azureRoleEligibleAssignments)},
	}

	rows := []table.Row{}
	for _, azureRoleEligibleAssignment := range azureRoleEligibleAssignments {
		rows = append(rows, table.Row{
			azureRoleEligibleAssignment.Role(),
			azureRoleEligibleAssignment.Resource(),
			azureRoleEligibleAssignment.ResourceType(),
			azureRoleEligibleAssignment.Membership(),
			azureRoleEligibleAssignment.Condition(),
			azureRoleEligibleAssignment.EndTime(),
			azureRoleEligibleAssignment.Scope(),
		})
	}

	return tuihelper.TableMultiChooser(azureRoleEligibleAssignments, columns, rows)
}

func renderOutput(selectedAzureRoles azurerm.AzureRoleEligibleAssignments, justification string, statuses []string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	builder := &strings.Builder{}
	builder.WriteString("# PIM Azure Role Assignment Request\n\n")
	builder.WriteString("## Request Status\n\n")
	builder.WriteString("| Role | Resource | Scope | Status |\n")
	builder.WriteString("|---|---|---|---|\n")
	for i, role := range selectedAzureRoles {
		builder.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", role.Role(), role.Resource(), role.Scope(), statuses[i]))
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
