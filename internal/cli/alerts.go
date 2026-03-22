package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var alertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "View and manage security alerts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 3")
	},
}

func init() {
	rootCmd.AddCommand(alertsCmd)
}
