package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen/github/workflows"
)

// A set of utility functions for creating jobs in GitHub Workflows.
// Experimental.
type WorkflowJobs interface {
}

// The jsii proxy struct for WorkflowJobs
type jsiiProxy_WorkflowJobs struct {
	_ byte // padding
}

// Experimental.
func NewWorkflowJobs() WorkflowJobs {
	_init_.Initialize()

	j := jsiiProxy_WorkflowJobs{}

	_jsii_.Create(
		"projen.github.WorkflowJobs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWorkflowJobs_Override(w WorkflowJobs) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.WorkflowJobs",
		nil, // no parameters
		w,
	)
}

// Creates a pull request with the changes of a patch file.
//
// Returns: Job.
// Experimental.
func WorkflowJobs_PullRequestFromPatch(options *PullRequestFromPatchOptions) *workflows.Job {
	_init_.Initialize()

	if err := validateWorkflowJobs_PullRequestFromPatchParameters(options); err != nil {
		panic(err)
	}
	var returns *workflows.Job

	_jsii_.StaticInvoke(
		"projen.github.WorkflowJobs",
		"pullRequestFromPatch",
		[]interface{}{options},
		&returns,
	)

	return returns
}

