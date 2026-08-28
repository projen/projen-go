package polaris


// Experimental.
type ParallelTranslateConfiguration struct {
	// Specifies whether cov-translate parallelization should be enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the number of cov-emit processes to be run in parallel by cov-translate when multiple files are seen on a single native compiler invocation.
	//
	// A value of 0 will use the number of logical processors in the machine.
	// Experimental.
	Processes *float64 `field:"optional" json:"processes" yaml:"processes"`
}

