package projen


// Volume configuration.
// Experimental.
type DockerComposeVolumeConfig struct {
	// Driver to use for the volume.
	// Default: - value is not provided.
	//
	// Experimental.
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// Options to provide to the driver.
	// Experimental.
	DriverOpts *map[string]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Set to true to indicate that the volume is externally created.
	// Default: - unset, indicating that docker-compose creates the volume.
	//
	// Experimental.
	External *bool `field:"optional" json:"external" yaml:"external"`
	// Name of the volume for when the volume name isn't going to work in YAML.
	// Default: - unset, indicating that docker-compose creates volumes as usual.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

