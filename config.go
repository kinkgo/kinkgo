package kinkgo

// Config holds runtime configuration for the TestSuite.
type Config[E Environment] struct {
	// Environment is the environment to run the test suite in.
	Environment E

	// Labels is a list of labels to apply to the test suite.
	Labels []string
}

// NewNopEnvironmentConfig returns a new Config with a NopEnvironment and labels.
func NewNopEnvironmentConfig(labels ...string) Config[NopEnvironment] {
	return Config[NopEnvironment]{
		Environment: NopEnvironment{},
		Labels:      labels,
	}
}

// NewConfig returns a new Config with the given environment and labels.
func NewConfig[E Environment](env E, labels ...string) Config[E] {
	return Config[E]{
		Environment: env,
		Labels:      labels,
	}
}
