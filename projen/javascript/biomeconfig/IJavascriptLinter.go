package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Linter options specific to the JavaScript linter.
// Experimental.
type IJavascriptLinter interface {
	// Control the linter for JavaScript (and its super languages) files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
}

// The jsii proxy for IJavascriptLinter
type jsiiProxy_IJavascriptLinter struct {
	_ byte // padding
}

func (j *jsiiProxy_IJavascriptLinter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptLinter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

