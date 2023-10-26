//go:build no_runtime_type_checking

package cdk

// Building without runtime type checking enabled, so all the below just return nil

func validateIntegrationTestBase_IsComponentParameters(x interface{}) error {
	return nil
}

func validateIntegrationTestBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegrationTestBaseParameters(project projen.Project, options *IntegrationTestBaseOptions) error {
	return nil
}

