//go:build !no_runtime_type_checking

package github

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateWorkflowActions_CheckoutWithPatchParameters(options *CheckoutWithPatchOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkflowActions_CreatePullRequestParameters(options *CreatePullRequestOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkflowActions_SetupGitIdentityParameters(id *GitIdentity) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(id, func() string { return "parameter id" }); err != nil {
		return err
	}

	return nil
}

func validateWorkflowActions_UploadGitPatchParameters(options *UploadGitPatchOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

