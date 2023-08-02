package projen


// Deprecated: use `ProjenrcJsonOptions`.
type ProjenrcOptions struct {
	// The name of the projenrc file.
	// Default: ".projenrc.json"
	//
	// Deprecated: use `ProjenrcJsonOptions`.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
}

