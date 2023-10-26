//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RequirementsFile) validateSynthesizeContentParameters(resolver projen.IResolver) error {
	return nil
}

func validateRequirementsFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validateRequirementsFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_RequirementsFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_RequirementsFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewRequirementsFileParameters(project projen.Project, filePath *string, options *RequirementsFileOptions) error {
	return nil
}

