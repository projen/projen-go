//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PullRequestTemplate) validateAddLineParameters(line *string) error {
	return nil
}

func (p *jsiiProxy_PullRequestTemplate) validateSynthesizeContentParameters(_arg projen.IResolver) error {
	return nil
}

func (j *jsiiProxy_PullRequestTemplate) validateSetExecutableParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_PullRequestTemplate) validateSetReadonlyParameters(val *bool) error {
	return nil
}

func validateNewPullRequestTemplateParameters(github GitHub, options *PullRequestTemplateOptions) error {
	return nil
}

