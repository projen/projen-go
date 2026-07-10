//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PythonSample) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (p *jsiiProxy_PythonSample) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validatePythonSample_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePythonSample_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPythonSampleParameters(project projen.Project, options *PythonSampleOptions) error {
	return nil
}

