//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileBase) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func validateFileBase_IsComponentParameters(x interface{}) error {
	return nil
}

func validateFileBase_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_FileBase) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_FileBase) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewFileBaseParameters(scope constructs.IConstruct, filePath *string, options *FileBaseOptions) error {
	return nil
}

