package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy OpenClaw in a secured sandbox",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 4")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
