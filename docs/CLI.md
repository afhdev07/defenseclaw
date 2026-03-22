# CLI Reference

All subcommands are registered on `defenseclaw`. Use `defenseclaw help <command>` for flags and examples.

## Commands

| Command | Description | Iteration |
|--------|-------------|-----------|
| `init` | Create `~/.defenseclaw` config and SQLite audit database | 1 |
| `scan skill <path>` | Run skill-scanner on a skill directory | 1 |
| `scan mcp <path-or-url>` | Run mcp-scanner on MCP manifest or endpoint | 1 |
| `scan aibom <path>` | Run aibom inventory on a project path | 1 |
| `deploy` | Orchestrated deploy with admission gate (OpenShell) | 4 |
| `block skill <id>` | Add a skill to the block list | 2 |
| `block mcp <id>` | Add an MCP server to the block list | 2 |
| `allow skill <id>` | Add a skill to the allow list | 2 |
| `allow mcp <id>` | Add an MCP server to the allow list | 2 |
| `list blocked` | List blocked skills and MCP servers | 2 |
| `list allowed` | List allowed skills and MCP servers | 2 |
| `quarantine` | Manage quarantined artifacts | 2 |
| `rescan` | Re-run last scan or scan a target again | 2 |
| `alerts` | Show recent alerts from scans and enforcement | 3 |
| `status` | Show DefenseClaw and scanner status | 4 |
| `tui` | Launch the four-panel terminal dashboard | 3 |
| `audit` | Query or export audit events (JSON/CSV) | 1 |

Commands marked for later iterations may be stubs or partially implemented until that iteration lands; use `defenseclaw help` for the current behavior on your build.
