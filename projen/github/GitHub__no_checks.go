//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitHub) validateAddDependabotParameters(options *DependabotOptions) error {
	return nil
}

func (g *jsiiProxy_GitHub) validateAddWorkflowParameters(name *string) error {
	return nil
}

func (g *jsiiProxy_GitHub) validateTryFindWorkflowParameters(name *string) error {
	return nil
}

func validateGitHub_OfParameters(project projen.Project) error {
	return nil
}

func validateNewGitHubParameters(project projen.Project, options *GitHubOptions) error {
	return nil
}

