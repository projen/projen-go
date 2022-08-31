//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TomlFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (t *jsiiProxy_TomlFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (t *jsiiProxy_TomlFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (t *jsiiProxy_TomlFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func (j *jsiiProxy_TomlFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_TomlFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewTomlFileParameters(project Project, filePath *string, options *TomlFileOptions) error {
	return nil
}

