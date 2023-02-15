package circleci


// Experimental.
type Machine struct {
	// The VM image to use.
	// See: https://circleci.com/docs/2.0/configuration-reference/#available-machine-images
	//
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// enable docker layer caching.
	// See: https://circleci.com/docs/2.0/configuration-reference/#available-machine-images
	//
	// Experimental.
	DockerLayerCaching *string `field:"optional" json:"dockerLayerCaching" yaml:"dockerLayerCaching"`
}

