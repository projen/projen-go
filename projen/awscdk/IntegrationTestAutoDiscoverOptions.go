package awscdk


// Options for `IntegrationTestAutoDiscover`.
// Experimental.
type IntegrationTestAutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Test source tree.
	// Experimental.
	Testdir *string `field:"required" json:"testdir" yaml:"testdir"`
	// Options for integration tests.
	// Experimental.
	IntegrationTestOptions *IntegrationTestCommonOptions `field:"optional" json:"integrationTestOptions" yaml:"integrationTestOptions"`
}

