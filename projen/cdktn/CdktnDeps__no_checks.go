//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdktnDeps) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (c *jsiiProxy_CdktnDeps) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateCdktnDeps_IsComponentParameters(x interface{}) error {
	return nil
}

func validateCdktnDeps_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdktnDepsParameters(project projen.Project, options *CdktnDepsOptions) error {
	return nil
}

