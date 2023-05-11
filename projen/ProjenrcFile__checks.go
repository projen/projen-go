//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func validateProjenrcFile_OfParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func validateNewProjenrcFileParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

