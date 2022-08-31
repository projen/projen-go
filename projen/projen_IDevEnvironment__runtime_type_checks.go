//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"
)

func (i *jsiiProxy_IDevEnvironment) validateAddDockerImageParameters(image DevEnvironmentDockerImage) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

