//go:build no_runtime_type_checking

package typescript

// Building without runtime type checking enabled, so all the below just return nil

func validateTsJestDiagnostics_CustomParameters(config *map[string]interface{}) error {
	return nil
}

