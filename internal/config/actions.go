package config

import (
	"fmt"
	"strings"
)

func DefaultSkillActions() SkillActionsConfig {
	return SkillActionsConfig{
		Critical: SeverityAction{File: FileActionQuarantine, Runtime: RuntimeDisable, Install: InstallBlock},
		High:     SeverityAction{File: FileActionQuarantine, Runtime: RuntimeDisable, Install: InstallBlock},
		Medium:   SeverityAction{File: FileActionNone, Runtime: RuntimeEnable, Install: InstallNone},
		Low:      SeverityAction{File: FileActionNone, Runtime: RuntimeEnable, Install: InstallNone},
		Info:     SeverityAction{File: FileActionNone, Runtime: RuntimeEnable, Install: InstallNone},
	}
}

// ForSeverity returns the configured action for a given severity string.
// Severity is matched case-insensitively; unknown values fall back to the Info action.
func (a *SkillActionsConfig) ForSeverity(severity string) SeverityAction {
	switch strings.ToUpper(severity) {
	case "CRITICAL":
		return a.Critical
	case "HIGH":
		return a.High
	case "MEDIUM":
		return a.Medium
	case "LOW":
		return a.Low
	default:
		return a.Info
	}
}

// ShouldDisable returns true if the runtime action for the given severity is "disable".
func (a *SkillActionsConfig) ShouldDisable(severity string) bool {
	return a.ForSeverity(severity).Runtime == RuntimeDisable
}

// ShouldQuarantine returns true if the file action for the given severity is "quarantine".
func (a *SkillActionsConfig) ShouldQuarantine(severity string) bool {
	return a.ForSeverity(severity).File == FileActionQuarantine
}

// ShouldInstallBlock returns true if the install action for the given severity is "block".
func (a *SkillActionsConfig) ShouldInstallBlock(severity string) bool {
	return a.ForSeverity(severity).Install == InstallBlock
}

func (a *SkillActionsConfig) Validate() error {
	entries := []struct {
		label  string
		action SeverityAction
	}{
		{"critical", a.Critical},
		{"high", a.High},
		{"medium", a.Medium},
		{"low", a.Low},
		{"info", a.Info},
	}

	for _, e := range entries {
		switch e.action.Runtime {
		case RuntimeDisable, RuntimeEnable:
		default:
			return fmt.Errorf("config: skill_actions.%s.runtime: invalid value %q (must be %q or %q)",
				e.label, e.action.Runtime, RuntimeDisable, RuntimeEnable)
		}
		switch e.action.File {
		case FileActionNone, FileActionQuarantine:
		default:
			return fmt.Errorf("config: skill_actions.%s.file: invalid value %q (must be %q or %q)",
				e.label, e.action.File, FileActionNone, FileActionQuarantine)
		}
		switch e.action.Install {
		case InstallBlock, InstallAllow, InstallNone:
		default:
			return fmt.Errorf("config: skill_actions.%s.install: invalid value %q (must be %q, %q, or %q)",
				e.label, e.action.Install, InstallBlock, InstallAllow, InstallNone)
		}
	}
	return nil
}
