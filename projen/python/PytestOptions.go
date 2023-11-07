package python


// Experimental.
type PytestOptions struct {
	// Stop the testing process after the first N failures.
	// Experimental.
	MaxFailures *float64 `field:"optional" json:"maxFailures" yaml:"maxFailures"`
	// Directory with tests.
	// Default: 'tests'.
	//
	// Experimental.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
	// Pytest version.
	// Default: "7.4.3"
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

