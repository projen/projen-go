//go:build no_runtime_type_checking

package uvconfig

// Building without runtime type checking enabled, so all the below just return nil

func validateExcludeNewerOverride_FromBooleanParameters(value *bool) error {
	return nil
}

func validateExcludeNewerOverride_FromStringParameters(value *string) error {
	return nil
}

