//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TextFile) validateAddLineParameters(line *string) error {
	return nil
}

func (t *jsiiProxy_TextFile) validateSynthesizeContentParameters(_arg IResolver) error {
	return nil
}

func validateTextFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateTextFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TextFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_TextFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewTextFileParameters(scope constructs.IConstruct, filePath *string, options *TextFileOptions) error {
	return nil
}

