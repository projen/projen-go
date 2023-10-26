//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func validatePytestSample_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePytestSample_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPytestSampleParameters(project projen.Project, options *PytestSampleOptions) error {
	return nil
}

