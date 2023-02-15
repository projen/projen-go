// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An interface providing the name of a docker compose service.
// Experimental.
type IDockerComposeServiceName interface {
	// The name of the docker compose service.
	// Experimental.
	ServiceName() *string
}

// The jsii proxy for IDockerComposeServiceName
type jsiiProxy_IDockerComposeServiceName struct {
	_ byte // padding
}

func (j *jsiiProxy_IDockerComposeServiceName) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

