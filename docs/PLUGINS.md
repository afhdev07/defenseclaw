# Plugin Development Guide

Custom scanners extend DefenseClaw without rewriting the core binary. Plugins implement the same abstraction as built-in scanners.

## Scanner interface (Go)

Built-in and plugin scanners satisfy this interface (from `internal/scanner`):

```go
type Scanner interface {
	Name() string
	Version() string
	SupportedTargets() []string
	Scan(ctx context.Context, target string) (*ScanResult, error)
}
```

DefenseClaw discovers plugin binaries under `~/.defenseclaw/plugins/` at `init` and registers them when they expose the expected gRPC service contract.

Plugins are **standalone binaries** that implement the scanner gRPC service; the main CLI shells out or connects according to the plugin protocol. Full plugin packaging and examples are planned for **iteration 5**.
