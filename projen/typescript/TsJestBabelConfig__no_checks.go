//go:build no_runtime_type_checking

package typescript

// Building without runtime type checking enabled, so all the below just return nil

func validateTsJestBabelConfig_CustomParameters(config *map[string]interface{}) error {
	return nil
}

func validateTsJestBabelConfig_FromFileParameters(filePath *string) error {
	return nil
}

