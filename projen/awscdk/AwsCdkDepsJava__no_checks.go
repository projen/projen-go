//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCdkDepsJava) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (a *jsiiProxy_AwsCdkDepsJava) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateAwsCdkDepsJava_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAwsCdkDepsJava_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAwsCdkDepsJavaParameters(project projen.Project, options *AwsCdkDepsOptions) error {
	return nil
}

