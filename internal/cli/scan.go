package cli

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/defenseclaw/defenseclaw/internal/scanner"
)

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan skills, MCP servers, and code for security issues",
	Long:  "Run all available scanners against the current directory or a specified path.",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runScanAll,
}

var scanSkillCmd = &cobra.Command{
	Use:   "skill <path>",
	Short: "Scan a skill directory with skill-scanner",
	Args:  cobra.ExactArgs(1),
	RunE:  runScanSkill,
}

var scanMCPCmd = &cobra.Command{
	Use:   "mcp <url>",
	Short: "Scan an MCP server with mcp-scanner",
	Args:  cobra.ExactArgs(1),
	RunE:  runScanMCP,
}

var scanAIBOMCmd = &cobra.Command{
	Use:   "aibom [path]",
	Short: "Generate AI bill of materials",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runScanAIBOM,
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.AddCommand(scanSkillCmd)
	scanCmd.AddCommand(scanMCPCmd)
	scanCmd.AddCommand(scanAIBOMCmd)
}

func runScanSkill(cmd *cobra.Command, args []string) error {
	s := scanner.NewSkillScanner(cfg.Scanners.SkillScanner)
	return execScanner(cmd.Context(), s, args[0])
}

func runScanMCP(cmd *cobra.Command, args []string) error {
	s := scanner.NewMCPScanner(cfg.Scanners.MCPScanner)
	return execScanner(cmd.Context(), s, args[0])
}

func runScanAIBOM(cmd *cobra.Command, args []string) error {
	target := "."
	if len(args) > 0 {
		target = args[0]
	}
	s := scanner.NewAIBOMScanner(cfg.Scanners.AIBOM)
	return execScanner(cmd.Context(), s, target)
}

func runScanAll(cmd *cobra.Command, args []string) error {
	target := "."
	if len(args) > 0 {
		target = args[0]
	}

	scanners := []scanner.Scanner{
		scanner.NewSkillScanner(cfg.Scanners.SkillScanner),
		scanner.NewMCPScanner(cfg.Scanners.MCPScanner),
		scanner.NewAIBOMScanner(cfg.Scanners.AIBOM),
	}

	var errs []string
	for _, s := range scanners {
		if err := execScanner(cmd.Context(), s, target); err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", s.Name(), err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("scan errors:\n  %s", strings.Join(errs, "\n  "))
	}
	return nil
}

func execScanner(ctx context.Context, s scanner.Scanner, target string) error {
	fmt.Printf("[scan] %s -> %s\n", s.Name(), target)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	result, err := s.Scan(ctx, target)
	if err != nil {
		return fmt.Errorf("scan: %w", err)
	}

	if auditLog != nil {
		if logErr := auditLog.LogScan(result); logErr != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to log scan result: %v\n", logErr)
		}
	}

	printScanResult(result)
	return nil
}

func printScanResult(r *scanner.ScanResult) {
	if r.IsClean() {
		fmt.Printf("  Clean (%s)\n", r.Duration.Round(time.Millisecond))
		return
	}

	fmt.Printf("  Findings: %d (duration: %s)\n", len(r.Findings), r.Duration.Round(time.Millisecond))
	for _, f := range r.Findings {
		fmt.Printf("  [%s] %s\n", f.Severity, f.Title)
		if f.Location != "" {
			fmt.Printf("    Location: %s\n", f.Location)
		}
		if f.Remediation != "" {
			fmt.Printf("    Fix: %s\n", f.Remediation)
		}
	}
}
