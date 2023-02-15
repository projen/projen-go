// CDK for software projects
package projen


// IPAM subnet configuration.
// Experimental.
type DockerComposeNetworkIpamSubnetConfig struct {
	// Subnet in CIDR format that represents a network segment.
	// Experimental.
	Subnet *string `field:"optional" json:"subnet" yaml:"subnet"`
}

