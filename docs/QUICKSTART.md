# Quick Start Guide

Get DefenseClaw running in under 5 minutes.

## Prerequisites

- **Go 1.22+** — to build from source
- **Python 3.9+** — for scanner dependencies
- **pip** — to install Python scanners

## 1. Build

```bash
git clone https://github.com/defenseclaw/defenseclaw.git
cd defenseclaw
make build
```

For DGX Spark (linux/arm64):

```bash
make build-linux-arm64
scp defenseclaw-linux-arm64 spark:/usr/local/bin/defenseclaw
```

## 2. Install Scanner Dependencies

```bash
bash scripts/setup-scanners.sh
```

This installs `skill-scanner`, `mcp-scanner`, and `aibom` via pip.

## 3. Initialize

```bash
defenseclaw init
```

This creates `~/.defenseclaw/` with:
- `config.yaml` — scanner paths, policy settings
- `audit.db` — SQLite audit log
- `quarantine/` — blocked skill storage
- `plugins/` — custom scanner plugins
- `policies/` — policy templates

## 4. First Scan

```bash
# Scan a clean skill
defenseclaw scan skill ./test/fixtures/skills/clean-skill/

# Scan a skill with known issues
defenseclaw scan skill ./test/fixtures/skills/malicious-skill/

# Generate AI bill of materials
defenseclaw scan aibom .

# Run all scanners against current directory
defenseclaw scan
```

## 5. Next Steps

- `defenseclaw block skill <name>` — block a skill (iteration 2)
- `defenseclaw tui` — open the dashboard (iteration 3)
- `defenseclaw deploy` — full orchestrated deploy (iteration 4)

See [CLI Reference](CLI.md) for all commands.
