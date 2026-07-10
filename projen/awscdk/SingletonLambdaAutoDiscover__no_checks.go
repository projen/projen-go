//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SingletonLambdaAutoDiscover) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (s *jsiiProxy_SingletonLambdaAutoDiscover) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateSingletonLambdaAutoDiscover_IsComponentParameters(x interface{}) error {
	return nil
}

func validateSingletonLambdaAutoDiscover_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSingletonLambdaAutoDiscoverParameters(project projen.Project, options *SingletonLambdaAutoDiscoverOptions) error {
	return nil
}

