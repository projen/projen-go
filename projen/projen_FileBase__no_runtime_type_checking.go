//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileBase) validateSynthesizeContentParameters(resolver IResolver) error {
	return nil
}

func (j *jsiiProxy_FileBase) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_FileBase) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewFileBaseParameters(project Project, filePath *string, options *FileBaseOptions) error {
	return nil
}

