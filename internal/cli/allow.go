package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var allowCmd = &cobra.Command{
	Use:   "allow",
	Short: "Allow a previously blocked skill or MCP server",
}

var allowSkillCmd = &cobra.Command{
	Use:   "skill <name>",
	Short: "Allow a skill",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 2")
	},
}

var allowMCPCmd = &cobra.Command{
	Use:   "mcp <url>",
	Short: "Allow an MCP server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 2")
	},
}

func init() {
	rootCmd.AddCommand(allowCmd)
	allowCmd.AddCommand(allowSkillCmd)
	allowCmd.AddCommand(allowMCPCmd)
}
