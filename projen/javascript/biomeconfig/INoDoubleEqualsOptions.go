package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Rule's options.
// Experimental.
type INoDoubleEqualsOptions interface {
	// If `true`, an exception is made when comparing with `null`, as it's often relied on to check both for `null` or `undefined`.
	//
	// If `false`, no such exception will be made.
	// Experimental.
	IgnoreNull() *bool
	// Experimental.
	SetIgnoreNull(i *bool)
}

// The jsii proxy for INoDoubleEqualsOptions
type jsiiProxy_INoDoubleEqualsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_INoDoubleEqualsOptions) IgnoreNull() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignoreNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INoDoubleEqualsOptions)SetIgnoreNull(val *bool) {
	_jsii_.Set(
		j,
		"ignoreNull",
		val,
	)
}

