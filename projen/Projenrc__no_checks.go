//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateProjenrc_OfParameters(project Project) error {
	return nil
}

func validateNewProjenrcParameters(project Project, options *ProjenrcJsonOptions) error {
	return nil
}

