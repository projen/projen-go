package polaris


// Experimental.
type ParseWarningsConfiguration struct {
	// Enables parse warnings, recovery warnings, and semantic warnings that are produced by the cov-build command so that they appear as defects in Coverity Connect.
	//
	// By default, this is disabled if the aggressiveness level is low, and enabled if the aggressiveness level is medium or high.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

