//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func validateProjenrc_OfParameters(project projen.Project) error {
	return nil
}

func validateNewProjenrcParameters(project projen.Project, options *ProjenrcOptions) error {
	return nil
}

