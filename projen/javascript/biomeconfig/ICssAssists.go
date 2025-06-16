package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the CSS assists behaves.
// Experimental.
type ICssAssists interface {
	// Control the assists for CSS files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
}

// The jsii proxy for ICssAssists
type jsiiProxy_ICssAssists struct {
	_ byte // padding
}

func (j *jsiiProxy_ICssAssists) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssAssists)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

