//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func validateTypescriptConfigExtends_FromPathsParameters(paths *[]*string) error {
	return nil
}

func validateTypescriptConfigExtends_FromTypescriptConfigsParameters(configs *[]TypescriptConfig) error {
	return nil
}

