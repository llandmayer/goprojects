package skeleton

// import (
// 	"cli-skeleton/pkg/core"
// 	"cli-skeleton/plugins"
// 	"fmt"
// 	"strings"

// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// type Model struct {
// 	State        *core.AppState
// 	Layout       Layout
// 	Width        int
// 	Height       int
// 	MenuItems    []string
// 	Mode         string
// 	Content      string
// 	ActiveIndex  int
// 	TextInput    textinput.Model
// 	activePlugin core.Plugin
// }

// func NewModel(state *core.AppState, layout Layout, menu []string, content string) *Model {
// 	t := textinput.New()
// 	t.Placeholder = "Enter command"
// 	t.CharLimit = 50
// 	t.Width = 30

// 	return &Model{
// 		State:       state,
// 		Layout:      layout,
// 		MenuItems:   menu,
// 		Content:     content,
// 		Mode:        "menu",
// 		ActiveIndex: 0,
// 		TextInput:   t,
// 	}
// }

// func (m *Model) Init() tea.Cmd {
// 	return nil
// }

// func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.Width = msg.Width
// 		m.Height = msg.Height

// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q":
// 			return m, tea.Quit
// 		case "enter":
// 			if m.State.Current() == core.MenuMode {
// 				selected := m.MenuItems[m.ActiveIndex]
// 				plug, ok := plugins.Registry[selected]
// 				if ok {
// 					m.activePlugin = plug
// 					// Push the plugin's main mode into AppState
// 					m.State.Push(core.Mode(plug.Name() + ".main"))
// 					cmd := plug.Init()
// 					return m, cmd
// 				}
// 			}
// 		case "up", "k":
// 			if m.ActiveIndex > 0 && m.State.Current() == core.MenuMode {
// 				m.ActiveIndex--
// 			}
// 		case "down", "j":
// 			if m.ActiveIndex < len(m.MenuItems)-1 && m.State.Current() == core.MenuMode {
// 				m.ActiveIndex++
// 			}
// 		case "esc":
// 			if m.activePlugin != nil {
// 				if m.State.Current() == core.Mode(m.activePlugin.Name()+".main") {
// 					m.State.Pop()
// 					m.activePlugin = nil
// 				} else {
// 					// Delegate ESC handling to the plugin
// 					newPlugin, cmd := m.activePlugin.Update(msg, m.State)
// 					m.activePlugin = newPlugin
// 					return m, cmd
// 				}
// 			}
// 		}

// 		// Delegate other inputs to the active plugin
// 		if m.activePlugin != nil {
// 			newPlugin, cmd := m.activePlugin.Update(msg, m.State)
// 			m.activePlugin = newPlugin
// 			return m, cmd
// 		}
// 	}

// 	return m, nil
// }

// func (m *Model) View() string {
// 	headerHeight := 1
// 	footerHeight := 1
// 	usableHeight := m.Height - (headerHeight + footerHeight + 4)
// 	sidebarWidth := 20
// 	contentWidth := m.Width - sidebarWidth - 4

// 	var content string
// 	switch {
// 	case m.State.Current() == core.MenuMode:
// 		content = m.Layout.Content(contentWidth, usableHeight, m.ActiveIndex, m.Content)

// 	case m.activePlugin != nil && strings.HasPrefix(string(m.State.Current()), m.activePlugin.Name()+"."):
// 		content = m.activePlugin.View(contentWidth, usableHeight, m.State)

// 	default:
// 		content = fmt.Sprintf("Unknown mode. Current: %s. Press ESC to return to the menu.", m.State.Current())
// 	}

// 	header := m.Layout.Header(m.Width - 2)
// 	sidebar := m.Layout.Sidebar(sidebarWidth, usableHeight, m.ActiveIndex, m.MenuItems)
// 	footer := m.Layout.Footer(m.Width, footerHeight, modeString(m.State.Current()), m.TextInput)

// 	return lipgloss.JoinVertical(
// 		lipgloss.Left,
// 		header,
// 		lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content),
// 		footer,
// 	)
// }

// func modeString(m core.Mode) string {
// 	switch m {
// 	case core.MenuMode:
// 		return "menu"
// 	case core.ContentMode:
// 		return "content"
// 	default:
// 		if len(m) > 0 {
// 			parts := string(m)
// 			return parts
// 		}
// 	}
// 	return "unknown"
// }
