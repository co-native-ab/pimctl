package menu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/co-native-ab/pimctl/internal/graph"
)

func getJustification(selectedEntraRoles graph.EntraRoleAssignmentRequests, reviewResult graph.ReviewResult) (justificationModel, error) {
	p := tea.NewProgram(initialRequestEligibleEntraRoleModel(selectedEntraRoles, reviewResult))
	rawResult, err := p.Run()
	if err != nil {
		return justificationModel{}, fmt.Errorf("failed to run program: %w", err)
	}

	result, ok := rawResult.(justificationModel)
	if !ok {
		return justificationModel{}, fmt.Errorf("failed to cast result to model")
	}

	if result.err != nil {
		return justificationModel{}, fmt.Errorf("error: %w", result.err)
	}

	if result.textInput.Value() == "" && !result.quit {
		return justificationModel{}, fmt.Errorf("justification is required")
	}

	return result, nil
}

type (
	errMsg error
)

type justificationModel struct {
	selectedEntraRoles graph.EntraRoleAssignmentRequests
	reviewResult       graph.ReviewResult

	textInput textinput.Model
	err       error
	success   bool
	quit      bool
}

func initialRequestEligibleEntraRoleModel(selectedEntraRoles graph.EntraRoleAssignmentRequests, reviewResult graph.ReviewResult) justificationModel {
	ti := textinput.New()
	ti.Placeholder = "Justification"
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 64

	return justificationModel{
		selectedEntraRoles: selectedEntraRoles,
		reviewResult:       reviewResult,
		textInput:          ti,
		err:                nil,
	}
}

func (m justificationModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m justificationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quit = true
			return m, tea.Quit
		case tea.KeyEnter:
			m.success = true
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m justificationModel) View() string {
	if m.quit || m.success {
		return ""
	}

	s := &strings.Builder{}
	reviewString := "approving"
	if m.reviewResult == graph.DenyReviewResult {
		reviewString = "denying"
	}

	fmt.Fprintf(s, "Reason for %s assignment(s):\n", reviewString)
	for _, selectedGroup := range m.selectedEntraRoles {
		fmt.Fprintf(s, "  %s by %s\n", selectedGroup.RoleDefinition.DisplayName, selectedGroup.Principal.DisplayName)
		fmt.Fprintf(s, "  > %q\n\n", selectedGroup.Justification)
	}

	fmt.Fprintf(s, "%s\n\n", m.textInput.View())
	fmt.Fprintf(s, "(esc to quit)\n")

	return s.String()
}
