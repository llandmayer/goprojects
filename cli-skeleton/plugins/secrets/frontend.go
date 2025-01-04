package secrets

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins/secrets/models"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var SecretsBorderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62")).
	Padding(0, 1)

var SecretsHighlightStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Bold(true)

type SecretsModel struct {
	width        int
	height       int
	data         []models.Component
	filtered     []models.Component
	index        int
	footer       string
	filterInput  textinput.Model
	filterActive bool
}

func NewSecretsModel() *SecretsModel {
	data, err := handler.ListSecrets()
	if err != nil {
		data = []models.Component{}
	}
	filterInput := textinput.New()
	filterInput.Placeholder = "Filter..."
	filterInput.CharLimit = 156
	filterInput.Focus()
	filterInput.Width = 20
	m := &SecretsModel{
		data:         data,
		filtered:     data,
		filterInput:  filterInput,
		filterActive: false,
		index:        0,
		footer:       "Press 1 to open something",
	}
	return m
}

func (m *SecretsModel) Init() tea.Cmd {
	return nil
}

func (m *SecretsModel) filterData() {
	if m.filterInput.Value() == "" {
		m.filtered = m.data
		return
	}
	val := strings.ToLower(m.filterInput.Value())
	var out []models.Component
	for _, s := range m.data {
		if strings.Contains(strings.ToLower(s.Account), val) ||
			strings.Contains(strings.ToLower(s.Region), val) ||
			strings.Contains(strings.ToLower(s.Environment), val) ||
			strings.Contains(strings.ToLower(s.ComponentName), val) {
			out = append(out, s)
		}
	}
	m.filtered = out
	if m.index >= len(m.filtered) && len(m.filtered) > 0 {
		m.index = len(m.filtered) - 1
	}
}

func (m *SecretsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "/":
			m.filterActive = !m.filterActive
			if !m.filterActive {
				m.filterInput.SetValue("")
				m.filtered = m.data
			}
		case "enter":
			if m.filterActive {
				m.filterData()
			}
		case "q":
			if !m.filterActive {
				return m, tea.Quit
			}
		case "up", "k":
			if !m.filterActive && m.index > 0 {
				m.index--
			}
		case "down", "j":
			if !m.filterActive && m.index < len(m.filtered)-1 {
				m.index++
			}
		case "1":
			if !m.filterActive && len(m.filtered) > 0 {
				secret := m.filtered[m.index]
				m.footer = fmt.Sprintf("Opened secret: %s", secret.ComponentName)
			}
		case "esc":
			if m.filterActive {
				m.filterActive = false
				m.filterInput.SetValue("")
				m.filtered = m.data
			} else {
				m.footer = ""
			}
		}
	}
	if m.filterActive {
		m.filterInput, cmd = m.filterInput.Update(msg)
	}
	return m, cmd
}

func (m *SecretsModel) renderRow(rowIndex int) string {
	s := m.filtered[rowIndex]
	line := fmt.Sprintf(
		"%2d  %-30s %-20s %-20s %-20s",
		rowIndex+1,
		s.Account,
		s.Region,
		s.Environment,
		s.ComponentName,
	)
	if rowIndex == m.index {
		return SecretsHighlightStyle.Render("> " + line)
	}
	return "  " + line
}

func (m *SecretsModel) View() string {
	var lines []string
	header := fmt.Sprintf("  %2s  %-30s %-20s %-20s %-20s", "ID", "ACCOUNT", "REGION", "ENVIRONMENT", "SYSTEM")
	lines = append(lines, header)
	lines = append(lines, strings.Repeat("-", len(header)))
	for i := range m.filtered {
		lines = append(lines, m.renderRow(i))
	}
	if m.filterActive {
		lines = append(lines, "", "Filter: "+m.filterInput.View())
	} else {
		lines = append(lines, "", m.footer, "Press '/' to filter, 'q' to quit, up/down to navigate, '1' to open.")
	}
	content := lipgloss.JoinVertical(lipgloss.Left, lines...)
	return SecretsBorderStyle.
		Width(m.width).
		Height(m.height).
		Render(content)
}

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
