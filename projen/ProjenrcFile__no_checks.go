//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateProjenrcFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateProjenrcFile_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProjenrcFile_OfParameters(project Project) error {
	return nil
}

func validateNewProjenrcFileParameters(scope constructs.IConstruct) error {
	return nil
}

