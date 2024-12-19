package menu

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/co-native-ab/pimctl/internal/graph"
)

var reviewResultChoices = []graph.ReviewResult{graph.ApproveReviewResult, graph.DenyReviewResult}

func getReviewResult(selectedGroups graph.GroupAssignmentRequests) (reviewResultModel, error) {
	p := tea.NewProgram(reviewResultModel{
		selectedGroups: selectedGroups,
	})
	reviewResultRaw, err := p.Run()
	if err != nil {
		return reviewResultModel{}, fmt.Errorf("failed to run program: %w", err)
	}

	reviewResult, ok := reviewResultRaw.(reviewResultModel)
	if !ok {
		return reviewResultModel{}, fmt.Errorf("failed to cast result to model")
	}

	return reviewResult, nil
}

type reviewResultModel struct {
	selectedGroups graph.GroupAssignmentRequests

	cursor  int
	choice  graph.ReviewResult
	success bool
	quit    bool
}

func (m reviewResultModel) Init() tea.Cmd {
	return nil
}

func (m reviewResultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quit = true
			return m, tea.Quit

		case "enter":
			m.success = true
			m.choice = reviewResultChoices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(reviewResultChoices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(reviewResultChoices) - 1
			}
		}
	}

	return m, nil
}

func (m reviewResultModel) View() string {
	if m.quit || m.success {
		return ""
	}

	s := &strings.Builder{}
	s.WriteString("Review the following requests\n")
	for _, selectedGroup := range m.selectedGroups {
		fmt.Fprintf(s, "  %s by %s\n", selectedGroup.Group.DisplayName, selectedGroup.Principal.DisplayName)
		fmt.Fprintf(s, "  > %q\n\n", selectedGroup.Justification)
	}

	for i := 0; i < len(reviewResultChoices); i++ {
		if m.cursor == i {
			s.WriteString("[x] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(reviewResultChoices[i].String())
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}
