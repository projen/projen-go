//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IDockerComposeVolumeConfig) validateAddVolumeConfigurationParameters(volumeName *string, configuration *DockerComposeVolumeConfig) error {
	if volumeName == nil {
		return fmt.Errorf("parameter volumeName is required, but nil was provided")
	}

	if configuration == nil {
		return fmt.Errorf("parameter configuration is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(configuration, func() string { return "parameter configuration" }); err != nil {
		return err
	}

	return nil
}

