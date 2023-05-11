//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateProjenrcFile_OfParameters(project Project) error {
	return nil
}

func validateNewProjenrcFileParameters(project Project) error {
	return nil
}

