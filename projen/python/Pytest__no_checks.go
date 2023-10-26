//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func validatePytest_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePytest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPytestParameters(project projen.Project, options *PytestOptions) error {
	return nil
}

