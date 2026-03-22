package scanner

import "context"

// Scanner defines the interface that all scanner implementations must satisfy.
// Built-in scanners wrap external CLI tools. Plugins implement this interface
// as standalone gRPC binaries discovered in ~/.defenseclaw/plugins/.
type Scanner interface {
	Name() string
	Version() string
	SupportedTargets() []string
	Scan(ctx context.Context, target string) (*ScanResult, error)
}
