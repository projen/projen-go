package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen/github/workflows"
)

// A set of utility functions for creating GitHub actions in workflows.
// Experimental.
type WorkflowActions interface {
}

// The jsii proxy struct for WorkflowActions
type jsiiProxy_WorkflowActions struct {
	_ byte // padding
}

// Experimental.
func NewWorkflowActions() WorkflowActions {
	_init_.Initialize()

	j := jsiiProxy_WorkflowActions{}

	_jsii_.Create(
		"projen.github.WorkflowActions",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWorkflowActions_Override(w WorkflowActions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.WorkflowActions",
		nil, // no parameters
		w,
	)
}

// Checks out a repository and applies a git patch that was created using `uploadGitPatch`.
//
// Returns: Job steps.
// Experimental.
func WorkflowActions_CheckoutWithPatch(options *CheckoutWithPatchOptions) *[]*workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowActions_CheckoutWithPatchParameters(options); err != nil {
		panic(err)
	}
	var returns *[]*workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowActions",
		"checkoutWithPatch",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// A step that creates a pull request based on the current repo state.
//
// Returns: Job steps.
// Experimental.
func WorkflowActions_CreatePullRequest(options *CreatePullRequestOptions) *[]*workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowActions_CreatePullRequestParameters(options); err != nil {
		panic(err)
	}
	var returns *[]*workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowActions",
		"createPullRequest",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Configures the git identity (user name and email).
//
// Returns: Job steps.
// Experimental.
func WorkflowActions_SetupGitIdentity(id *GitIdentity) *[]*workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowActions_SetupGitIdentityParameters(id); err != nil {
		panic(err)
	}
	var returns *[]*workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowActions",
		"setupGitIdentity",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Creates a .patch file from the current git diff and uploads it as an artifact. Use `checkoutWithPatch` to download and apply in another job.
//
// If a patch was uploaded, the action can optionally fail the job.
//
// Returns: Job steps.
// Experimental.
func WorkflowActions_UploadGitPatch(options *UploadGitPatchOptions) *[]*workflows.JobStep {
	_init_.Initialize()

	if err := validateWorkflowActions_UploadGitPatchParameters(options); err != nil {
		panic(err)
	}
	var returns *[]*workflows.JobStep

	_jsii_.StaticInvoke(
		"projen.github.WorkflowActions",
		"uploadGitPatch",
		[]interface{}{options},
		&returns,
	)

	return returns
}

