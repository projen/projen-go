//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNewLambdaAutoDiscoverParameters(project projen.Project, options *LambdaAutoDiscoverOptions) error {
	return nil
}

