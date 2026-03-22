# Testing Guide

## Unit tests

Run unit tests from the repository root:

```bash
go test ./test/unit/...
```

(If `test/unit` is not present yet, use `go test ./...` and narrow packages as they are added.)

## E2E tests

End-to-end tests exercise the binary and fixtures:

```bash
go test ./test/e2e/...
```

## Manual testing on DGX Spark

1. Cross-compile for Linux arm64: `make build-linux-arm64`
2. Copy the binary to the Spark host (for example with `scp`)
3. On the host: run `defenseclaw init`, then `defenseclaw scan skill` against fixtures under `test/fixtures/skills/`
4. Optional: run `bash scripts/test-e2e-spark.sh` after building `./defenseclaw` on that platform

## Manual testing on macOS

1. Build for Apple Silicon: `make build-darwin-arm64`
2. Run `./defenseclaw init` and `./defenseclaw scan skill ./test/fixtures/skills/clean-skill/`
3. Optional: `bash scripts/test-e2e-mac.sh` from the repo root with `./defenseclaw` present

## Test fixtures

| Path | Purpose |
|------|---------|
| `test/fixtures/skills/clean-skill/` | Minimal skill with manifest and benign `skill.py` for happy-path scans |
| `test/fixtures/skills/malicious-skill/` | Skill with intentional anti-patterns for scanner and policy testing |
| `test/fixtures/mcps/clean-mcp.json` | Benign MCP manifest for MCP scanner tests |
| `test/fixtures/mcps/malicious-mcp.json` | MCP manifest with suspicious tool description for detection tests |
| `test/fixtures/code/clean.py` | Simple Python that should pass static checks |
| `test/fixtures/code/hardcoded-secret.py` | Intentional hardcoded credential for CodeGuard-style tests |
