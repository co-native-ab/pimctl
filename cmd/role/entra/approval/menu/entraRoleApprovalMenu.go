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
)

var Cmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu to manage Entra role approvals others have requested",
	Long:  "Menu to manage all the Entra role approvals others have requested in PIM",
	RunE:  runEntraRolesApprovalMenu,
}

func init() {}

func runEntraRolesApprovalMenu(cmd *cobra.Command, _ []string) error {
	return entraRolesApprovalMenu(cmd.Context())
}

func entraRolesApprovalMenu(ctx context.Context) error {
	graphClient, err := cmdhelper.NewGraphClientWithCachedCredential()
	if err != nil {
		return fmt.Errorf("failed to create graph client with cached credential: %w", err)
	}

	entraRoleApprovalRequests, err := graphClient.PIMEntraRoleApprovalRequests(ctx)
	if err != nil {
		return fmt.Errorf("failed to get approval requests: %w", err)
	}

	chooseEntraRolesResult, quit, err := chooseEntraRoles(entraRoleApprovalRequests, time.Now())
	if err != nil {
		return fmt.Errorf("failed to choose group: %w", err)
	}

	if quit {
		return nil
	}

	reviewResult, err := getReviewResult(chooseEntraRolesResult)
	if err != nil {
		return fmt.Errorf("failed to get review result: %w", err)
	}

	if reviewResult.quit {
		return nil
	}

	reviewResultChoice := reviewResult.choice
	justificationResult, err := getJustification(chooseEntraRolesResult, reviewResultChoice)
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	justification := justificationResult.textInput.Value()
	for _, entraRoleApprovalRequest := range chooseEntraRolesResult {
		err = graphClient.PIMEntraRoleAssignmentApprovalByApprovalID(ctx, entraRoleApprovalRequest.ApprovalID, justification, reviewResultChoice)
		if err != nil {
			return fmt.Errorf("failed to approve or deny group assignment request: %w", err)
		}
	}

	renderOutput(chooseEntraRolesResult, justification, reviewResultChoice.String())

	return nil
}

func chooseEntraRoles(entraRoleApprovalRequests graph.EntraRoleAssignmentRequests, now time.Time) (graph.EntraRoleAssignmentRequests, bool, error) {
	columns := []table.Column{
		{Title: "Role", Width: tuihelper.LongestStringLength("roleDefinition.displayName", 10, entraRoleApprovalRequests)},
		{Title: "Requestor", Width: tuihelper.LongestStringLength("principal.displayName", 10, entraRoleApprovalRequests)},
		{Title: "Request Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Resource Type", Width: 10},
		{Title: "Reason", Width: tuihelper.LongestStringLength("justification", 10, entraRoleApprovalRequests)},
		{Title: "Ticket Number", Width: tuihelper.LongestStringLength("ticketInfo.ticketNumber", 10, entraRoleApprovalRequests)},
		{Title: "Ticket System", Width: tuihelper.LongestStringLength("ticketInfo.ticketSystem", 10, entraRoleApprovalRequests)},
		{Title: "Start Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "End Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Approval ID", Width: 0},
	}

	rows := []table.Row{}
	for _, entraRoleApprovalRequest := range entraRoleApprovalRequests {
		scheduleInfo := entraRoleApprovalRequest.ScheduleInfo
		scheduleInfo.StartDateTime = now
		rows = append(rows, table.Row{
			entraRoleApprovalRequest.RoleDefinition.DisplayName,
			entraRoleApprovalRequest.Principal.DisplayName,
			entraRoleApprovalRequest.RequestTime(),
			entraRoleApprovalRequest.Scope(),
			entraRoleApprovalRequest.Justification,
			entraRoleApprovalRequest.TicketInfo.TicketNumber,
			entraRoleApprovalRequest.TicketInfo.TicketSystem,
			scheduleInfo.StartTime(),
			scheduleInfo.EndTime(),
		})
	}

	return tuihelper.TableMultiChooser(entraRoleApprovalRequests, columns, rows)
}

func renderOutput(selectedEntraRoles graph.EntraRoleAssignmentRequests, justification, reviewResult string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	s := &strings.Builder{}
	s.WriteString("# PIM Entra Role Approval\n\n")
	s.WriteString("## Request Status\n\n")
	fmt.Fprintf(s, "Review Result: %s\n\n", reviewResult)
	s.WriteString("| Requested By | Role | Justification |\n")
	s.WriteString("|---|---|---|\n")
	for _, entraRoleApprovalRequest := range selectedEntraRoles {
		fmt.Fprintf(s, "| %s | %s | %s |\n", entraRoleApprovalRequest.Principal.DisplayName, entraRoleApprovalRequest.RoleDefinition.DisplayName, entraRoleApprovalRequest.Justification)
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
