//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNewIntegrationTestBaseParameters(project projen.Project, options *IntegrationTestBaseOptions) error {
	return nil
}

