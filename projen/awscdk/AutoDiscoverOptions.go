package awscdk


// Options for `AutoDiscover`.
// Experimental.
type AutoDiscoverOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Project source tree (relative to project output directory).
	// Experimental.
	Srcdir *string `field:"required" json:"srcdir" yaml:"srcdir"`
	// Options for AWS Lambda functions.
	// Experimental.
	LambdaOptions *LambdaFunctionCommonOptions `field:"optional" json:"lambdaOptions" yaml:"lambdaOptions"`
	// Options for lambda extensions.
	// Experimental.
	LambdaExtensionOptions *LambdaExtensionCommonOptions `field:"optional" json:"lambdaExtensionOptions" yaml:"lambdaExtensionOptions"`
	// Test source tree.
	// Experimental.
	Testdir *string `field:"required" json:"testdir" yaml:"testdir"`
	// Options for integration tests.
	// Experimental.
	IntegrationTestOptions *IntegrationTestCommonOptions `field:"optional" json:"integrationTestOptions" yaml:"integrationTestOptions"`
	// Auto-discover edge lambda functions.
	// Default: true.
	//
	// Experimental.
	EdgeLambdaAutoDiscover *bool `field:"optional" json:"edgeLambdaAutoDiscover" yaml:"edgeLambdaAutoDiscover"`
	// Auto-discover integration tests.
	// Default: true.
	//
	// Experimental.
	IntegrationTestAutoDiscover *bool `field:"optional" json:"integrationTestAutoDiscover" yaml:"integrationTestAutoDiscover"`
	// Auto-discover lambda functions.
	// Default: true.
	//
	// Experimental.
	LambdaAutoDiscover *bool `field:"optional" json:"lambdaAutoDiscover" yaml:"lambdaAutoDiscover"`
	// Auto-discover lambda extensions.
	// Default: true.
	//
	// Experimental.
	LambdaExtensionAutoDiscover *bool `field:"optional" json:"lambdaExtensionAutoDiscover" yaml:"lambdaExtensionAutoDiscover"`
}

