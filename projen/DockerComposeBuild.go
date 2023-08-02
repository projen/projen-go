package projen


// Build arguments for creating a docker image.
// Experimental.
type DockerComposeBuild struct {
	// Docker build context directory.
	// Experimental.
	Context *string `field:"required" json:"context" yaml:"context"`
	// Build args.
	// Default: - none are provided.
	//
	// Experimental.
	Args *map[string]*string `field:"optional" json:"args" yaml:"args"`
	// A dockerfile to build from.
	// Default: "Dockerfile".
	//
	// Experimental.
	Dockerfile *string `field:"optional" json:"dockerfile" yaml:"dockerfile"`
}

