//go:build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerComposeService) validateAddDependsOnParameters(serviceName IDockerComposeServiceName) error {
	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddEnvironmentParameters(name *string, value *string) error {
	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddLabelParameters(name *string, value *string) error {
	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddNetworkParameters(network IDockerComposeNetworkBinding) error {
	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddPortParameters(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) error {
	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddVolumeParameters(volume IDockerComposeVolumeBinding) error {
	return nil
}

func validateNewDockerComposeServiceParameters(serviceName *string, serviceDescription *DockerComposeServiceDescription) error {
	return nil
}

