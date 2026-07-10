//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GithubWorkflow) validateAddJobParameters(id *string, job interface{}) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateAddJobsParameters(jobs *map[string]interface{}) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateAppendStepParameters(jobId *string, step *workflows.JobStep) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateGetJobParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateGetStepParameters(jobId *string, stepId *string) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateInsertStepAfterParameters(jobId *string, referenceStepId *string, step *workflows.JobStep) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateInsertStepBeforeParameters(jobId *string, referenceStepId *string, step *workflows.JobStep) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateOnParameters(events *workflows.Triggers) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validatePatchStepParameters(jobId *string, stepId *string, patch *workflows.JobStep) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validatePostProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateProjectCreationParameters(initProject *projen.InitProject) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateRemoveJobParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateRemoveStepParameters(jobId *string, stepId *string) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateReplaceStepParameters(jobId *string, stepId *string, replacementStep *workflows.JobStep) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateUpdateJobParameters(id *string, job interface{}) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateUpdateJobsParameters(jobs *map[string]interface{}) error {
	return nil
}

func validateGithubWorkflow_IsComponentParameters(x interface{}) error {
	return nil
}

func validateGithubWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGithubWorkflowParameters(github GitHub, name *string, options *GithubWorkflowOptions) error {
	return nil
}

