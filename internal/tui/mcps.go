package tui

type MCPsPanel struct{}

func NewMCPsPanel() MCPsPanel { return MCPsPanel{} }
func (p MCPsPanel) View() string { return "MCP servers panel — iteration 3" }
