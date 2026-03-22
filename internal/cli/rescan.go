package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rescanCmd = &cobra.Command{
	Use:   "rescan",
	Short: "Re-scan all installed skills and MCP servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 4")
	},
}

func init() {
	rootCmd.AddCommand(rescanCmd)
}
