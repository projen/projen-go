//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationTestAutoDiscover) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (i *jsiiProxy_IntegrationTestAutoDiscover) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateIntegrationTestAutoDiscover_IsComponentParameters(x interface{}) error {
	return nil
}

func validateIntegrationTestAutoDiscover_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegrationTestAutoDiscoverParameters(project projen.Project, options *IntegrationTestAutoDiscoverOptions) error {
	return nil
}

