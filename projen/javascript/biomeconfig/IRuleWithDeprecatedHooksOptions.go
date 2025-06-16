package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithDeprecatedHooksOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IDeprecatedHooksOptions
	// Experimental.
	SetOptions(o IDeprecatedHooksOptions)
}

// The jsii proxy for IRuleWithDeprecatedHooksOptions
type jsiiProxy_IRuleWithDeprecatedHooksOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithDeprecatedHooksOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithDeprecatedHooksOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithDeprecatedHooksOptions) Options() IDeprecatedHooksOptions {
	var returns IDeprecatedHooksOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithDeprecatedHooksOptions)SetOptions(val IDeprecatedHooksOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

