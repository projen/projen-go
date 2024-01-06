//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateTaskWorkflowJob_IsComponentParameters(x interface{}) error {
	return nil
}

func validateTaskWorkflowJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTaskWorkflowJobParameters(scope constructs.IConstruct, task projen.Task, options *TaskWorkflowJobOptions) error {
	return nil
}

