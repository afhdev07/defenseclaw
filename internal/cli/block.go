package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Block a skill or MCP server",
}

var blockSkillCmd = &cobra.Command{
	Use:   "skill <name>",
	Short: "Block a skill",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 2")
	},
}

var blockMCPCmd = &cobra.Command{
	Use:   "mcp <url>",
	Short: "Block an MCP server",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented — coming in iteration 2")
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
	blockCmd.AddCommand(blockSkillCmd)
	blockCmd.AddCommand(blockMCPCmd)
}
