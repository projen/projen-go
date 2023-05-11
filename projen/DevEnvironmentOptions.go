package projen


// Base options for configuring a container-based development environment.
// Experimental.
type DevEnvironmentOptions struct {
	// A Docker image or Dockerfile for the container.
	// Experimental.
	DockerImage DevEnvironmentDockerImage `field:"optional" json:"dockerImage" yaml:"dockerImage"`
	// An array of ports that should be exposed from the container.
	// Experimental.
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
	// An array of tasks that should be run when the container starts.
	// Experimental.
	Tasks *[]Task `field:"optional" json:"tasks" yaml:"tasks"`
	// An array of extension IDs that specify the extensions that should be installed inside the container when it is created.
	// Experimental.
	VscodeExtensions *[]*string `field:"optional" json:"vscodeExtensions" yaml:"vscodeExtensions"`
}

