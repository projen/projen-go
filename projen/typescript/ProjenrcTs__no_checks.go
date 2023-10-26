//go:build no_runtime_type_checking

package typescript

// Building without runtime type checking enabled, so all the below just return nil

func validateProjenrcTs_IsComponentParameters(x interface{}) error {
	return nil
}

func validateProjenrcTs_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProjenrcTs_OfParameters(project projen.Project) error {
	return nil
}

func validateNewProjenrcTsParameters(project projen.Project, options *ProjenrcTsOptions) error {
	return nil
}

