//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (y *jsiiProxy_YamlFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (y *jsiiProxy_YamlFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (y *jsiiProxy_YamlFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (y *jsiiProxy_YamlFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func validateYamlFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateYamlFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_YamlFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_YamlFile) validateSetLineWidthParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_YamlFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewYamlFileParameters(scope constructs.IConstruct, filePath *string, options *YamlFileOptions) error {
	return nil
}

