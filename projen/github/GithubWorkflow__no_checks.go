//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GithubWorkflow) validateAddJobParameters(id *string, job interface{}) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateAddJobsParameters(jobs *map[string]interface{}) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateGetJobParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateOnParameters(events *workflows.Triggers) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateRemoveJobParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateUpdateJobParameters(id *string, job interface{}) error {
	return nil
}

func (g *jsiiProxy_GithubWorkflow) validateUpdateJobsParameters(jobs *map[string]interface{}) error {
	return nil
}

func validateNewGithubWorkflowParameters(github GitHub, name *string, options *GithubWorkflowOptions) error {
	return nil
}

