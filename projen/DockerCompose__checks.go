//go:build !no_runtime_type_checking

package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_DockerCompose) validateAddServiceParameters(serviceName *string, description *DockerComposeServiceDescription) error {
	if serviceName == nil {
		return fmt.Errorf("parameter serviceName is required, but nil was provided")
	}

	if description == nil {
		return fmt.Errorf("parameter description is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(description, func() string { return "parameter description" }); err != nil {
		return err
	}

	return nil
}

func validateDockerCompose_BindVolumeParameters(sourcePath *string, targetPath *string) error {
	if sourcePath == nil {
		return fmt.Errorf("parameter sourcePath is required, but nil was provided")
	}

	if targetPath == nil {
		return fmt.Errorf("parameter targetPath is required, but nil was provided")
	}

	return nil
}

func validateDockerCompose_NamedVolumeParameters(volumeName *string, targetPath *string, options *DockerComposeVolumeConfig) error {
	if volumeName == nil {
		return fmt.Errorf("parameter volumeName is required, but nil was provided")
	}

	if targetPath == nil {
		return fmt.Errorf("parameter targetPath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDockerCompose_NetworkParameters(networkName *string, options *DockerComposeNetworkConfig) error {
	if networkName == nil {
		return fmt.Errorf("parameter networkName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDockerCompose_PortMappingParameters(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) error {
	if publishedPort == nil {
		return fmt.Errorf("parameter publishedPort is required, but nil was provided")
	}

	if targetPort == nil {
		return fmt.Errorf("parameter targetPort is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDockerCompose_ServiceNameParameters(serviceName *string) error {
	if serviceName == nil {
		return fmt.Errorf("parameter serviceName is required, but nil was provided")
	}

	return nil
}

func validateNewDockerComposeParameters(project Project, props *DockerComposeProps) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

