//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskWorkflow) validateAddJobParameters(id *string, job interface{}) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateAddJobsParameters(jobs *map[string]interface{}) error {
	return nil
}

func (t *jsiiProxy_TaskWorkflow) validateOnParameters(events *workflows.Triggers) error {
	return nil
}

func validateNewTaskWorkflowParameters(github GitHub, options *TaskWorkflowOptions) error {
	return nil
}

