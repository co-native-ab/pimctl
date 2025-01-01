package menu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/co-native-ab/pimctl/internal/graph"
)

func getJustification(selectedEntraRoles graph.EntraRoleEligibleAssignments) (justificationModel, error) {
	p := tea.NewProgram(initialRequestEligibleEntraRoleModel(selectedEntraRoles))
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
	textInput          textinput.Model
	selectedEntraRoles graph.EntraRoleEligibleAssignments
	err                error
	success            bool
	quit               bool
}

func initialRequestEligibleEntraRoleModel(selectedEntraRoles graph.EntraRoleEligibleAssignments) justificationModel {
	ti := textinput.New()
	ti.Placeholder = "Justification"
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 64

	return justificationModel{
		textInput:          ti,
		selectedEntraRoles: selectedEntraRoles,
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

	builder := &strings.Builder{}
	builder.WriteString("Justification for request(s) for the following Entra role(s):\n")
	for _, group := range m.selectedEntraRoles {
		builder.WriteString(fmt.Sprintf("  %s\n", group.RoleDefinition.DisplayName))
	}
	builder.WriteString("\n")
	builder.WriteString(m.textInput.View())
	builder.WriteString("\n\n")
	builder.WriteString("(esc to quit)")
	builder.WriteString("\n")

	return builder.String()
}
