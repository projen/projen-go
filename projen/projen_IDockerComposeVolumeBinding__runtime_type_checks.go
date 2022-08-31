//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"
)

func (i *jsiiProxy_IDockerComposeVolumeBinding) validateBindParameters(volumeConfig IDockerComposeVolumeConfig) error {
	if volumeConfig == nil {
		return fmt.Errorf("parameter volumeConfig is required, but nil was provided")
	}

	return nil
}

