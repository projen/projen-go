package awscdk


// Options for `IntegrationTest`.
// Experimental.
type IntegrationTestOptions struct {
	// Destroy the test app after a successful deployment.
	//
	// If disabled, leaves the
	// app deployed in the dev account.
	// Default: true.
	//
	// Experimental.
	DestroyAfterDeploy *bool `field:"optional" json:"destroyAfterDeploy" yaml:"destroyAfterDeploy"`
	// Enables path metadata, adding `aws:cdk:path`, with the defining construct's path, to the CloudFormation metadata for each synthesized resource.
	// Default: false.
	//
	// Experimental.
	PathMetadata *bool `field:"optional" json:"pathMetadata" yaml:"pathMetadata"`
	// A path from the project root directory to a TypeScript file which contains the integration test app.
	//
	// This is relative to the root directory of the project.
	//
	// Example:
	//   "test/subdir/foo.integ.ts"
	//
	// Experimental.
	Entrypoint *string `field:"required" json:"entrypoint" yaml:"entrypoint"`
	// The path of the tsconfig.json file to use when running integration test cdk apps.
	// Experimental.
	TsconfigPath *string `field:"required" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Name of the integration test.
	// Default: - Derived from the entrypoint filename.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// AWS CDK dependency manager.
	// Experimental.
	CdkDeps AwsCdkDeps `field:"required" json:"cdkDeps" yaml:"cdkDeps"`
	// A list of stacks within the integration test to deploy/destroy.
	// Default: ["**"].
	//
	// Experimental.
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
}

