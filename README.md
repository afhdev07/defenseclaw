# DefenseClaw

Enterprise governance layer for [OpenClaw](https://github.com/nvidia/openclaw). Scans skills, MCP servers, and code before they run. Enforces block/allow lists. Provides a terminal dashboard for security operators.

## Quick Start

```bash
# Install (coming soon)
curl -sSf https://get.defenseclaw.dev | sh

# Initialize
defenseclaw init

# Scan a skill
defenseclaw scan skill ./path/to/skill

# Open the dashboard
defenseclaw tui
```

## What It Does

- **Scan before run** — skills, MCP servers, A2A agents, code, AI dependencies
- **Block/allow lists** — operator-managed enforcement for skills and MCP servers
- **Terminal dashboard** — scan findings, policy violations, enforcement actions
- **Audit trail** — every action logged to SQLite with timestamps and context

## Documentation

- [Quick Start Guide](docs/QUICKSTART.md)
- [Architecture](docs/ARCHITECTURE.md)
- [CLI Reference](docs/CLI.md)
- [TUI Guide](docs/TUI.md)
- [Plugin Development](docs/PLUGINS.md)
- [Testing](docs/TESTING.md)
- [Contributing](docs/CONTRIBUTING.md)

## Building from Source

```bash
make build              # Current platform
make build-linux-arm64  # DGX Spark
make build-darwin-arm64 # Apple Silicon
make test               # Run tests
```

Requires Go 1.22+.

## License

Apache 2.0 — see [LICENSE](LICENSE).
