package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithConsistentArrayTypeOptions interface {
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
	Options() IConsistentArrayTypeOptions
	// Experimental.
	SetOptions(o IConsistentArrayTypeOptions)
}

// The jsii proxy for IRuleWithConsistentArrayTypeOptions
type jsiiProxy_IRuleWithConsistentArrayTypeOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithConsistentArrayTypeOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithConsistentArrayTypeOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithConsistentArrayTypeOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithConsistentArrayTypeOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithConsistentArrayTypeOptions) Options() IConsistentArrayTypeOptions {
	var returns IConsistentArrayTypeOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithConsistentArrayTypeOptions)SetOptions(val IConsistentArrayTypeOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

