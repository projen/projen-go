//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package typescript

import (
	"fmt"
)

func validateNewTypedocDocgenParameters(project TypeScriptProject) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

