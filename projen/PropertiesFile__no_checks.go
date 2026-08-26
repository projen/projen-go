//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PropertiesFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (p *jsiiProxy_PropertiesFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PropertiesFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (p *jsiiProxy_PropertiesFile) validatePostProjectCreationParameters(initProject *InitProject) error {
	return nil
}

func (p *jsiiProxy_PropertiesFile) validateProjectCreationParameters(initProject *InitProject) error {
	return nil
}

func (p *jsiiProxy_PropertiesFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func validatePropertiesFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePropertiesFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PropertiesFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_PropertiesFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewPropertiesFileParameters(scope constructs.IConstruct, filePath *string, options *PropertiesFileOptions) error {
	return nil
}

