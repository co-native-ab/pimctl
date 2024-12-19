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

	chooseGroupsResult, quit, err := chooseGroups(groupApprovalRequests, time.Now())
	if err != nil {
		return fmt.Errorf("failed to choose group: %w", err)
	}

	if quit {
		return nil
	}

	reviewResult, err := getReviewResult(chooseGroupsResult)
	if err != nil {
		return fmt.Errorf("failed to get review result: %w", err)
	}

	if reviewResult.quit {
		return nil
	}

	reviewResultChoice := reviewResult.choice
	justificationResult, err := getJustification(chooseGroupsResult, reviewResultChoice)
	if err != nil {
		return fmt.Errorf("failed to get justification: %w", err)
	}

	if justificationResult.quit {
		return nil
	}

	justification := justificationResult.textInput.Value()
	for _, groupApprovalRequest := range chooseGroupsResult {
		err = graphClient.PIMGroupAssignmentApprovalByApprovalID(ctx, groupApprovalRequest.ApprovalID, justification, reviewResultChoice)
		if err != nil {
			return fmt.Errorf("failed to approve or deny group assignment request: %w", err)
		}
	}

	renderOutput(chooseGroupsResult, justification, reviewResultChoice.String())

	return nil
}

func chooseGroups(groupApprovalRequests graph.GroupAssignmentRequests, now time.Time) (graph.GroupAssignmentRequests, bool, error) {
	columns := []table.Column{
		{Title: "Role", Width: tuihelper.LongestStringLength("accessId", 10, groupApprovalRequests)},
		{Title: "Requestor", Width: tuihelper.LongestStringLength("principal.displayName", 10, groupApprovalRequests)},
		{Title: "Request Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Group", Width: tuihelper.LongestStringLength("group.displayName", 10, groupApprovalRequests)},
		{Title: "Group Type", Width: 13},
		{Title: "Reason", Width: tuihelper.LongestStringLength("justification", 10, groupApprovalRequests)},
		{Title: "Ticket Number", Width: tuihelper.LongestStringLength("ticketInfo.ticketNumber", 10, groupApprovalRequests)},
		{Title: "Ticket System", Width: tuihelper.LongestStringLength("ticketInfo.ticketSystem", 10, groupApprovalRequests)},
		{Title: "Start Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "End Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Approval ID", Width: 0},
	}

	rows := []table.Row{}
	for _, groupApprovalRequest := range groupApprovalRequests {
		scheduleInfo := groupApprovalRequest.ScheduleInfo
		scheduleInfo.StartDateTime = now
		rows = append(rows, table.Row{
			tuihelper.TitleCase(groupApprovalRequest.AccessID),
			groupApprovalRequest.Principal.DisplayName,
			groupApprovalRequest.RequestTime(),
			groupApprovalRequest.Group.DisplayName,
			groupApprovalRequest.Group.GroupType(),
			groupApprovalRequest.Justification,
			groupApprovalRequest.TicketInfo.TicketNumber,
			groupApprovalRequest.TicketInfo.TicketSystem,
			scheduleInfo.StartTime(),
			scheduleInfo.EndTime(),
			groupApprovalRequest.ApprovalID,
		})
	}

	return tuihelper.TableMultiChooser(groupApprovalRequests, columns, rows)
}

func renderOutput(selectedGroups graph.GroupAssignmentRequests, justification, reviewResult string) error {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(128),
	)

	s := &strings.Builder{}
	s.WriteString("# PIM Group Approval\n\n")
	s.WriteString("## Request Status\n\n")
	fmt.Fprintf(s, "Review Result: %s\n\n", reviewResult)
	s.WriteString("| Requested By | Group Name | Justification |\n")
	s.WriteString("|---|---|---|\n")
	for _, groupApprovalRequest := range selectedGroups {
		fmt.Fprintf(s, "| %s | %s | %s |\n", groupApprovalRequest.Principal.DisplayName, groupApprovalRequest.Group.DisplayName, groupApprovalRequest.Justification)
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
