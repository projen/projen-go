//go:build !no_runtime_type_checking

package projen

import (
	"fmt"
)

func (i *jsiiProxy_IScriptRunner) validateConfigForParameters(entrypoint *string) error {
	if entrypoint == nil {
		return fmt.Errorf("parameter entrypoint is required, but nil was provided")
	}

	return nil
}

