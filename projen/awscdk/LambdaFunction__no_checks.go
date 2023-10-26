//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateLambdaFunction_IsComponentParameters(x interface{}) error {
	return nil
}

func validateLambdaFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaFunctionParameters(project projen.Project, options *LambdaFunctionOptions) error {
	return nil
}

