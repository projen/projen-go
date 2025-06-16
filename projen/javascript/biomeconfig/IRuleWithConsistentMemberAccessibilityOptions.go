package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithConsistentMemberAccessibilityOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IConsistentMemberAccessibilityOptions
	// Experimental.
	SetOptions(o IConsistentMemberAccessibilityOptions)
}

// The jsii proxy for IRuleWithConsistentMemberAccessibilityOptions
type jsiiProxy_IRuleWithConsistentMemberAccessibilityOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithConsistentMemberAccessibilityOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithConsistentMemberAccessibilityOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithConsistentMemberAccessibilityOptions) Options() IConsistentMemberAccessibilityOptions {
	var returns IConsistentMemberAccessibilityOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithConsistentMemberAccessibilityOptions)SetOptions(val IConsistentMemberAccessibilityOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

