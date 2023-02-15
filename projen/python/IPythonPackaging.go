package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/projen/projen-go/projen"
)

// Experimental.
type IPythonPackaging interface {
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
}

// The jsii proxy for IPythonPackaging
type jsiiProxy_IPythonPackaging struct {
	_ byte // padding
}

func (j *jsiiProxy_IPythonPackaging) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

