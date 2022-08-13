// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// A docker-compose service.
// Experimental.
type DockerComposeService interface {
	IDockerComposeServiceName
	// Command to run in the container.
	// Experimental.
	Command() *[]*string
	// Other services that this service depends on.
	// Experimental.
	DependsOn() *[]IDockerComposeServiceName
	// Environment variables.
	// Experimental.
	Environment() *map[string]*string
	// Docker image.
	// Experimental.
	Image() *string
	// Docker image build instructions.
	// Experimental.
	ImageBuild() *DockerComposeBuild
	// Published ports.
	// Experimental.
	Ports() *[]*DockerComposeServicePort
	// Name of the service.
	// Experimental.
	ServiceName() *string
	// Volumes mounted in the container.
	// Experimental.
	Volumes() *[]IDockerComposeVolumeBinding
	// Make the service depend on another service.
	// Experimental.
	AddDependsOn(serviceName IDockerComposeServiceName)
	// Add an environment variable.
	// Experimental.
	AddEnvironment(name *string, value *string)
	// Add a port mapping.
	// Experimental.
	AddPort(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions)
	// Add a volume to the service.
	// Experimental.
	AddVolume(volume IDockerComposeVolumeBinding)
}

// The jsii proxy struct for DockerComposeService
type jsiiProxy_DockerComposeService struct {
	jsiiProxy_IDockerComposeServiceName
}

func (j *jsiiProxy_DockerComposeService) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) DependsOn() *[]IDockerComposeServiceName {
	var returns *[]IDockerComposeServiceName
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) ImageBuild() *DockerComposeBuild {
	var returns *DockerComposeBuild
	_jsii_.Get(
		j,
		"imageBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Ports() *[]*DockerComposeServicePort {
	var returns *[]*DockerComposeServicePort
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerComposeService) Volumes() *[]IDockerComposeVolumeBinding {
	var returns *[]IDockerComposeVolumeBinding
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}


// Experimental.
func NewDockerComposeService(serviceName *string, serviceDescription *DockerComposeServiceDescription) DockerComposeService {
	_init_.Initialize()

	j := jsiiProxy_DockerComposeService{}

	_jsii_.Create(
		"projen.DockerComposeService",
		[]interface{}{serviceName, serviceDescription},
		&j,
	)

	return &j
}

// Experimental.
func NewDockerComposeService_Override(d DockerComposeService, serviceName *string, serviceDescription *DockerComposeServiceDescription) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.DockerComposeService",
		[]interface{}{serviceName, serviceDescription},
		d,
	)
}

func (d *jsiiProxy_DockerComposeService) AddDependsOn(serviceName IDockerComposeServiceName) {
	_jsii_.InvokeVoid(
		d,
		"addDependsOn",
		[]interface{}{serviceName},
	)
}

func (d *jsiiProxy_DockerComposeService) AddEnvironment(name *string, value *string) {
	_jsii_.InvokeVoid(
		d,
		"addEnvironment",
		[]interface{}{name, value},
	)
}

func (d *jsiiProxy_DockerComposeService) AddPort(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) {
	_jsii_.InvokeVoid(
		d,
		"addPort",
		[]interface{}{publishedPort, targetPort, options},
	)
}

func (d *jsiiProxy_DockerComposeService) AddVolume(volume IDockerComposeVolumeBinding) {
	_jsii_.InvokeVoid(
		d,
		"addVolume",
		[]interface{}{volume},
	)
}

