package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithRestrictedGlobalsOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IRestrictedGlobalsOptions
	// Experimental.
	SetOptions(o IRestrictedGlobalsOptions)
}

// The jsii proxy for IRuleWithRestrictedGlobalsOptions
type jsiiProxy_IRuleWithRestrictedGlobalsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithRestrictedGlobalsOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithRestrictedGlobalsOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithRestrictedGlobalsOptions) Options() IRestrictedGlobalsOptions {
	var returns IRestrictedGlobalsOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithRestrictedGlobalsOptions)SetOptions(val IRestrictedGlobalsOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

