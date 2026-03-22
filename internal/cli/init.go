package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/defenseclaw/defenseclaw/internal/audit"
	"github.com/defenseclaw/defenseclaw/internal/config"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize DefenseClaw environment",
	Long:  "Creates ~/.defenseclaw/, default config, SQLite database, and checks for scanner dependencies.",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(_ *cobra.Command, _ []string) error {
	env := config.DetectEnvironment()
	fmt.Printf("  Environment: %s\n", env)

	defaults := config.DefaultConfig()

	dirs := []string{defaults.DataDir, defaults.QuarantineDir, defaults.PluginDir, defaults.PolicyDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return fmt.Errorf("init: create %s: %w", dir, err)
		}
	}
	fmt.Println("  Directories: created")

	if err := defaults.Save(); err != nil {
		return fmt.Errorf("init: write config: %w", err)
	}
	fmt.Printf("  Config: %s\n", config.ConfigPath())

	store, err := audit.NewStore(defaults.AuditDB)
	if err != nil {
		return fmt.Errorf("init: create audit db: %w", err)
	}
	defer store.Close()

	if err := store.Init(); err != nil {
		return fmt.Errorf("init: initialize schema: %w", err)
	}
	fmt.Printf("  Audit DB: %s\n", defaults.AuditDB)

	logger := audit.NewLogger(store)
	_ = logger.LogAction("init", defaults.DataDir, fmt.Sprintf("environment=%s", env))

	scanners := []struct {
		name    string
		bin     string
		install string
	}{
		{"skill-scanner", defaults.Scanners.SkillScanner, "pip install cisco-ai-skill-scanner"},
		{"mcp-scanner", defaults.Scanners.MCPScanner, "pip install mcp-scanner"},
		{"aibom", defaults.Scanners.AIBOM, "pip install aibom"},
	}

	for _, s := range scanners {
		if _, err := exec.LookPath(s.bin); err != nil {
			fmt.Printf("  Scanner %s: not found (install with: %s)\n", s.name, s.install)
		} else {
			fmt.Printf("  Scanner %s: found\n", s.name)
		}
	}

	if _, err := exec.LookPath(defaults.OpenShell.Binary); err != nil {
		switch env {
		case config.EnvMacOS:
			fmt.Println("  OpenShell: not available on macOS (sandbox enforcement will be skipped)")
		default:
			fmt.Println("  OpenShell: not found (sandbox enforcement will not be active)")
		}
	} else {
		fmt.Println("  OpenShell: found")
	}

	fmt.Println("\nDefenseClaw initialized. Run 'defenseclaw scan' to start scanning.")
	return nil
}
