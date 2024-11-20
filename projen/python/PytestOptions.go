package python


// Experimental.
type PytestOptions struct {
	// Stop the testing process after the first N failures.
	// Experimental.
	MaxFailures *float64 `field:"optional" json:"maxFailures" yaml:"maxFailures"`
	// Location of sample tests.
	//
	// Typically the same directory where project tests will be located.
	// Default: "tests".
	//
	// Deprecated: Reference `sampleTestdir` on the project instead; to change the directory where tests are discovered from, use `testMatch`.
	Testdir *string `field:"optional" json:"testdir" yaml:"testdir"`
	// List of paths to test files or directories.
	//
	// Useful when all project tests are in a known location to speed up
	// test collection and to avoid picking up undesired tests by accident.
	//
	// Leave empty to discover all test_*.py or *_test.py files, per Pytest default.
	//
	// The array will be concatenated and passed as a single argument to pytest.
	//
	// Example:
	//   ["tests/unit", "tests/qa"]
	//
	// Default: [].
	//
	// Experimental.
	TestMatch *[]*string `field:"optional" json:"testMatch" yaml:"testMatch"`
	// Pytest version.
	// Default: "7.4.3"
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

