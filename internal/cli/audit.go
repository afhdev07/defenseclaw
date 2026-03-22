package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "View audit log and generate reports",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 5")
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
}
