package awscdk


// Common options for auto discovering project subcomponents.
// Experimental.
type AutoDiscoverCommonOptions struct {
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// Path to the tsconfig file to use for integration tests.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
}

