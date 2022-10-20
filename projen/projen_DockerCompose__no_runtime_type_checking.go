//go:build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerCompose) validateAddServiceParameters(serviceName *string, description *DockerComposeServiceDescription) error {
	return nil
}

func validateDockerCompose_BindVolumeParameters(sourcePath *string, targetPath *string) error {
	return nil
}

func validateDockerCompose_NamedVolumeParameters(volumeName *string, targetPath *string, options *DockerComposeVolumeConfig) error {
	return nil
}

func validateDockerCompose_PortMappingParameters(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) error {
	return nil
}

func validateDockerCompose_ServiceNameParameters(serviceName *string) error {
	return nil
}

func validateNewDockerComposeParameters(project Project, props *DockerComposeProps) error {
	return nil
}

