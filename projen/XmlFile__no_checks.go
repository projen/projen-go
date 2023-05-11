//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (x *jsiiProxy_XmlFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (x *jsiiProxy_XmlFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (x *jsiiProxy_XmlFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (x *jsiiProxy_XmlFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func (j *jsiiProxy_XmlFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_XmlFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewXmlFileParameters(project Project, filePath *string, options *XmlFileOptions) error {
	return nil
}

