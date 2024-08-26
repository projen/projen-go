package cdk


// Options for `JsiiDocgen`.
// Experimental.
type JsiiDocgenOptions struct {
	// File path for generated docs.
	// Default: "API.md"
	//
	// Experimental.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// A semver version string to install a specific version of jsii-docgen.
	// Default: '*'.
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

