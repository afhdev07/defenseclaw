package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show DefenseClaw status",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 4")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
