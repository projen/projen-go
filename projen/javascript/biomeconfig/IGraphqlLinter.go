package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the GraphQL linter behaves.
// Experimental.
type IGraphqlLinter interface {
	// Control the formatter for GraphQL files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
}

// The jsii proxy for IGraphqlLinter
type jsiiProxy_IGraphqlLinter struct {
	_ byte // padding
}

func (j *jsiiProxy_IGraphqlLinter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlLinter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

