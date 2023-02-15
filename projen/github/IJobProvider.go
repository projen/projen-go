package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen/github/workflows"
)

// Experimental.
type IJobProvider interface {
	// Generates a collection of named GitHub workflow jobs.
	// Experimental.
	RenderJobs() *map[string]*workflows.Job
}

// The jsii proxy for IJobProvider
type jsiiProxy_IJobProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_IJobProvider) RenderJobs() *map[string]*workflows.Job {
	var returns *map[string]*workflows.Job

	_jsii_.Invoke(
		i,
		"renderJobs",
		nil, // no parameters
		&returns,
	)

	return returns
}

