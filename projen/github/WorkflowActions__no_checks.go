//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowActions_CheckoutWithPatchParameters(options *CheckoutWithPatchOptions) error {
	return nil
}

func validateWorkflowActions_CreatePullRequestParameters(options *CreatePullRequestOptions) error {
	return nil
}

func validateWorkflowActions_SetupGitIdentityParameters(id *GitIdentity) error {
	return nil
}

func validateWorkflowActions_UploadGitPatchParameters(options *UploadGitPatchOptions) error {
	return nil
}

