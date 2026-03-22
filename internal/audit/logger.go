package audit

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/defenseclaw/defenseclaw/internal/scanner"
)

type Logger struct {
	store *Store
}

func NewLogger(store *Store) *Logger {
	return &Logger{store: store}
}

func (l *Logger) LogScan(result *scanner.ScanResult) error {
	scanID := uuid.New().String()
	raw, _ := result.JSON()

	if err := l.store.InsertScanResult(
		scanID, result.Scanner, result.Target, result.Timestamp,
		result.Duration.Milliseconds(), len(result.Findings),
		string(result.MaxSeverity()), string(raw),
	); err != nil {
		return err
	}

	for _, f := range result.Findings {
		tagsJSON, _ := json.Marshal(f.Tags)
		findingID := uuid.New().String()
		if err := l.store.InsertFinding(
			findingID, scanID, string(f.Severity), f.Title,
			f.Description, f.Location, f.Remediation, f.Scanner,
			string(tagsJSON),
		); err != nil {
			return err
		}
	}

	return l.store.LogEvent(Event{
		Timestamp: time.Now().UTC(),
		Action:    "scan",
		Target:    result.Target,
		Details: fmt.Sprintf("scanner=%s findings=%d max_severity=%s duration=%s",
			result.Scanner, len(result.Findings), result.MaxSeverity(), result.Duration),
		Severity: string(result.MaxSeverity()),
	})
}

func (l *Logger) LogAction(action, target, details string) error {
	return l.store.LogEvent(Event{
		Timestamp: time.Now().UTC(),
		Action:    action,
		Target:    target,
		Details:   details,
		Severity:  "INFO",
	})
}
