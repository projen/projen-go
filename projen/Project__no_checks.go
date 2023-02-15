//go:build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Project) validateAddGitIgnoreParameters(pattern *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateAddPackageIgnoreParameters(_pattern *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateAddTaskParameters(name *string, props *TaskOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateAddTipParameters(message *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateAnnotateGeneratedParameters(_glob *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateRemoveTaskParameters(name *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateRunTaskCommandParameters(task Task) error {
	return nil
}

func (p *jsiiProxy_Project) validateTryFindFileParameters(filePath *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateTryFindJsonFileParameters(filePath *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateTryFindObjectFileParameters(filePath *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateTryRemoveFileParameters(filePath *string) error {
	return nil
}

func validateNewProjectParameters(options *ProjectOptions) error {
	return nil
}

