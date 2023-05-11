//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObjectFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (o *jsiiProxy_ObjectFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_ObjectFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (o *jsiiProxy_ObjectFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func (j *jsiiProxy_ObjectFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_ObjectFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewObjectFileParameters(project Project, filePath *string, options *ObjectFileOptions) error {
	return nil
}

