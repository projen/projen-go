package projen


// A service port mapping.
// Experimental.
type DockerComposeServicePort struct {
	// Port mapping mode.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Network protocol.
	// Experimental.
	Protocol DockerComposeProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// Published port number.
	// Experimental.
	Published *float64 `field:"required" json:"published" yaml:"published"`
	// Target port number.
	// Experimental.
	Target *float64 `field:"required" json:"target" yaml:"target"`
}

