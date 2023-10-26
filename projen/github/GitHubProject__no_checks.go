//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitHubProject) validateAddGitIgnoreParameters(pattern *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateAddPackageIgnoreParameters(_pattern *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateAddTaskParameters(name *string, props *projen.TaskOptions) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateAddTipParameters(message *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateAnnotateGeneratedParameters(glob *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateRemoveTaskParameters(name *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateRunTaskCommandParameters(task projen.Task) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateTryFindFileParameters(filePath *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateTryFindJsonFileParameters(filePath *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateTryFindObjectFileParameters(filePath *string) error {
	return nil
}

func (g *jsiiProxy_GitHubProject) validateTryRemoveFileParameters(filePath *string) error {
	return nil
}

func validateGitHubProject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGitHubProject_IsProjectParameters(x interface{}) error {
	return nil
}

func validateGitHubProject_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGitHubProjectParameters(options *GitHubProjectOptions) error {
	return nil
}

