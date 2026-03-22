package tui

import "github.com/charmbracelet/lipgloss"

var (
	TabStyle         = lipgloss.NewStyle().Padding(0, 2)
	ActiveTabStyle   = lipgloss.NewStyle().Padding(0, 2).Bold(true).Foreground(lipgloss.Color("39"))
	StatusStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	SeverityCritical = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	SeverityHigh     = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
	SeverityMedium   = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	SeverityLow      = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
)
