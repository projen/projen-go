//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PyprojectTomlFile) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (p *jsiiProxy_PyprojectTomlFile) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (p *jsiiProxy_PyprojectTomlFile) validateAddToArrayParameters(path *string) error {
	return nil
}

func (p *jsiiProxy_PyprojectTomlFile) validateSynthesizeContentParameters(resolver projen.IResolver) error {
	return nil
}

func validatePyprojectTomlFile_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePyprojectTomlFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PyprojectTomlFile) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_PyprojectTomlFile) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewPyprojectTomlFileParameters(scope constructs.IConstruct, config *PyprojectToml) error {
	return nil
}

