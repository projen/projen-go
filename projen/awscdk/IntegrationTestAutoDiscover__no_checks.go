//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateIntegrationTestAutoDiscover_IsComponentParameters(x interface{}) error {
	return nil
}

func validateIntegrationTestAutoDiscover_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegrationTestAutoDiscoverParameters(project projen.Project, options *IntegrationTestAutoDiscoverOptions) error {
	return nil
}

