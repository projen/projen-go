//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package python

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PythonProject) validateAddDependencyParameters(spec *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateAddDevDependencyParameters(spec *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateAddGitIgnoreParameters(pattern *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateAddPackageIgnoreParameters(_pattern *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateAddTipParameters(message *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateAnnotateGeneratedParameters(glob *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateRemoveTaskParameters(name *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateRunTaskCommandParameters(task projen.Task) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateTryFindFileParameters(filePath *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateTryFindJsonFileParameters(filePath *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateTryFindObjectFileParameters(filePath *string) error {
	return nil
}

func (p *jsiiProxy_PythonProject) validateTryRemoveFileParameters(filePath *string) error {
	return nil
}

func validateNewPythonProjectParameters(options *PythonProjectOptions) error {
	return nil
}

