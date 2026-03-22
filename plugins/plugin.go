package plugins

import (
	"context"
	"time"
)

type Scanner interface {
	Name() string
	Version() string
	SupportedTargets() []string
	Scan(ctx context.Context, target string) (*ScanResult, error)
}

type ScanResult struct {
	Scanner   string        `json:"scanner"`
	Target    string        `json:"target"`
	Timestamp time.Time     `json:"timestamp"`
	Findings  []Finding     `json:"findings"`
	Duration  time.Duration `json:"duration"`
}

type Finding struct {
	ID          string   `json:"id"`
	Severity    string   `json:"severity"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Remediation string   `json:"remediation"`
	Scanner     string   `json:"scanner"`
	Tags        []string `json:"tags"`
}
