//go:build no_runtime_type_checking

package projen

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Gitpod) validateAddCustomTaskParameters(options *GitpodTask) error {
	return nil
}

func (g *jsiiProxy_Gitpod) validateAddDockerImageParameters(image DevEnvironmentDockerImage) error {
	return nil
}

func (g *jsiiProxy_Gitpod) validateAddPrebuildsParameters(config *GitpodPrebuilds) error {
	return nil
}

func validateGitpod_IsComponentParameters(x interface{}) error {
	return nil
}

func validateGitpod_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGitpodParameters(project Project, options *GitpodOptions) error {
	return nil
}

