package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Open the interactive terminal dashboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 3")
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
