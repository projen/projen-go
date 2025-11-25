//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func validatePoetryPyproject_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePoetryPyproject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPoetryPyprojectParameters(scope constructs.IConstruct, options *PoetryPyprojectOptions) error {
	return nil
}

