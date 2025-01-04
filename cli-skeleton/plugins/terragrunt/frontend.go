package terragrunt

import (
	"cli-skeleton/pkg/core"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// -----------------------------------------------------------------------------
// 1) The actual Bubble Tea model
// -----------------------------------------------------------------------------

type TerragruntModel struct {
	width  int
	height int
	data   []string
}

func NewTerragruntModel() *TerragruntModel {
	return &TerragruntModel{
		data: []string{"terragrunt", "terra", "secret3"},
	}
}

func (m *TerragruntModel) Init() tea.Cmd {
	return nil
}

func (m *TerragruntModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *TerragruntModel) View() string {
	style := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Padding(1).
		Border(lipgloss.RoundedBorder())
	content := "Secrets:\n"
	for _, secret := range m.data {
		content += fmt.Sprintf("- %s\n", secret)
	}
	content = lipgloss.NewStyle().Width(m.width - 4).Render(content) // Adjust content width
	return style.Render(content + "\nPress 'q' to quit.")
}

// -----------------------------------------------------------------------------
// 2) The plugin wrapper that implements core.FrontendPlugin
// -----------------------------------------------------------------------------

type TerragruntFrontendPlugin struct {
	model *TerragruntModel
}

func NewTerragruntFrontendPlugin() *TerragruntFrontendPlugin {
	return &TerragruntFrontendPlugin{
		model: NewTerragruntModel(),
	}
}

func (p *TerragruntFrontendPlugin) Name() string {
	return "terragrunt"
}

func (p *TerragruntFrontendPlugin) InitTUI(state *core.AppState) tea.Model {
	// If you need 'state', store or use it here
	return p.model
}

func (p *TerragruntFrontendPlugin) Init() tea.Cmd {
	return nil
}

func (p *TerragruntFrontendPlugin) Update(msg tea.Msg, state *core.AppState) (core.FrontendPlugin, tea.Cmd) {
	newModel, cmd := p.model.Update(msg)
	if updated, ok := newModel.(*TerragruntModel); ok {
		p.model = updated
	}
	return p, cmd
}

func (p *TerragruntFrontendPlugin) View(width, height int, state *core.AppState) string {
	p.model.width = width
	p.model.height = height
	return p.model.View()
}
