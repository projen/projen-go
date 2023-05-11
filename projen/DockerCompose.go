package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Create a docker-compose YAML file.
// Experimental.
type DockerCompose interface {
	Component
	// Experimental.
	Project() Project
	// Add a service to the docker-compose file.
	// Experimental.
	AddService(serviceName *string, description *DockerComposeServiceDescription) DockerComposeService
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for DockerCompose
type jsiiProxy_DockerCompose struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_DockerCompose) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewDockerCompose(project Project, props *DockerComposeProps) DockerCompose {
	_init_.Initialize()

	if err := validateNewDockerComposeParameters(project, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DockerCompose{}

	_jsii_.Create(
		"projen.DockerCompose",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDockerCompose_Override(d DockerCompose, project Project, props *DockerComposeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.DockerCompose",
		[]interface{}{project, props},
		d,
	)
}

// Create a bind volume that binds a host path to the target path in the container.
// Experimental.
func DockerCompose_BindVolume(sourcePath *string, targetPath *string) IDockerComposeVolumeBinding {
	_init_.Initialize()

	if err := validateDockerCompose_BindVolumeParameters(sourcePath, targetPath); err != nil {
		panic(err)
	}
	var returns IDockerComposeVolumeBinding

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"bindVolume",
		[]interface{}{sourcePath, targetPath},
		&returns,
	)

	return returns
}

// Create a named volume and mount it to the target path.
//
// If you use this
// named volume in several services, the volume will be shared. In this
// case, the volume configuration of the first-provided options are used.
// Experimental.
func DockerCompose_NamedVolume(volumeName *string, targetPath *string, options *DockerComposeVolumeConfig) IDockerComposeVolumeBinding {
	_init_.Initialize()

	if err := validateDockerCompose_NamedVolumeParameters(volumeName, targetPath, options); err != nil {
		panic(err)
	}
	var returns IDockerComposeVolumeBinding

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"namedVolume",
		[]interface{}{volumeName, targetPath, options},
		&returns,
	)

	return returns
}

// Create a named network and mount it to the target path.
//
// If you use this
// named network in several services, the network will be shared. In this
// case, the network configuration of the first-provided options are used.
// Experimental.
func DockerCompose_Network(networkName *string, options *DockerComposeNetworkConfig) IDockerComposeNetworkBinding {
	_init_.Initialize()

	if err := validateDockerCompose_NetworkParameters(networkName, options); err != nil {
		panic(err)
	}
	var returns IDockerComposeNetworkBinding

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"network",
		[]interface{}{networkName, options},
		&returns,
	)

	return returns
}

// Create a port mapping.
// Experimental.
func DockerCompose_PortMapping(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) *DockerComposeServicePort {
	_init_.Initialize()

	if err := validateDockerCompose_PortMappingParameters(publishedPort, targetPort, options); err != nil {
		panic(err)
	}
	var returns *DockerComposeServicePort

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"portMapping",
		[]interface{}{publishedPort, targetPort, options},
		&returns,
	)

	return returns
}

// Depends on a service name.
// Experimental.
func DockerCompose_ServiceName(serviceName *string) IDockerComposeServiceName {
	_init_.Initialize()

	if err := validateDockerCompose_ServiceNameParameters(serviceName); err != nil {
		panic(err)
	}
	var returns IDockerComposeServiceName

	_jsii_.StaticInvoke(
		"projen.DockerCompose",
		"serviceName",
		[]interface{}{serviceName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerCompose) AddService(serviceName *string, description *DockerComposeServiceDescription) DockerComposeService {
	if err := d.validateAddServiceParameters(serviceName, description); err != nil {
		panic(err)
	}
	var returns DockerComposeService

	_jsii_.Invoke(
		d,
		"addService",
		[]interface{}{serviceName, description},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerCompose) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DockerCompose) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DockerCompose) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

