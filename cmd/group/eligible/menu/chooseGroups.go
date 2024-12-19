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

func chooseGroups(groupEligibleAssignments graph.GroupEligibleAssignments) (graph.GroupEligibleAssignments, bool, error) {
	columns := []table.Column{
		{Title: "", Width: 3},
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
			"[ ]",
			titleCase(groupEligibleAssignment.AccessID),
			groupEligibleAssignment.Group.DisplayName,
			groupEligibleAssignment.Group.GroupType(),
			titleCase(groupEligibleAssignment.MemberType),
			groupEligibleAssignment.ScheduleInfo.EndTime(),
			groupEligibleAssignment.GroupID,
		})
	}

	keyMap := table.DefaultKeyMap()
	pageDownKeys := []string{}
	for _, v := range keyMap.PageDown.Keys() {
		if v == " " {
			continue
		}
		pageDownKeys = append(pageDownKeys, v)
	}
	keyMap.PageDown.SetKeys(pageDownKeys...)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
		table.WithKeyMap(keyMap),
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

	chooseGroupResultRaw, err := tea.NewProgram(initChooseGroupModel(t, groupEligibleAssignments)).Run()
	if err != nil {
		return graph.GroupEligibleAssignments{}, false, fmt.Errorf("failed to run menu: %w", err)
	}

	chooseGroupResult, ok := chooseGroupResultRaw.(chooseGroupsModel)
	if !ok {
		return graph.GroupEligibleAssignments{}, false, fmt.Errorf("failed to cast choosen group result to model")
	}

	selectedGroups := graph.GroupEligibleAssignments{}
	for _, v := range chooseGroupResult.selected {
		selectedGroups = append(selectedGroups, v)
	}

	return selectedGroups, chooseGroupResult.quit, nil
}

func initChooseGroupModel(t table.Model, choices graph.GroupEligibleAssignments) chooseGroupsModel {
	return chooseGroupsModel{
		table:    t,
		choices:  choices,
		selected: map[int]graph.GroupEligibleAssignment{},
	}
}

type chooseGroupsModel struct {
	table    table.Model
	choices  graph.GroupEligibleAssignments
	selected map[int]graph.GroupEligibleAssignment
	success  bool
	quit     bool
}

func (m chooseGroupsModel) Init() tea.Cmd { return nil }

func (m chooseGroupsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quit = true
			return m, tea.Quit
		case " ":
			current := m.choices[m.table.Cursor()]
			if _, ok := m.selected[m.table.Cursor()]; ok {
				delete(m.selected, m.table.Cursor())
			} else {
				m.selected[m.table.Cursor()] = current
			}
		case "enter":
			current := m.choices[m.table.Cursor()]
			m.selected[m.table.Cursor()] = current
			m.success = true
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m chooseGroupsModel) View() string {
	if m.quit || m.success {
		return ""
	}

	rows := []table.Row{}
	for i, row := range m.table.Rows() {
		if _, ok := m.selected[i]; ok {
			row[0] = "[x]"
		} else {
			row[0] = "[ ]"
		}
		rows = append(rows, row)
	}
	m.table.SetRows(rows)

	return baseStyle.Render(m.table.View()) + "\n  " + m.table.HelpView() + "\n" + "  (space to select, enter to confirm, q to quit)\n"
}
