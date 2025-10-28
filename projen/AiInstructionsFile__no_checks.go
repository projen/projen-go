//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AiInstructionsFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func validateAiInstructionsFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateAiInstructionsFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AiInstructionsFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_AiInstructionsFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewAiInstructionsFileParameters(scope constructs.IConstruct, filePath *string, options *FileBaseOptions) error {
	return nil
}

