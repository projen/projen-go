package polaris


// Specifies how the CLI should handle caching when performing capture/analysis.
// Experimental.
type CachingConfiguration struct {
	// A true value indicates caching will be used when performing remote analysis.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

