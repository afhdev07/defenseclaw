package scanner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

type AIBOMScanner struct {
	BinaryPath string
}

func NewAIBOMScanner(binaryPath string) *AIBOMScanner {
	if binaryPath == "" {
		binaryPath = "aibom"
	}
	return &AIBOMScanner{BinaryPath: binaryPath}
}

func (s *AIBOMScanner) Name() string              { return "aibom" }
func (s *AIBOMScanner) Version() string            { return "1.0.0" }
func (s *AIBOMScanner) SupportedTargets() []string { return []string{"skill", "mcp", "code"} }

func (s *AIBOMScanner) Scan(ctx context.Context, target string) (*ScanResult, error) {
	start := time.Now()

	cmd := exec.CommandContext(ctx, s.BinaryPath, "generate", "--format", "json", target)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	duration := time.Since(start)

	result := &ScanResult{
		Scanner:   s.Name(),
		Target:    target,
		Timestamp: start,
		Duration:  duration,
	}

	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return nil, fmt.Errorf("scanner: %s not found at %q — install with: pip install aibom", s.Name(), s.BinaryPath)
		}
		if stdout.Len() == 0 {
			return nil, fmt.Errorf("scanner: %s failed: %s", s.Name(), stderr.String())
		}
	}

	if stdout.Len() > 0 {
		findings, parseErr := parseAIBOMOutput(stdout.Bytes())
		if parseErr != nil {
			return nil, fmt.Errorf("scanner: failed to parse %s output: %w", s.Name(), parseErr)
		}
		result.Findings = findings
	}

	return result, nil
}

type aibomOutput struct {
	Components []aibomComponent `json:"components"`
}

type aibomComponent struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Type     string   `json:"type"`
	Licenses []string `json:"licenses"`
}

func parseAIBOMOutput(data []byte) ([]Finding, error) {
	var out aibomOutput
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}

	findings := make([]Finding, 0, len(out.Components))
	for _, c := range out.Components {
		findings = append(findings, Finding{
			ID:       fmt.Sprintf("AIBOM-%s-%s", c.Name, c.Version),
			Severity: SeverityInfo,
			Title:    fmt.Sprintf("Component: %s@%s", c.Name, c.Version),
			Description: fmt.Sprintf("Type: %s, Licenses: %v", c.Type, c.Licenses),
			Scanner:  "aibom",
		})
	}
	return findings, nil
}
