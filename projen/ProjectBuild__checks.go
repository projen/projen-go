//go:build !no_runtime_type_checking

// CDK for software projects
package projen

import (
	"fmt"
)

func validateNewProjectBuildParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

