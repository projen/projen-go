package workflows


// Version requirement for tools.
// Experimental.
type ToolRequirement struct {
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Whether to enable automatic dependency caching.
	// Default: false.
	//
	// Experimental.
	Cache *bool `field:"optional" json:"cache" yaml:"cache"`
}

