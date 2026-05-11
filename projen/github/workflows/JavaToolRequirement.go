package workflows


// Version requirement for Java tools.
// Experimental.
type JavaToolRequirement struct {
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Whether to enable automatic dependency caching.
	// Default: false.
	//
	// Experimental.
	Cache *bool `field:"optional" json:"cache" yaml:"cache"`
	// The JDK distribution to use.
	// Default: "corretto".
	//
	// Experimental.
	Distribution *string `field:"optional" json:"distribution" yaml:"distribution"`
	// The package manager to use for caching (e.g. "maven", "gradle", "sbt"). Required when `cache` is true.
	// Experimental.
	PackageManager *string `field:"optional" json:"packageManager" yaml:"packageManager"`
}

