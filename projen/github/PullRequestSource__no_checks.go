//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validatePullRequestSource_FromBranchParameters(options *PullRequestSourceBranchOptions) error {
	return nil
}

func validatePullRequestSource_FromUsersParameters(options *PullRequestSourceUsersOptions) error {
	return nil
}

