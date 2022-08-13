package workflows


// Options pertaining to container environments.
// Experimental.
type ContainerOptions struct {
	// The Docker image to use as the container to run the action.
	//
	// The value can
	// be the Docker Hub image name or a registry name.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// f the image's container registry requires authentication to pull the image, you can use credentials to set a map of the username and password.
	//
	// The credentials are the same values that you would provide to the docker
	// login command.
	// Experimental.
	Credentials *ContainerCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Sets a map of environment variables in the container.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Additional Docker container resource options.
	// See: https://docs.docker.com/engine/reference/commandline/create/#options
	//
	// Experimental.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Sets an array of ports to expose on the container.
	// Experimental.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
	// Sets an array of volumes for the container to use.
	//
	// You can use volumes to
	// share data between services or other steps in a job. You can specify
	// named Docker volumes, anonymous Docker volumes, or bind mounts on the
	// host.
	//
	// To specify a volume, you specify the source and destination path:
	// `<source>:<destinationPath>`.
	// Experimental.
	Volumes *[]*string `field:"optional" json:"volumes" yaml:"volumes"`
}

