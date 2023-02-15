// CDK for software projects
package projen


// Options for port mappings.
// Experimental.
type DockerComposePortMappingOptions struct {
	// Port mapping protocol.
	// Experimental.
	Protocol DockerComposeProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

