package polaris


// Options for `PolarisCoverityGo`.
//
// Extends base options with Go-specific defaults.
// Experimental.
type PolarisCoverityGoOptions struct {
	// Experimental.
	Commit *CommitConfiguration `field:"required" json:"commit" yaml:"commit"`
	// Experimental.
	Analyze *AnalysisConfiguration `field:"optional" json:"analyze" yaml:"analyze"`
	// Experimental.
	Caching *CachingConfiguration `field:"optional" json:"caching" yaml:"caching"`
	// Experimental.
	Capture *CaptureConfiguration `field:"optional" json:"capture" yaml:"capture"`
	// Specifies the version of the configuration file in use.
	// Experimental.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

