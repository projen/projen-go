package release


// Experimental.
type ContinuousReleaseOptions struct {
	// Paths for which pushes should trigger a release.
	// Experimental.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

