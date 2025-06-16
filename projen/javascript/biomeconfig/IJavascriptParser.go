package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the JavaScript parser behaves.
// Experimental.
type IJavascriptParser interface {
	// It enables the experimental and unsafe parsing of parameter decorators.
	//
	// These decorators belong to an old proposal, and they are subject to change.
	// Experimental.
	UnsafeParameterDecoratorsEnabled() *bool
	// Experimental.
	SetUnsafeParameterDecoratorsEnabled(u *bool)
}

// The jsii proxy for IJavascriptParser
type jsiiProxy_IJavascriptParser struct {
	_ byte // padding
}

func (j *jsiiProxy_IJavascriptParser) UnsafeParameterDecoratorsEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"unsafeParameterDecoratorsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptParser)SetUnsafeParameterDecoratorsEnabled(val *bool) {
	_jsii_.Set(
		j,
		"unsafeParameterDecoratorsEnabled",
		val,
	)
}

