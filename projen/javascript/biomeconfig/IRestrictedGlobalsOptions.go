package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for the rule `noRestrictedGlobals`.
// Experimental.
type IRestrictedGlobalsOptions interface {
	// A list of names that should trigger the rule.
	// Experimental.
	DeniedGlobals() *[]*string
	// Experimental.
	SetDeniedGlobals(d *[]*string)
}

// The jsii proxy for IRestrictedGlobalsOptions
type jsiiProxy_IRestrictedGlobalsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRestrictedGlobalsOptions) DeniedGlobals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deniedGlobals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestrictedGlobalsOptions)SetDeniedGlobals(val *[]*string) {
	_jsii_.Set(
		j,
		"deniedGlobals",
		val,
	)
}

