//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateProjenrcJson_IsComponentParameters(x interface{}) error {
	return nil
}

func validateProjenrcJson_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProjenrcJson_OfParameters(project Project) error {
	return nil
}

func validateNewProjenrcJsonParameters(project Project, options *ProjenrcJsonOptions) error {
	return nil
}

