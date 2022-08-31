// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Volume binding information.
// Experimental.
type IDockerComposeVolumeBinding interface {
	// Binds the requested volume to the docker-compose volume configuration and provide mounting instructions for synthesis.
	//
	// Returns: mounting instructions for the service.
	// Experimental.
	Bind(volumeConfig IDockerComposeVolumeConfig) *DockerComposeVolumeMount
}

// The jsii proxy for IDockerComposeVolumeBinding
type jsiiProxy_IDockerComposeVolumeBinding struct {
	_ byte // padding
}

func (i *jsiiProxy_IDockerComposeVolumeBinding) Bind(volumeConfig IDockerComposeVolumeConfig) *DockerComposeVolumeMount {
	if err := i.validateBindParameters(volumeConfig); err != nil {
		panic(err)
	}
	var returns *DockerComposeVolumeMount

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{volumeConfig},
		&returns,
	)

	return returns
}

