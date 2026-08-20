//go:build no_runtime_type_checking

package cdktn

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CdktnConfig) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (c *jsiiProxy_CdktnConfig) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validateCdktnConfig_IsComponentParameters(x interface{}) error {
	return nil
}

func validateCdktnConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCdktnConfigParameters(scope constructs.IConstruct, options *CdktnConfigOptions) error {
	return nil
}

