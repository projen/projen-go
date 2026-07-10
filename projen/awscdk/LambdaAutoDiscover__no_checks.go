//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaAutoDiscover) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (l *jsiiProxy_LambdaAutoDiscover) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateLambdaAutoDiscover_IsComponentParameters(x interface{}) error {
	return nil
}

func validateLambdaAutoDiscover_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaAutoDiscoverParameters(project projen.Project, options *LambdaAutoDiscoverOptions) error {
	return nil
}

