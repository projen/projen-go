//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Tasks) validateAddEnvironmentParameters(name *string, value *string) error {
	return nil
}

func (t *jsiiProxy_Tasks) validateAddTaskParameters(name *string, options *TaskOptions) error {
	return nil
}

func (t *jsiiProxy_Tasks) validateRemoveTaskParameters(name *string) error {
	return nil
}

func (t *jsiiProxy_Tasks) validateTryFindParameters(name *string) error {
	return nil
}

func validateNewTasksParameters(project Project) error {
	return nil
}

