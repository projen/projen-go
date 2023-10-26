//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func validateVenv_IsComponentParameters(x interface{}) error {
	return nil
}

func validateVenv_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewVenvParameters(project projen.Project, options *VenvOptions) error {
	return nil
}

