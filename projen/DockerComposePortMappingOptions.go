package projen


// Options for port mappings.
// Experimental.
type DockerComposePortMappingOptions struct {
	// Port mapping protocol.
	// Default: DockerComposeProtocol.TCP
	//
	// Experimental.
	Protocol DockerComposeProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

