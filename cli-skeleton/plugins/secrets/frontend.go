package secrets

import (
	"cli-skeleton/pkg/core"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SecretsModel struct {
	state *core.AppState
	data  []string
}

func NewSecretsModel(state *core.AppState) *SecretsModel {
	return &SecretsModel{
		state: state,
		data:  []string{"secret1", "secret2", "secret3"}, // Example data
	}
}

func (m *SecretsModel) Init() tea.Cmd {
	return nil
}

func (m *SecretsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *SecretsModel) View() string {
	style := lipgloss.NewStyle().Padding(1).Border(lipgloss.RoundedBorder())
	content := "Secrets:\n"
	for _, secret := range m.data {
		content += fmt.Sprintf("- %s\n", secret)
	}
	return style.Render(content + "\nPress 'q' to quit.")
}
