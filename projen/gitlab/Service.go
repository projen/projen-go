package gitlab


// Used to specify an additional Docker image to run scripts in.
//
// The service image is linked to the image specified in the.
// See: https://docs.gitlab.com/ee/ci/yaml/#services
//
// Experimental.
type Service struct {
	// Full name of the image that should be used.
	//
	// It should contain the Registry part if needed.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Additional alias that can be used to access the service from the job's container.
	//
	// Read Accessing the services for more information.
	// Experimental.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Command or script that should be used as the container's command.
	//
	// It will be translated to arguments passed to Docker after the image's name. The syntax is similar to Dockerfile's CMD directive, where each shell token is a separate string in the array.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Command or script that should be executed as the container's entrypoint.
	//
	// It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array.
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
}

