package terragrunt

import (
	"cli-skeleton/pkg/core"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// var _ core.Plugin = (*TerragruntPlugin)(nil)

type pluginMode string

const (
	MainMode pluginMode = "main"
	SubMode1 pluginMode = "sub1"
	SubMode2 pluginMode = "sub2"
)

type TerragruntPlugin struct {
	currentMode pluginMode
	counter     int
}

func (p *TerragruntPlugin) currentPluginMode(state *core.AppState) pluginMode {
	current := state.Current()
	switch current {
	case core.Mode(p.Name() + "." + string(MainMode)):
		return MainMode
	case core.Mode(p.Name() + "." + string(SubMode1)):
		return SubMode1
	case core.Mode(p.Name() + "." + string(SubMode2)):
		return SubMode2
	default:
		return ""
	}
}

func (p *TerragruntPlugin) Name() string {
	return "terragrunt"
}

func (p *TerragruntPlugin) Init() tea.Cmd {
	p.currentMode = MainMode
	return nil
}

func (p *TerragruntPlugin) Update(msg tea.Msg, state *core.AppState) (core.Plugin, tea.Cmd) {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "1":
			mode := core.Mode(p.Name() + "." + string(SubMode1))
			state.Push(mode)
		case "2":
			mode := core.Mode(p.Name() + "." + string(SubMode2))
			state.Push(mode)
		case "up":
			if p.currentPluginMode(state) == SubMode1 {
				p.counter++
			}
		case "down":
			if p.currentPluginMode(state) == SubMode1 {
				if p.counter > 0 {
					p.counter--
				}
			}
		case "esc":
			if p.currentPluginMode(state) != MainMode {
				state.Pop()
			} else {
				return p, func() tea.Msg {
					return core.GoBackMsg{}
				}
			}
		case "ctrl+c":
			return p, tea.Quit
		}
	}
	return p, nil
}
func (p *TerragruntPlugin) View(width, height int, state *core.AppState) string {
	style := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	switch p.currentPluginMode(state) {
	case MainMode:
		return style.Render("Terragrunt Plugin\nPress 1 for SubMode1\nPress 2 for SubMode2\nESC to exit")
	case SubMode1:
		return style.Render(fmt.Sprintf("SubMode1\nCounter: %d\nESC to go back", p.counter))
	case SubMode2:
		return style.Render("SubMode2\nESC to go back")
	default:
		return style.Render(fmt.Sprintf("Unknown Mode %s", state.Current()))
	}
}
