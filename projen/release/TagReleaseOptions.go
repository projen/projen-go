package release


// Experimental.
type TagReleaseOptions struct {
	// Tag patterns for which pushes should trigger a release.
	// Experimental.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

