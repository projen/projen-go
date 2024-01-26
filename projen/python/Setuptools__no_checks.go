//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func validateSetuptools_IsComponentParameters(x interface{}) error {
	return nil
}

func validateSetuptools_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSetuptoolsParameters(project projen.Project, options *SetuptoolsOptions) error {
	return nil
}

