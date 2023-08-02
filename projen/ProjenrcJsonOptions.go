package projen


// Experimental.
type ProjenrcJsonOptions struct {
	// The name of the projenrc file.
	// Default: ".projenrc.json"
	//
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
}

