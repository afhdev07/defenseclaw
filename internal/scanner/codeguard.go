package scanner

import (
	"context"
	"fmt"
)

type CodeGuardScanner struct {
	BinaryPath string
}

func NewCodeGuardScanner(binaryPath string) *CodeGuardScanner {
	if binaryPath == "" {
		binaryPath = "codeguard"
	}
	return &CodeGuardScanner{BinaryPath: binaryPath}
}

func (s *CodeGuardScanner) Name() string              { return "codeguard" }
func (s *CodeGuardScanner) Version() string            { return "0.1.0" }
func (s *CodeGuardScanner) SupportedTargets() []string { return []string{"code"} }

func (s *CodeGuardScanner) Scan(_ context.Context, _ string) (*ScanResult, error) {
	return nil, fmt.Errorf("scanner: codeguard integration not yet implemented — coming in iteration 4")
}
