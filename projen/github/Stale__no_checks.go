//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateStale_IsComponentParameters(x interface{}) error {
	return nil
}

func validateStale_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewStaleParameters(github GitHub, options *StaleOptions) error {
	return nil
}

