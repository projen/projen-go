package cdk8s


// Options for `AutoDiscover`.
// Experimental.
type AutoDiscoverOptions struct {
	// Test source tree.
	// Experimental.
	Testdir *string `field:"required" json:"testdir" yaml:"testdir"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Automatically discover integration tests.
	// Experimental.
	IntegrationTestAutoDiscover *bool `field:"optional" json:"integrationTestAutoDiscover" yaml:"integrationTestAutoDiscover"`
}

