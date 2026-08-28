package polaris


// Experimental.
type CompilerConfiguration struct {
	// Specifies a list of arguments to pass to "cov-configure" to generate the compiler configuration to use during capture.
	//
	// This key is mutually exclusive with the "file" key.
	// Experimental.
	CovConfigure *[]*[]*string `field:"optional" json:"covConfigure" yaml:"covConfigure"`
	// Specifies a pre-generated compiler configuration file to use.
	//
	// This key is mutually exclusive with the "cov-configure" key.
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
}

