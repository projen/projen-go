//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func validateTaskShell_CommandParameters(command *[]*string) error {
	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

