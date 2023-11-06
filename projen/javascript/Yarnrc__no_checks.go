//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func validateYarnrc_IsComponentParameters(x interface{}) error {
	return nil
}

func validateYarnrc_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewYarnrcParameters(project projen.Project, version *string, options *YarnrcOptions) error {
	return nil
}

