//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_DockerComposeService) validateAddDependsOnParameters(serviceName IDockerComposeServiceName) error {
	if serviceName == nil {
		return fmt.Errorf("parameter serviceName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddEnvironmentParameters(name *string, value *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddNetworkParameters(network IDockerComposeNetworkBinding) error {
	if network == nil {
		return fmt.Errorf("parameter network is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DockerComposeService) validateAddPortParameters(publishedPort *float64, targetPort *float64, options *DockerComposePortMappingOptions) error {
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

func (d *jsiiProxy_DockerComposeService) validateAddVolumeParameters(volume IDockerComposeVolumeBinding) error {
	if volume == nil {
		return fmt.Errorf("parameter volume is required, but nil was provided")
	}

	return nil
}

func validateNewDockerComposeServiceParameters(serviceName *string, serviceDescription *DockerComposeServiceDescription) error {
	if serviceName == nil {
		return fmt.Errorf("parameter serviceName is required, but nil was provided")
	}

	if serviceDescription == nil {
		return fmt.Errorf("parameter serviceDescription is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(serviceDescription, func() string { return "parameter serviceDescription" }); err != nil {
		return err
	}

	return nil
}

