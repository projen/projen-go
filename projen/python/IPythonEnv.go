package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IPythonEnv interface {
	// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
	// Experimental.
	SetupEnvironment()
}

// The jsii proxy for IPythonEnv
type jsiiProxy_IPythonEnv struct {
	_ byte // padding
}

func (i *jsiiProxy_IPythonEnv) SetupEnvironment() {
	_jsii_.InvokeVoid(
		i,
		"setupEnvironment",
		nil, // no parameters
	)
}

