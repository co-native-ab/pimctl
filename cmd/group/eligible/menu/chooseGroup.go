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

func chooseGroup(groupEligibleAssignments graph.GroupEligibleAssignments) (chooseGroupModel, error) {
	columns := []table.Column{
		{Title: "Role", Width: longestStringLength("accessId", 10, groupEligibleAssignments)},
		{Title: "Group", Width: longestStringLength("group.displayName", 10, groupEligibleAssignments)},
		{Title: "Group Type", Width: 13},
		{Title: "Membership", Width: 10},
		{Title: "End Time", Width: len(time.Now().Local().Format(time.RFC3339))},
		{Title: "Group ID", Width: 0},
	}

	rows := []table.Row{}
	for _, groupEligibleAssignment := range groupEligibleAssignments {
		rows = append(rows, table.Row{
			titleCase(groupEligibleAssignment.AccessID),
			groupEligibleAssignment.Group.DisplayName,
			groupEligibleAssignment.Group.GroupType(),
			titleCase(groupEligibleAssignment.MemberType),
			groupEligibleAssignment.ScheduleInfo.EndTime(),
			groupEligibleAssignment.GroupID,
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

	chooseGroupResultRaw, err := tea.NewProgram(initChooseGroupModel(t)).Run()
	if err != nil {
		return chooseGroupModel{}, fmt.Errorf("failed to run menu: %w", err)
	}

	chooseGroupResult, ok := chooseGroupResultRaw.(chooseGroupModel)
	if !ok {
		return chooseGroupModel{}, fmt.Errorf("failed to cast choosen group result to model")
	}

	return chooseGroupResult, nil
}

func initChooseGroupModel(t table.Model) chooseGroupModel {
	return chooseGroupModel{
		table: t,
	}
}

type chooseGroupModel struct {
	table                    table.Model
	selectedGroupID          string
	selectedGroupDisplayName string
	success                  bool
	quit                     bool
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
			m.selectedGroupID = m.table.SelectedRow()[5]
			m.selectedGroupDisplayName = m.table.SelectedRow()[1]
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
