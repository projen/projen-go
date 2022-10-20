//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNewLambdaExtensionAutoDiscoverParameters(project projen.Project, options *LambdaExtensionAutoDiscoverOptions) error {
	return nil
}

