package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List blocked or allowed items",
}

var listBlockedCmd = &cobra.Command{
	Use:   "blocked",
	Short: "List blocked skills and MCP servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 2")
	},
}

var listAllowedCmd = &cobra.Command{
	Use:   "allowed",
	Short: "List allowed skills and MCP servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 2")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listBlockedCmd)
	listCmd.AddCommand(listAllowedCmd)
}
