//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

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

func validateNewDevContainerParameters(project projen.Project, options *DevContainerOptions) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

