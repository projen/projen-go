package polaris


// Options for `PolarisCoverityJava`.
//
// Extends base options with Java-specific defaults.
// Experimental.
type PolarisCoverityJavaOptions struct {
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

