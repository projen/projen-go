//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func validateNewComponentParameters(project Project) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

