package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var quarantineCmd = &cobra.Command{
	Use:   "quarantine <skill>",
	Short: "Immediately quarantine a skill",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 2")
	},
}

func init() {
	rootCmd.AddCommand(quarantineCmd)
}
