package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithFixNoOptions interface {
	// The kind of the code actions emitted by the rule.
	// Experimental.
	Fix() *string
	// Experimental.
	SetFix(f *string)
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
}

// The jsii proxy for IRuleWithFixNoOptions
type jsiiProxy_IRuleWithFixNoOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithFixNoOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithFixNoOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithFixNoOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithFixNoOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

