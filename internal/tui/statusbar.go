package tui

type StatusBar struct{}

func NewStatusBar() StatusBar { return StatusBar{} }
func (s StatusBar) View() string { return "Status bar — iteration 3" }
