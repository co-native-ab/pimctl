package menu

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/glamour"
	"github.com/co-native-ab/pimctl/internal/cmdhelper"
	"github.com/co-native-ab/pimctl/internal/graph"
	"github.com/co-native-ab/pimctl/internal/tuihelper"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Cmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu to manage eligible Entra roles you have access to",
	Long:  "Menu to manage all the eligible Entra roles you have access to in PIM",
	RunE:  runEligibleEntraRolesMenu,
}

func init() {
	Cmd.Flags().String("scope-id", "/", "The scope id to request the role assignment for")
}

type eligibleEntraRolesMenuOptions struct {
	entraRoleScopeID string
}

func newEligibleEntraRolesMenuOptions(flags *pflag.FlagSet) (*eligibleEntraRolesMenuOptions, error) {
	entraRoleScopeID, err := flags.GetString("scope-id")
	if err != nil {
		return nil, fmt.Errorf("failed to get scope-id: %w", err)
	}

	return &eligibleEntraRolesMenuOptions{
		entraRoleScopeID: entraRoleScopeID,
	}, nil
}

func runEligibleEntraRolesMenu(cmd *cobra.Command, _ []string) error {
	opts, err := newEligibleEntraRolesMenuOptions(cmd.Flags())
	if err != nil {
		return fmt.Errorf("failed to get options: %w", err)
	}

	return eligibleEntraRolesMenu(cmd.Context(), opts)
}

func eligibleEntraRolesMenu(ctx context.Context, opts *eligibleEntraRolesMenuOptions) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleEligibleAssignments, err := graphClient.PIMEntraRoleEligibleAssignments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get eligible groups: %w", err)
	}

	selectedEntraRoles, quit, err := chooseGroups(entraRoleEligibleAssignments)
	if err != nil {
		return fmt.Errorf("failed to choose group: %w", err)
	}

	if quit {
		return nil
	}

	justificationResult, err := getJustification(selectedEntraRoles)
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	justification := justificationResult.textInput.Value()
	statuses := []string{}
	for _, group := range selectedEntraRoles {
		status, err := cmdhelper.PIMEntraRoleAssignmentScheduleRequest(ctx, graphClient, 0, group.RoleDefinitionID, justification, opts.entraRoleScopeID)
		if err != nil {
			return fmt.Errorf("failed to schedule group assignment request for %s: %w", group.RoleDefinition.DisplayName, err)
		}
		statuses = append(statuses, status)
	}

	renderOutput(selectedEntraRoles, justification, statuses)

	return nil
}

func chooseGroups(entraRoleEligibleAssignments graph.EntraRoleEligibleAssignments) (graph.EntraRoleEligibleAssignments, bool, error) {
	columns := []table.Column{
		{Title: "Role", Width: tuihelper.LongestStringLength("roleDefinition.displayName", 10, entraRoleEligibleAssignments)},
		{Title: "Scope", Width: 10},
		{Title: "Membership", Width: 10},
		{Title: "End Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Role Definition ID", Width: 0},
	}

	rows := []table.Row{}
	for _, entraRoleEligibleAssignment := range entraRoleEligibleAssignments {
		rows = append(rows, table.Row{
			tuihelper.TitleCase(entraRoleEligibleAssignment.RoleDefinition.DisplayName),
			entraRoleEligibleAssignment.Scope(),
			tuihelper.TitleCase(entraRoleEligibleAssignment.MemberType),
			entraRoleEligibleAssignment.ScheduleInfo.EndTime(),
			entraRoleEligibleAssignment.RoleDefinitionID,
		})
	}

	return tuihelper.TableMultiChooser(entraRoleEligibleAssignments, columns, rows)
}

func renderOutput(selectedEntraRoles graph.EntraRoleEligibleAssignments, justification string, statuses []string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	builder := &strings.Builder{}
	builder.WriteString("# PIM Entra Role Assignment Request\n\n")
	builder.WriteString("## Request Status\n\n")
	builder.WriteString("| Role | Role Definition ID | Status |\n")
	builder.WriteString("|---|---|---|\n")
	for i, group := range selectedEntraRoles {
		builder.WriteString(fmt.Sprintf("| %s | %s | %s |\n", group.RoleDefinition.DisplayName, group.RoleDefinitionID, statuses[i]))
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
