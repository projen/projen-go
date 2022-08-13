package python


// Options for python test code sample.
// Experimental.
type PytestSampleOptions struct {
	// Name of the python package as used in imports and filenames.
	// Experimental.
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// Test directory.
	// Experimental.
	Testdir *string `field:"required" json:"testdir" yaml:"testdir"`
}

