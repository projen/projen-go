//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateAutoDiscover_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAutoDiscover_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAutoDiscoverParameters(project projen.Project, options *AutoDiscoverOptions) error {
	return nil
}

