package tui

type SkillsPanel struct{}

func NewSkillsPanel() SkillsPanel { return SkillsPanel{} }
func (p SkillsPanel) View() string { return "Skills panel — iteration 3" }
