//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func validateVersion_IsComponentParameters(x interface{}) error {
	return nil
}

func validateVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewVersionParameters(scope constructs.IConstruct, options *VersionOptions) error {
	return nil
}

