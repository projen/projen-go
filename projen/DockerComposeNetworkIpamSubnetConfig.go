package projen


// IPAM subnet configuration.
// Experimental.
type DockerComposeNetworkIpamSubnetConfig struct {
	// Subnet in CIDR format that represents a network segment.
	// Default: - value is not provided.
	//
	// Experimental.
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
}

