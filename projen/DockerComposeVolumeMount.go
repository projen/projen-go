// CDK for software projects
package projen


// Service volume mounting information.
// Experimental.
type DockerComposeVolumeMount struct {
	// Volume source.
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Volume target.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Type of volume.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

