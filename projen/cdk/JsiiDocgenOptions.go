package cdk


// Options for `JsiiDocgen`.
// Experimental.
type JsiiDocgenOptions struct {
	// File path for generated docs.
	// Default: "API.md"
	//
	// Experimental.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
}

