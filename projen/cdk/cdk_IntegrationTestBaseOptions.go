package cdk


// Options for IntegrationTest.
// Experimental.
type IntegrationTestBaseOptions struct {
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
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

