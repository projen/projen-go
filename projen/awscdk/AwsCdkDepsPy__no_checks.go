//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateAwsCdkDepsPy_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAwsCdkDepsPy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAwsCdkDepsPyParameters(project projen.Project, options *AwsCdkDepsOptions) error {
	return nil
}

