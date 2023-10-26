//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskWorkflow) validateAddJobParameters(id *string, job interface{}) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateAddJobsParameters(jobs *map[string]interface{}) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateGetJobParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateOnParameters(events *workflows.Triggers) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateRemoveJobParameters(id *string) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateUpdateJobParameters(id *string, job interface{}) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateUpdateJobsParameters(jobs *map[string]interface{}) error {
	return nil
}

func validateTaskWorkflow_IsComponentParameters(x interface{}) error {
	return nil
}

func validateTaskWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTaskWorkflowParameters(github GitHub, options *TaskWorkflowOptions) error {
	return nil
}

