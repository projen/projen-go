package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the CSS linter behaves.
// Experimental.
type ICssLinter interface {
	// Control the linter for CSS files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
}

// The jsii proxy for ICssLinter
type jsiiProxy_ICssLinter struct {
	_ byte // padding
}

func (j *jsiiProxy_ICssLinter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssLinter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

