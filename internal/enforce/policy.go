package enforce

import "context"

type PolicyEngine struct{}

func NewPolicyEngine() *PolicyEngine { return &PolicyEngine{} }

func (e *PolicyEngine) CheckBlockList(_ context.Context, _ string) (bool, error) {
	return false, nil
}

func (e *PolicyEngine) CheckAllowList(_ context.Context, _ string) (bool, error) {
	return false, nil
}
