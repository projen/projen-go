//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IniFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (i *jsiiProxy_IniFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_IniFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (i *jsiiProxy_IniFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func (j *jsiiProxy_IniFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_IniFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewIniFileParameters(project Project, filePath *string, options *IniFileOptions) error {
	return nil
}

