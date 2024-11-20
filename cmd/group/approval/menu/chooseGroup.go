package menu

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/co-native-ab/pimctl/internal/graph"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func titleCase(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func longestStringLength[T any](property string, defaultLength int, input []T) int {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return defaultLength
	}

	var data []map[string]interface{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return defaultLength
	}

	maxLength := 0
	for _, item := range data {
		result, err := jsonpath.Get(fmt.Sprintf("$.%s", property), item)
		if err != nil {
			return defaultLength
		}

		v, ok := result.(string)
		if !ok {
			return defaultLength
		}

		length := len(v)
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func chooseGroup(groupApprovalRequests graph.GroupAssignmentRequests, now time.Time) (chooseGroupModel, error) {
	columns := []table.Column{
		{Title: "Role", Width: longestStringLength("accessId", 10, groupApprovalRequests)},
		{Title: "Requestor", Width: longestStringLength("principal.displayName", 10, groupApprovalRequests)},
		{Title: "Request Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Group", Width: longestStringLength("group.displayName", 10, groupApprovalRequests)},
		{Title: "Group Type", Width: 13},
		{Title: "Reason", Width: longestStringLength("justification", 10, groupApprovalRequests)},
		{Title: "Ticket Number", Width: longestStringLength("ticketInfo.ticketNumber", 10, groupApprovalRequests)},
		{Title: "Ticket System", Width: longestStringLength("ticketInfo.ticketSystem", 10, groupApprovalRequests)},
		{Title: "Start Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "End Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Approval ID", Width: 0},
	}

	rows := []table.Row{}
	for _, groupApprovalRequest := range groupApprovalRequests {
		scheduleInfo := groupApprovalRequest.ScheduleInfo
		scheduleInfo.StartDateTime = now
		rows = append(rows, table.Row{
			titleCase(groupApprovalRequest.AccessID),
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

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	groups := map[string]graph.GroupAssignmentRequest{}
	for _, groupApprovalRequest := range groupApprovalRequests {
		groups[groupApprovalRequest.ApprovalID] = groupApprovalRequest
	}

	chooseGroupResultRaw, err := tea.NewProgram(initChooseGroupModel(t, groups)).Run()
	if err != nil {
		return chooseGroupModel{}, fmt.Errorf("failed to run menu: %w", err)
	}

	chooseGroupResult, ok := chooseGroupResultRaw.(chooseGroupModel)
	if !ok {
		return chooseGroupModel{}, fmt.Errorf("failed to cast choosen group result to model")
	}

	return chooseGroupResult, nil
}

func initChooseGroupModel(t table.Model, groups map[string]graph.GroupAssignmentRequest) chooseGroupModel {
	return chooseGroupModel{
		table:  t,
		groups: groups,
	}
}

type chooseGroupModel struct {
	table         table.Model
	groups        map[string]graph.GroupAssignmentRequest
	selectedGroup graph.GroupAssignmentRequest
	success       bool
	quit          bool
}

func (m chooseGroupModel) Init() tea.Cmd { return nil }

func (m chooseGroupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quit = true
			return m, tea.Quit
		case "enter":
			approvalID := m.table.SelectedRow()[10]
			m.selectedGroup = m.groups[approvalID]
			m.success = true
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m chooseGroupModel) View() string {
	if m.quit || m.success {
		return ""
	}

	return baseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n"
}
