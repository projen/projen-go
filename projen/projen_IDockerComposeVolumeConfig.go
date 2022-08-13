// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Storage for volume configuration.
// Experimental.
type IDockerComposeVolumeConfig interface {
	// Add volume configuration to the repository.
	// Experimental.
	AddVolumeConfiguration(volumeName *string, configuration *DockerComposeVolumeConfig)
}

// The jsii proxy for IDockerComposeVolumeConfig
type jsiiProxy_IDockerComposeVolumeConfig struct {
	_ byte // padding
}

func (i *jsiiProxy_IDockerComposeVolumeConfig) AddVolumeConfiguration(volumeName *string, configuration *DockerComposeVolumeConfig) {
	_jsii_.InvokeVoid(
		i,
		"addVolumeConfiguration",
		[]interface{}{volumeName, configuration},
	)
}

