package gitlab


// Asset configuration for a release.
// Experimental.
type Assets struct {
	// Include asset links in the release.
	// Experimental.
	Links *[]*Link `field:"required" json:"links" yaml:"links"`
}

