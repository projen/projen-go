//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateIntegrationTest_IsComponentParameters(x interface{}) error {
	return nil
}

func validateIntegrationTest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegrationTestParameters(project projen.Project, options *IntegrationTestOptions) error {
	return nil
}

