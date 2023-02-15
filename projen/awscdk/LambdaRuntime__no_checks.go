//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNewLambdaRuntimeParameters(functionRuntime *string, esbuildTarget *string, options *LambdaRuntimeOptions) error {
	return nil
}

