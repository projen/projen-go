//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validatePullRequestBackport_IsComponentParameters(x interface{}) error {
	return nil
}

func validatePullRequestBackport_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPullRequestBackportParameters(scope constructs.IConstruct, options *PullRequestBackportOptions) error {
	return nil
}

