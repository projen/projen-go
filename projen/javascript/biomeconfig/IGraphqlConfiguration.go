package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options applied to GraphQL files.
// Experimental.
type IGraphqlConfiguration interface {
	// GraphQL formatter options.
	// Experimental.
	Formatter() IGraphqlFormatter
	// Experimental.
	SetFormatter(f IGraphqlFormatter)
	// Experimental.
	Linter() IGraphqlLinter
	// Experimental.
	SetLinter(l IGraphqlLinter)
}

// The jsii proxy for IGraphqlConfiguration
type jsiiProxy_IGraphqlConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IGraphqlConfiguration) Formatter() IGraphqlFormatter {
	var returns IGraphqlFormatter
	_jsii_.Get(
		j,
		"formatter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlConfiguration)SetFormatter(val IGraphqlFormatter) {
	_jsii_.Set(
		j,
		"formatter",
		val,
	)
}

func (j *jsiiProxy_IGraphqlConfiguration) Linter() IGraphqlLinter {
	var returns IGraphqlLinter
	_jsii_.Get(
		j,
		"linter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlConfiguration)SetLinter(val IGraphqlLinter) {
	_jsii_.Set(
		j,
		"linter",
		val,
	)
}

