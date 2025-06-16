package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Linter options specific to the JSON linter.
// Experimental.
type IJsonAssists interface {
	// Control the linter for JSON (and its super languages) files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
}

// The jsii proxy for IJsonAssists
type jsiiProxy_IJsonAssists struct {
	_ byte // padding
}

func (j *jsiiProxy_IJsonAssists) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonAssists)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

