//go:build no_runtime_type_checking

// CDK for software projects
package projen

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Task) validateBuiltinParameters(name *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateEnvParameters(name *string, value *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateExecParameters(command *string, options *TaskStepOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validatePrependParameters(shell *string, options *TaskStepOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validatePrependExecParameters(shell *string, options *TaskStepOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validatePrependSayParameters(message *string, options *TaskStepOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validatePrependSpawnParameters(subtask Task, options *TaskStepOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateResetParameters(options *TaskStepOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateSayParameters(message *string, options *TaskStepOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateSpawnParameters(subtask Task, options *TaskStepOptions) error {
	return nil
}

func validateNewTaskParameters(name *string, props *TaskOptions) error {
	return nil
}

