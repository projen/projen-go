//go:build no_runtime_type_checking

package cdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNewIntegrationTestAutoDiscoverBaseParameters(project projen.Project, options *IntegrationTestAutoDiscoverBaseOptions) error {
	return nil
}

