//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"
)

func validateDevEnvironmentDockerImage_FromFileParameters(dockerFile *string) error {
	if dockerFile == nil {
		return fmt.Errorf("parameter dockerFile is required, but nil was provided")
	}

	return nil
}

func validateDevEnvironmentDockerImage_FromImageParameters(image *string) error {
	if image == nil {
		return fmt.Errorf("parameter image is required, but nil was provided")
	}

	return nil
}

