package circleci


// Options for docker executor.
// See: https://circleci.com/docs/2.0/configuration-reference/#docker
//
// Experimental.
type Docker struct {
	// The name of a custom docker image to use.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Authentication for registries using standard docker login credentials.
	// Experimental.
	Auth *map[string]*string `field:"optional" json:"auth" yaml:"auth"`
	// Authentication for AWS Elastic Container Registry (ECR).
	// Experimental.
	AwsAuth *map[string]*string `field:"optional" json:"awsAuth" yaml:"awsAuth"`
	// The command used as pid 1 (or args for entrypoint) when launching the container.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The command used as executable when launching the container.
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// A map of environment variable names and values.
	// Experimental.
	Environment *map[string]interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The name the container is reachable by.
	//
	// By default, container services are accessible through localhost.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Which user to run commands as within the Docker container.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
}

