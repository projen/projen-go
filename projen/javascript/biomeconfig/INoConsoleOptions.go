package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type INoConsoleOptions interface {
	// Allowed calls on the console object.
	// Experimental.
	Allow() *[]*string
	// Experimental.
	SetAllow(a *[]*string)
}

// The jsii proxy for INoConsoleOptions
type jsiiProxy_INoConsoleOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_INoConsoleOptions) Allow() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INoConsoleOptions)SetAllow(val *[]*string) {
	if err := j.validateSetAllowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allow",
		val,
	)
}

