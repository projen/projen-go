//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewLambdaRuntimeParameters(functionRuntime *string, esbuildTarget *string, options *LambdaRuntimeOptions) error {
	if functionRuntime == nil {
		return fmt.Errorf("parameter functionRuntime is required, but nil was provided")
	}

	if esbuildTarget == nil {
		return fmt.Errorf("parameter esbuildTarget is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

