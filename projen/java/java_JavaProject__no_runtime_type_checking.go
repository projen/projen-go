//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package java

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JavaProject) validateAddDependencyParameters(spec *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateAddGitIgnoreParameters(pattern *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateAddPackageIgnoreParameters(_pattern *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateAddPluginParameters(spec *string, options *PluginOptions) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateAddTestDependencyParameters(spec *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateAddTipParameters(message *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateAnnotateGeneratedParameters(glob *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateRemoveTaskParameters(name *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateRunTaskCommandParameters(task projen.Task) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateTryFindFileParameters(filePath *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateTryFindJsonFileParameters(filePath *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateTryFindObjectFileParameters(filePath *string) error {
	return nil
}

func (j *jsiiProxy_JavaProject) validateTryRemoveFileParameters(filePath *string) error {
	return nil
}

func validateNewJavaProjectParameters(options *JavaProjectOptions) error {
	return nil
}

