package unit

import (
	"testing"

	"github.com/defenseclaw/defenseclaw/internal/scanner"
)

func TestScanResultIsClean(t *testing.T) {
	r := &scanner.ScanResult{}
	if !r.IsClean() {
		t.Error("expected empty result to be clean")
	}
}

func TestScanResultMaxSeverity(t *testing.T) {
	tests := []struct {
		name     string
		findings []scanner.Finding
		want     scanner.Severity
	}{
		{"no findings", nil, scanner.SeverityInfo},
		{"single high", []scanner.Finding{{Severity: scanner.SeverityHigh}}, scanner.SeverityHigh},
		{"mixed", []scanner.Finding{
			{Severity: scanner.SeverityLow},
			{Severity: scanner.SeverityCritical},
			{Severity: scanner.SeverityMedium},
		}, scanner.SeverityCritical},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &scanner.ScanResult{Findings: tt.findings}
			if got := r.MaxSeverity(); got != tt.want {
				t.Errorf("MaxSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanResultCountBySeverity(t *testing.T) {
	r := &scanner.ScanResult{
		Findings: []scanner.Finding{
			{Severity: scanner.SeverityHigh},
			{Severity: scanner.SeverityHigh},
			{Severity: scanner.SeverityLow},
		},
	}
	if got := r.CountBySeverity(scanner.SeverityHigh); got != 2 {
		t.Errorf("CountBySeverity(HIGH) = %d, want 2", got)
	}
}
