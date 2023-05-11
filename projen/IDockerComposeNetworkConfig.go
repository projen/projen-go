package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Storage for network configuration.
// Experimental.
type IDockerComposeNetworkConfig interface {
	// Add network configuration to the repository.
	// Experimental.
	AddNetworkConfiguration(networkName *string, configuration *DockerComposeNetworkConfig)
}

// The jsii proxy for IDockerComposeNetworkConfig
type jsiiProxy_IDockerComposeNetworkConfig struct {
	_ byte // padding
}

func (i *jsiiProxy_IDockerComposeNetworkConfig) AddNetworkConfiguration(networkName *string, configuration *DockerComposeNetworkConfig) {
	if err := i.validateAddNetworkConfigurationParameters(networkName, configuration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addNetworkConfiguration",
		[]interface{}{networkName, configuration},
	)
}

