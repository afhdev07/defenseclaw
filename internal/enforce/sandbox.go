package enforce

type SandboxPolicy struct{}

func NewSandboxPolicy() *SandboxPolicy { return &SandboxPolicy{} }
