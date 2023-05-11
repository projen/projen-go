package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Network binding information.
// Experimental.
type IDockerComposeNetworkBinding interface {
	// Binds the requested network to the docker-compose network configuration and provide mounting instructions for synthesis.
	//
	// Returns: the service name.
	// Experimental.
	Bind(networkConfig IDockerComposeNetworkConfig) *string
}

// The jsii proxy for IDockerComposeNetworkBinding
type jsiiProxy_IDockerComposeNetworkBinding struct {
	_ byte // padding
}

func (i *jsiiProxy_IDockerComposeNetworkBinding) Bind(networkConfig IDockerComposeNetworkConfig) *string {
	if err := i.validateBindParameters(networkConfig); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{networkConfig},
		&returns,
	)

	return returns
}

