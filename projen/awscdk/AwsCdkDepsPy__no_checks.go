//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCdkDepsPy) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (a *jsiiProxy_AwsCdkDepsPy) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateAwsCdkDepsPy_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAwsCdkDepsPy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAwsCdkDepsPyParameters(project projen.Project, options *AwsCdkDepsOptions) error {
	return nil
}

