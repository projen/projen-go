// CDK for software projects
package projen


// Description of a docker-compose.yml service.
// Experimental.
type DockerComposeServiceDescription struct {
	// Provide a command to the docker container.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Names of other services this service depends on.
	// Experimental.
	DependsOn *[]IDockerComposeServiceName `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Add environment variables.
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
	// Map some ports.
	// Experimental.
	Ports *[]*DockerComposeServicePort `field:"optional" json:"ports" yaml:"ports"`
	// Mount some volumes into the service.
	//
	// Use one of the following to create volumes:.
	// See: DockerCompose.namedVolume() to create & mount a named volume
	//
	// Experimental.
	Volumes *[]IDockerComposeVolumeBinding `field:"optional" json:"volumes" yaml:"volumes"`
}

