//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNewAwsCdkDepsJsParameters(project projen.Project, options *AwsCdkDepsOptions) error {
	return nil
}

