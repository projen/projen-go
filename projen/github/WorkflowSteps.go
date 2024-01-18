package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen/github/workflows"
)

// A collection of very commonly used, individual, GitHub Workflow Job steps.
// Experimental.
type WorkflowSteps interface {
}

// The jsii proxy struct for WorkflowSteps
type jsiiProxy_WorkflowSteps struct {
	_ byte // padding
}

// Experimental.
func NewWorkflowSteps() WorkflowSteps {
	_init_.Initialize()

	j := jsiiProxy_WorkflowSteps{}

	_jsii_.Create(
		"projen.github.WorkflowSteps",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWorkflowSteps_Override(w WorkflowSteps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.WorkflowSteps",
		nil, // no parameters
		w,
	)
}

// Checks out a repository.
//
// Returns: A JobStep that checks out a repository.
// Experimental.
func WorkflowSteps_Checkout(options *CheckoutOptions) *workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowSteps_CheckoutParameters(options); err != nil {
		panic(err)
	}
	var returns *workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowSteps",
		"checkout",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Configures the git identity (user name and email).
//
// Returns: Job step that configures the provided git identity.
// Experimental.
func WorkflowSteps_SetupGitIdentity(options *SetupGitIdentityOptions) *workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowSteps_SetupGitIdentityParameters(options); err != nil {
		panic(err)
	}
	var returns *workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowSteps",
		"setupGitIdentity",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Checks if a tag exists.
//
// Requires a checkout step to have been run before this step with "fetch-depth" set to "0".
//
// Outputs:
// - `exists`: A string value of 'true' or 'false' indicating if the tag exists.
//
// Returns: Job step that checks if the provided tag exists.
// Experimental.
func WorkflowSteps_TagExists(tag *string, options *workflows.JobStepConfiguration) *workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowSteps_TagExistsParameters(tag, options); err != nil {
		panic(err)
	}
	var returns *workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowSteps",
		"tagExists",
		[]interface{}{tag, options},
		&returns,
	)

	return returns
}

// Uploads an artifact.
//
// Returns: A JobStep that uploads an artifact.
// Experimental.
func WorkflowSteps_UploadArtifact(options *UploadArtifactOptions) *workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowSteps_UploadArtifactParameters(options); err != nil {
		panic(err)
	}
	var returns *workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowSteps",
		"uploadArtifact",
		[]interface{}{options},
		&returns,
	)

	return returns
}

