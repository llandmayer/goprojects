// secrets/frontend.go

package secrets

import (
	"cli-skeleton/pkg/core"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// -----------------------------------------------------------------------------
// 1) The actual Bubble Tea model
// -----------------------------------------------------------------------------

type SecretsModel struct {
	width  int
	height int
	data   []string
}

func NewSecretsModel() *SecretsModel {
	return &SecretsModel{
		data: []string{"secret1", "secret2", "secret3"},
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
	style := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Padding(1).
		Border(lipgloss.RoundedBorder())

	content := "Secrets:\n"
	for _, secret := range m.data {
		content += fmt.Sprintf("- %s\n", secret)
	}
	return style.Render(content + "\nPress 'q' to quit.")
}

// -----------------------------------------------------------------------------
// 2) The plugin wrapper that implements core.FrontendPlugin
// -----------------------------------------------------------------------------

type SecretsFrontendPlugin struct {
	model *SecretsModel
}

func NewSecretsFrontendPlugin() *SecretsFrontendPlugin {
	return &SecretsFrontendPlugin{
		model: NewSecretsModel(),
	}
}

func (p *SecretsFrontendPlugin) Name() string {
	return "secrets"
}

func (p *SecretsFrontendPlugin) InitTUI(state *core.AppState) tea.Model {
	return p.model
}

func (p *SecretsFrontendPlugin) Init() tea.Cmd {
	return nil
}

func (p *SecretsFrontendPlugin) Update(msg tea.Msg, state *core.AppState) (core.FrontendPlugin, tea.Cmd) {
	newModel, cmd := p.model.Update(msg)
	if updated, ok := newModel.(*SecretsModel); ok {
		p.model = updated
	}
	return p, cmd
}

func (p *SecretsFrontendPlugin) View(width, height int, state *core.AppState) string {
	p.model.width = width
	p.model.height = height
	return p.model.View()
}
