//go:build no_runtime_type_checking

package cdk8s

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationTest) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (i *jsiiProxy_IntegrationTest) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateIntegrationTest_IsComponentParameters(x interface{}) error {
	return nil
}

func validateIntegrationTest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewIntegrationTestParameters(project projen.Project, options *IntegrationTestOptions) error {
	return nil
}

