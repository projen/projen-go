package projen


// Description of a docker-compose.yml service.
// Experimental.
type DockerComposeServiceDescription struct {
	// Provide a command to the docker container.
	// Default: - use the container's default command.
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Names of other services this service depends on.
	// Default: - no dependencies.
	//
	// Experimental.
	DependsOn *[]IDockerComposeServiceName `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Entrypoint to run in the container.
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Add environment variables.
	// Default: - no environment variables are provided.
	//
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Use a docker image.
	//
	// Note: You must specify either `build` or `image` key.
	// See: imageBuild.
	//
	// Experimental.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Build a docker image.
	//
	// Note: You must specify either `imageBuild` or `image` key.
	// See: image.
	//
	// Experimental.
	ImageBuild *DockerComposeBuild `field:"optional" json:"imageBuild" yaml:"imageBuild"`
	// Add labels.
	// Default: - no labels are provided.
	//
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Add some networks to the service.
	// See: DockerCompose.network () to create & mount a named network
	//
	// Experimental.
	Networks *[]IDockerComposeNetworkBinding `field:"optional" json:"networks" yaml:"networks"`
	// Map some ports.
	// Default: - no ports are mapped.
	//
	// Experimental.
	Ports *[]*DockerComposeServicePort `field:"optional" json:"ports" yaml:"ports"`
	// Mount some volumes into the service.
	//
	// Use one of the following to create volumes:.
	// See: DockerCompose.namedVolume () to create & mount a named volume
	//
	// Experimental.
	Volumes *[]IDockerComposeVolumeBinding `field:"optional" json:"volumes" yaml:"volumes"`
}

