package awscdk


// Options for `LambdaAutoDiscover`.
// Experimental.
type LambdaAutoDiscoverOptions struct {
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
}

