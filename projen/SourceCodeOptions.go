// CDK for software projects
package projen


// Options for `SourceCodeFile`.
// Experimental.
type SourceCodeOptions struct {
	// Indentation size.
	// Experimental.
	Indent *float64 `field:"optional" json:"indent" yaml:"indent"`
	// Whether the generated file should be readonly.
	// Experimental.
	Readonly *bool `field:"optional" json:"readonly" yaml:"readonly"`
}

