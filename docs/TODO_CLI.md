# CLI Hierarchy TODO

New top-level command structure for DefenseClaw CLI.

## Commands

- [ ] `defenseclaw init` — Initialize DefenseClaw in current workspace
- [ ] `defenseclaw setup` — Configure DefenseClaw settings and integrations
- [ ] `defenseclaw deploy` — Deploy agent with enforcement policies
- [ ] `defenseclaw sidecar` — Run DefenseClaw as a sidecar process
- [ ] `defenseclaw skill` — Manage skills (scan, block, allow, list)
- [ ] `defenseclaw plugin` — Manage plugins (install, list, remove)
- [ ] `defenseclaw mcp` — Manage MCP servers (scan, block, allow, list)
- [ ] `defenseclaw aibom` — Generate and manage AI Bill of Materials
- [ ] `defenseclaw status` — Show current enforcement status and health
- [ ] `defenseclaw help` — Display help information
- [ ] `defenseclaw alerts` — View and manage security alerts
- [ ] `defenseclaw tui` — Launch interactive terminal UI

## Notes

- Each command should follow existing Cobra patterns in `internal/cli/`
- All commands return `error` — Cobra handles exit codes
- Every public function takes `ctx context.Context` as first arg
