package cdk8s


// Experimental.
type IntegrationTestAutoDiscoverOptions struct {
	// Test source tree.
	// Experimental.
	Testdir *string `field:"required" json:"testdir" yaml:"testdir"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
}

