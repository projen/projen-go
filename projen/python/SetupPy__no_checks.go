//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SetupPy) validateSynthesizeContentParameters(resolver projen.IResolver) error {
	return nil
}

func (j *jsiiProxy_SetupPy) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_SetupPy) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewSetupPyParameters(project projen.Project, options *SetupPyOptions) error {
	return nil
}

