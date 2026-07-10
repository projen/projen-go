//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaExtension) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (l *jsiiProxy_LambdaExtension) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateLambdaExtension_IsComponentParameters(x interface{}) error {
	return nil
}

func validateLambdaExtension_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLambdaExtensionParameters(project projen.Project, options *LambdaExtensionOptions) error {
	return nil
}

