package projen


// Options for `SourceCodeFile`.
// Experimental.
type SourceCodeOptions struct {
	// Indentation size.
	// Default: 2.
	//
	// Experimental.
	Indent *float64 `field:"optional" json:"indent" yaml:"indent"`
	// Whether the generated file should be readonly.
	// Default: true.
	//
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

