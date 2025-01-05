package tuihelper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PaesslerAG/jsonpath"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func TitleCase(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func LongestStringLengthFn[T any](fn func(T) string, defaultLength int, input []T) int {
	maxLength := defaultLength
	for _, item := range input {
		result := fn(item)

		length := len(result)
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func LongestStringLength[T any](property string, defaultLength int, input []T) int {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return defaultLength
	}

	var data []map[string]interface{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return defaultLength
	}

	maxLength := defaultLength
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

func TableMultiChooser[T any](choices []T, inputColumns []table.Column, inputRows []table.Row) ([]T, bool, error) {
	columns := []table.Column{
		{Title: "", Width: 3},
	}
	columns = append(columns, inputColumns...)

	rows := []table.Row{}
	for _, inputRow := range inputRows {
		row := append([]string{"[ ]"}, inputRow...)
		rows = append(rows, row)
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

	rawResult, err := tea.NewProgram(initMultiChooserModel(t, choices)).Run()
	if err != nil {
		return nil, false, fmt.Errorf("failed to run menu: %w", err)
	}

	result, ok := rawResult.(multiChooserModel[T])
	if !ok {
		return nil, false, fmt.Errorf("failed to cast choosen group result to model")
	}

	selected := []T{}
	for _, v := range result.selected {
		selected = append(selected, v)
	}

	return selected, result.quit, nil
}

func initMultiChooserModel[T any](t table.Model, choices []T) multiChooserModel[T] {
	return multiChooserModel[T]{
		table:    t,
		choices:  choices,
		selected: map[int]T{},
	}
}

type multiChooserModel[T any] struct {
	table    table.Model
	choices  []T
	selected map[int]T
	success  bool
	quit     bool
}

func (m multiChooserModel[T]) Init() tea.Cmd { return nil }

func (m multiChooserModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m multiChooserModel[T]) View() string {
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
