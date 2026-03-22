package sandbox

import "fmt"

type OpenShell struct {
	BinaryPath string
	PolicyDir  string
}

func New(binaryPath, policyDir string) *OpenShell {
	return &OpenShell{BinaryPath: binaryPath, PolicyDir: policyDir}
}

func (o *OpenShell) IsAvailable() bool {
	return false
}

func (o *OpenShell) Start() error {
	return fmt.Errorf("sandbox: openshell integration not yet implemented — coming in iteration 2")
}
