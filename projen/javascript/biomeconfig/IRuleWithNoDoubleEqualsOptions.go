package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithNoDoubleEqualsOptions interface {
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
	// Rule's options.
	// Experimental.
	Options() INoDoubleEqualsOptions
	// Experimental.
	SetOptions(o INoDoubleEqualsOptions)
}

// The jsii proxy for IRuleWithNoDoubleEqualsOptions
type jsiiProxy_IRuleWithNoDoubleEqualsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithNoDoubleEqualsOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoDoubleEqualsOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithNoDoubleEqualsOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoDoubleEqualsOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithNoDoubleEqualsOptions) Options() INoDoubleEqualsOptions {
	var returns INoDoubleEqualsOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoDoubleEqualsOptions)SetOptions(val INoDoubleEqualsOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

