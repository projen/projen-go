//go:build no_runtime_type_checking

package github

// Building without runtime type checking enabled, so all the below just return nil

func validateWorkflowSteps_CheckoutParameters(options *CheckoutOptions) error {
	return nil
}

func validateWorkflowSteps_SetupGitIdentityParameters(options *SetupGitIdentityOptions) error {
	return nil
}

func validateWorkflowSteps_TagExistsParameters(tag *string, options *workflows.JobStepConfiguration) error {
	return nil
}

func validateWorkflowSteps_UploadArtifactParameters(options *UploadArtifactOptions) error {
	return nil
}

