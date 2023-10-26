//go:build no_runtime_type_checking

package build

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildJobParameters(id *string, job *workflows.Job) error {
	return nil
}

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildJobCommandsParameters(id *string, commands *[]*string, options *AddPostBuildJobCommandsOptions) error {
	return nil
}

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildJobTaskParameters(task projen.Task, options *AddPostBuildJobTaskOptions) error {
	return nil
}

func (b *jsiiProxy_BuildWorkflow) validateAddPostBuildStepsParameters(steps *[]*workflows.JobStep) error {
	return nil
}

func validateBuildWorkflow_IsComponentParameters(x interface{}) error {
	return nil
}

func validateBuildWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBuildWorkflowParameters(project projen.Project, options *BuildWorkflowOptions) error {
	return nil
}

