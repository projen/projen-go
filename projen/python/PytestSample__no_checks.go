//go:build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PytestSample) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (p *jsiiProxy_PytestSample) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func validatePytestSample_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePytestSample_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPytestSampleParameters(project projen.Project, options *PytestSampleOptions) error {
	return nil
}

