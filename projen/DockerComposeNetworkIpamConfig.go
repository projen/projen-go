package projen


// IPAM configuration.
// Experimental.
type DockerComposeNetworkIpamConfig struct {
	// A list with zero or more config blocks specifying custom IPAM configuration.
	// Default: - value is not provided.
	//
	// Experimental.
	Config *[]*DockerComposeNetworkIpamSubnetConfig `field:"optional" json:"config" yaml:"config"`
	// Driver to use for custom IPAM config.
	// Default: - value is not provided.
	//
	// Experimental.
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
}

