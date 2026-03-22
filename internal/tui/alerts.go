package tui

type AlertsPanel struct{}

func NewAlertsPanel() AlertsPanel { return AlertsPanel{} }
func (p AlertsPanel) View() string { return "Alerts panel — iteration 3" }
