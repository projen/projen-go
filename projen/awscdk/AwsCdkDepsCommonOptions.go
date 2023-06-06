package awscdk


// Options for `AwsCdkDeps`.
// Experimental.
type AwsCdkDepsCommonOptions struct {
	// Minimum version of the AWS CDK to depend on.
	// Experimental.
	CdkVersion *string `field:"required" json:"cdkVersion" yaml:"cdkVersion"`
	// Warning: NodeJS only.
	//
	// Install the.
	// Deprecated: The.
	CdkAssert *bool `field:"optional" json:"cdkAssert" yaml:"cdkAssert"`
	// Install the assertions library?
	//
	// Only needed for CDK 1.x. If using CDK 2.x then
	// assertions is already included in 'aws-cdk-lib'.
	// Experimental.
	CdkAssertions *bool `field:"optional" json:"cdkAssertions" yaml:"cdkAssertions"`
	// Which AWS CDKv1 modules this project requires.
	// Deprecated: For CDK 2.x use "deps" instead. (or "peerDeps" if you're building a library)
	CdkDependencies *[]*string `field:"optional" json:"cdkDependencies" yaml:"cdkDependencies"`
	// If this is enabled (default), all modules declared in `cdkDependencies` will be also added as normal `dependencies` (as well as `peerDependencies`).
	//
	// This is to ensure that downstream consumers actually have your CDK dependencies installed
	// when using npm < 7 or yarn, where peer dependencies are not automatically installed.
	// If this is disabled, `cdkDependencies` will be added to `devDependencies` to ensure
	// they are present during development.
	//
	// Note: this setting only applies to construct library projects.
	// Deprecated: Not supported in CDK v2.
	CdkDependenciesAsDeps *bool `field:"optional" json:"cdkDependenciesAsDeps" yaml:"cdkDependenciesAsDeps"`
	// AWS CDK modules required for testing.
	// Deprecated: For CDK 2.x use 'devDeps' (in node.js projects) or 'testDeps' (in java projects) instead
	CdkTestDependencies *[]*string `field:"optional" json:"cdkTestDependencies" yaml:"cdkTestDependencies"`
	// Use pinned version instead of caret version for CDK.
	//
	// You can use this to prevent mixed versions for your CDK dependencies and to prevent auto-updates.
	// If you use experimental features this will let you define the moment you include breaking changes.
	// Experimental.
	CdkVersionPinning *bool `field:"optional" json:"cdkVersionPinning" yaml:"cdkVersionPinning"`
	// Minimum version of the `constructs` library to depend on.
	// Experimental.
	ConstructsVersion *string `field:"optional" json:"constructsVersion" yaml:"constructsVersion"`
}

