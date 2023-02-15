//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNewAwsCdkDepsParameters(project projen.Project, options *AwsCdkDepsOptions) error {
	return nil
}

