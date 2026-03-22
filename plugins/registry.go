package plugins

import "fmt"

type Registry struct {
	scanners []Scanner
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Discover(_ string) error {
	return fmt.Errorf("plugins: discovery not yet implemented — coming in iteration 5")
}

func (r *Registry) Scanners() []Scanner {
	return r.scanners
}
