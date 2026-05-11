package workflows


// Version requirement for Python tools.
// Experimental.
type PythonToolRequirement struct {
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Whether to enable automatic dependency caching.
	// Default: false.
	//
	// Experimental.
	Cache *bool `field:"optional" json:"cache" yaml:"cache"`
	// The package manager to use for caching (e.g. "pip", "pipenv", "poetry"). Required when `cache` is true.
	// Experimental.
	PackageManager *string `field:"optional" json:"packageManager" yaml:"packageManager"`
}

