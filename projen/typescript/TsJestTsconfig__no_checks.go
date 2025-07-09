//go:build no_runtime_type_checking

package typescript

// Building without runtime type checking enabled, so all the below just return nil

func validateTsJestTsconfig_CustomParameters(config *javascript.TypeScriptCompilerOptions) error {
	return nil
}

func validateTsJestTsconfig_FromFileParameters(filePath *string) error {
	return nil
}

