//go:build !no_runtime_type_checking

package github

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen/github/workflows"
)

func validateWorkflowSteps_CheckoutParameters(options *CheckoutOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkflowSteps_SetupGitIdentityParameters(options *SetupGitIdentityOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkflowSteps_TagExistsParameters(tag *string, options *workflows.JobStepConfiguration) error {
	if tag == nil {
		return fmt.Errorf("parameter tag is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkflowSteps_UploadArtifactParameters(options *UploadArtifactOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

