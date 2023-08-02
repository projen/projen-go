package projen


// Network configuration.
// Experimental.
type DockerComposeNetworkConfig struct {
	// Set to true to indicate that standalone containers can attach to this network, in addition to services.
	// Default: - unset.
	//
	// Experimental.
	Attachable *bool `field:"optional" json:"attachable" yaml:"attachable"`
	// Set to true to indicate that the network is a bridge network.
	// Default: - unset.
	//
	// Experimental.
	Bridge *bool `field:"optional" json:"bridge" yaml:"bridge"`
	// Driver to use for the network.
	// Default: - value is not provided.
	//
	// Experimental.
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// Options for the configured driver.
	//
	// Those options are driver-dependent - consult the driver’s documentation for more information.
	// Default: - value is not provided.
	//
	// Experimental.
	DriverOpts *map[string]interface{} `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Set to true to indicate that the network is externally created.
	// Default: - unset, indicating that docker-compose creates the network.
	//
	// Experimental.
	External *bool `field:"optional" json:"external" yaml:"external"`
	// Set to true to indicate that you want to create an externally isolated overlay network.
	// Default: - unset.
	//
	// Experimental.
	Internal *bool `field:"optional" json:"internal" yaml:"internal"`
	// Specify custom IPAM config.
	// Default: - unset.
	//
	// Experimental.
	Ipam *DockerComposeNetworkIpamConfig `field:"optional" json:"ipam" yaml:"ipam"`
	// Attach labels to the network.
	// Default: - unset.
	//
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name of the network for when the network name isn't going to work in YAML.
	// Default: - unset, indicating that docker-compose creates networks as usual.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Set to true to indicate that the network is an overlay network.
	// Default: - unset.
	//
	// Experimental.
	Overlay *bool `field:"optional" json:"overlay" yaml:"overlay"`
}

