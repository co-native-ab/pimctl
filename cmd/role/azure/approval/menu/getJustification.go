package menu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/co-native-ab/pimctl/internal/azurerm"
)

func getJustification(selectedAzureRoles azurerm.AzureRoleAssignmentRequests, reviewResult azurerm.ReviewResult) (justificationModel, error) {
	p := tea.NewProgram(initialRequestEligibleAzureRoleModel(selectedAzureRoles, reviewResult))
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
	selectedAzureRoles azurerm.AzureRoleAssignmentRequests
	reviewResult       azurerm.ReviewResult

	textInput textinput.Model
	err       error
	success   bool
	quit      bool
}

func initialRequestEligibleAzureRoleModel(selectedAzureRoles azurerm.AzureRoleAssignmentRequests, reviewResult azurerm.ReviewResult) justificationModel {
	ti := textinput.New()
	ti.Placeholder = "Justification"
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 64

	return justificationModel{
		selectedAzureRoles: selectedAzureRoles,
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
	if m.reviewResult == azurerm.DenyReviewResult {
		reviewString = "denying"
	}

	fmt.Fprintf(s, "Reason for %s assignment(s):\n", reviewString)
	for _, selectedAzureRole := range m.selectedAzureRoles {
		fmt.Fprintf(s, "  %s (%s) by %s\n", selectedAzureRole.Role(), selectedAzureRole.Scope(), selectedAzureRole.Requestor())
		fmt.Fprintf(s, "  > %q\n\n", selectedAzureRole.Justification())
	}

	fmt.Fprintf(s, "%s\n\n", m.textInput.View())
	fmt.Fprintf(s, "(esc to quit)\n")

	return s.String()
}
