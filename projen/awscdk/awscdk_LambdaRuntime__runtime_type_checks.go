//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"
)

func validateNewLambdaRuntimeParameters(functionRuntime *string, esbuildTarget *string) error {
	if functionRuntime == nil {
		return fmt.Errorf("parameter functionRuntime is required, but nil was provided")
	}

	if esbuildTarget == nil {
		return fmt.Errorf("parameter esbuildTarget is required, but nil was provided")
	}

	return nil
}

