package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IActions interface {
	// Experimental.
	Source() ISource
	// Experimental.
	SetSource(s ISource)
}

// The jsii proxy for IActions
type jsiiProxy_IActions struct {
	_ byte // padding
}

func (j *jsiiProxy_IActions) Source() ISource {
	var returns ISource
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActions)SetSource(val ISource) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

