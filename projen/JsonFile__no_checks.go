//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JsonFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (j *jsiiProxy_JsonFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (j *jsiiProxy_JsonFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (j *jsiiProxy_JsonFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func validateJsonFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateJsonFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_JsonFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_JsonFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewJsonFileParameters(scope constructs.IConstruct, filePath *string, options *JsonFileOptions) error {
	return nil
}

