package gitlab


// Specifies the docker image to use for the job or globally for all jobs.
//
// Job configuration
// takes precedence over global setting. Requires a certain kind of Gitlab runner executor.
// See: https://docs.gitlab.com/ee/ci/yaml/#image
//
// Experimental.
type Image struct {
	// Full name of the image that should be used.
	//
	// It should contain the Registry part if needed.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Command or script that should be executed as the container's entrypoint.
	//
	// It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array.
	// Experimental.
	Entrypoint *[]interface{} `field:"optional" json:"entrypoint" yaml:"entrypoint"`
}

