//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateComponent_IsComponentParameters(x interface{}) error {
	return nil
}

func validateComponent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewComponentParameters(scope constructs.IConstruct) error {
	return nil
}

