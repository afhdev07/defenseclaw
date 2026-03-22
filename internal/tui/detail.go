package tui

type DetailModal struct {
	visible bool
}

func NewDetailModal() DetailModal { return DetailModal{} }
func (d DetailModal) View() string { return "Detail modal — iteration 3" }
