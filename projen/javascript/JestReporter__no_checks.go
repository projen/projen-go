//go:build no_runtime_type_checking

package javascript

// Building without runtime type checking enabled, so all the below just return nil

func validateNewJestReporterParameters(name *string) error {
	return nil
}

