// CDK for software projects
package projen


// Network protocol for port mapping.
// Experimental.
type DockerComposeProtocol string

const (
	// TCP protocol.
	// Experimental.
	DockerComposeProtocol_TCP DockerComposeProtocol = "TCP"
	// UDP protocol.
	// Experimental.
	DockerComposeProtocol_UDP DockerComposeProtocol = "UDP"
)

