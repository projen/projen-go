package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Linter options specific to the JSON linter.
// Experimental.
type IJsonLinter interface {
	// Control the linter for JSON (and its super languages) files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
}

// The jsii proxy for IJsonLinter
type jsiiProxy_IJsonLinter struct {
	_ byte // padding
}

func (j *jsiiProxy_IJsonLinter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonLinter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

