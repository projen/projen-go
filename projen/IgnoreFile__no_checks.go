//go:build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IgnoreFile) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func (j *jsiiProxy_IgnoreFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_IgnoreFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewIgnoreFileParameters(project Project, filePath *string) error {
	return nil
}

