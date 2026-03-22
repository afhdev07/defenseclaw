package unit

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/defenseclaw/defenseclaw/internal/audit"
)

func TestStoreInitAndLog(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	store, err := audit.NewStore(dbPath)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}
	defer store.Close()
	defer os.Remove(dbPath)

	if err := store.Init(); err != nil {
		t.Fatalf("Init: %v", err)
	}

	err = store.LogEvent(audit.Event{
		Action:   "test",
		Target:   "target",
		Details:  "test event",
		Severity: "INFO",
	})
	if err != nil {
		t.Fatalf("LogEvent: %v", err)
	}

	events, err := store.ListEvents(10)
	if err != nil {
		t.Fatalf("ListEvents: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	if events[0].Action != "test" {
		t.Errorf("expected action 'test', got %q", events[0].Action)
	}
}

func TestStoreInsertScanResult(t *testing.T) {
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")

	store, err := audit.NewStore(dbPath)
	if err != nil {
		t.Fatalf("NewStore: %v", err)
	}
	defer store.Close()

	if err := store.Init(); err != nil {
		t.Fatalf("Init: %v", err)
	}

	err = store.InsertScanResult(
		"scan-001", "skill-scanner", "/path/to/skill",
		time.Now(), 1500, 2, "HIGH", `{"scanner":"skill-scanner"}`,
	)
	if err != nil {
		t.Fatalf("InsertScanResult: %v", err)
	}
}
