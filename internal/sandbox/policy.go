package sandbox

type Policy struct {
	AllowedEndpoints []string `yaml:"allowed_endpoints"`
	DeniedEndpoints  []string `yaml:"denied_endpoints"`
	Permissions      []string `yaml:"permissions"`
}

func DefaultPolicy() *Policy {
	return &Policy{}
}
