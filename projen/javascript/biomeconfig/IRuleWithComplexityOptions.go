package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithComplexityOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IComplexityOptions
	// Experimental.
	SetOptions(o IComplexityOptions)
}

// The jsii proxy for IRuleWithComplexityOptions
type jsiiProxy_IRuleWithComplexityOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithComplexityOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithComplexityOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithComplexityOptions) Options() IComplexityOptions {
	var returns IComplexityOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithComplexityOptions)SetOptions(val IComplexityOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

