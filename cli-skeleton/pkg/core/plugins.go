package core

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Plugin interface {
	Name() string
	Init() tea.Cmd
	Update(msg tea.Msg, state *AppState) (Plugin, tea.Cmd)
	View(width, height int, state *AppState) string
}

var Registry = map[string]Plugin{}

func Register(name string, p Plugin) {
	Registry[name] = p
}

func GetPlugin(name string) Plugin {
	return Registry[name]
}

func ValidatePlugin(plugin any) error {
	if _, ok := plugin.(Plugin); !ok {
		return fmt.Errorf("plugin %T does not implement core.Plugin", plugin)
	}
	return nil
}

func RunPlugin(plugin Plugin) error {
	state := NewAppState()
	program := tea.NewProgram(&pluginWrapper{
		plugin: plugin,
		state:  state,
		width:  80,
		height: 24,
	}, tea.WithAltScreen())

	_, err := program.Run()
	return err
}

type pluginWrapper struct {
	plugin Plugin
	state  *AppState
	width  int
	height int
}

func (p *pluginWrapper) Init() tea.Cmd {
	return p.plugin.Init()
}

func (p *pluginWrapper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
	}

	updatedPlugin, cmd := p.plugin.Update(msg, p.state)
	p.plugin = updatedPlugin
	return p, cmd
}

func (p *pluginWrapper) View() string {
	return p.plugin.View(p.width, p.height, p.state)
}
