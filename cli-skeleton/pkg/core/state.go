package core

type Mode string

const (
	MenuMode    Mode = "menu"
	ContentMode Mode = "content"
)

type AppState struct {
	stack []Mode
}

func NewAppState() *AppState {
	return &AppState{stack: []Mode{MenuMode}}
}

func (s *AppState) Current() Mode {
	return s.stack[len(s.stack)-1]
}

func (s *AppState) Push(m Mode) {
	s.stack = append(s.stack, m)
}

func (s *AppState) Pop() {
	if len(s.stack) > 1 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *AppState) IsPluginMode(pluginName string, mode string) bool {
	current := s.Current()
	return current == Mode(pluginName+"."+mode)
}
