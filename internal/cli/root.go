package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/defenseclaw/defenseclaw/internal/audit"
	"github.com/defenseclaw/defenseclaw/internal/config"
)

var (
	cfg        *config.Config
	auditStore *audit.Store
	auditLog   *audit.Logger
)

var rootCmd = &cobra.Command{
	Use:   "defenseclaw",
	Short: "Enterprise governance layer for OpenClaw",
	Long: `DefenseClaw secures OpenClaw deployments by scanning skills, MCP servers,
and code before they run, enforcing block/allow lists, and providing a
terminal dashboard for governance.`,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		switch cmd.Name() {
		case "init", "help", "completion":
			return nil
		}

		var err error
		cfg, err = config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config — run 'defenseclaw init' first: %w", err)
		}

		auditStore, err = audit.NewStore(cfg.AuditDB)
		if err != nil {
			return fmt.Errorf("failed to open audit store: %w", err)
		}

		auditLog = audit.NewLogger(auditStore)
		return nil
	},
	PersistentPostRun: func(_ *cobra.Command, _ []string) {
		if auditStore != nil {
			auditStore.Close()
		}
	},
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
