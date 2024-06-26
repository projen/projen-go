//go:build !no_runtime_type_checking

package vscode

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

func (d *jsiiProxy_DevContainer) validateAddDockerImageParameters(image projen.DevEnvironmentDockerImage) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DevContainer) validateAddFeaturesParameters(features *[]*DevContainerFeature) error {
	for idx_5b8a8b, v := range *features {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter features[%#v]", idx_5b8a8b) }); err != nil {
			return err
		}
	}

	return nil
}

func validateDevContainer_IsComponentParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateDevContainer_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDevContainerParameters(project projen.Project, options *DevContainerOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

