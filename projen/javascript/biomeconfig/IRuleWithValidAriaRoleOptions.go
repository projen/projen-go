package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithValidAriaRoleOptions interface {
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
	Options() IValidAriaRoleOptions
	// Experimental.
	SetOptions(o IValidAriaRoleOptions)
}

// The jsii proxy for IRuleWithValidAriaRoleOptions
type jsiiProxy_IRuleWithValidAriaRoleOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithValidAriaRoleOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithValidAriaRoleOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithValidAriaRoleOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithValidAriaRoleOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithValidAriaRoleOptions) Options() IValidAriaRoleOptions {
	var returns IValidAriaRoleOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithValidAriaRoleOptions)SetOptions(val IValidAriaRoleOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

