//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validatePullRequestLint_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePullRequestLint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPullRequestLintParameters(github GitHub, options *PullRequestLintOptions) error {
	return nil
}

