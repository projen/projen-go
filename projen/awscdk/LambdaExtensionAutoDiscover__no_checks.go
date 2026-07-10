//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaExtensionAutoDiscover) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (l *jsiiProxy_LambdaExtensionAutoDiscover) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateLambdaExtensionAutoDiscover_IsComponentParameters(x interface{}) error {
	return nil
}

func validateLambdaExtensionAutoDiscover_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaExtensionAutoDiscoverParameters(project projen.Project, options *LambdaExtensionAutoDiscoverOptions) error {
	return nil
}

