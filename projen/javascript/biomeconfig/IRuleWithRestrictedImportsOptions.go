package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithRestrictedImportsOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IRestrictedImportsOptions
	// Experimental.
	SetOptions(o IRestrictedImportsOptions)
}

// The jsii proxy for IRuleWithRestrictedImportsOptions
type jsiiProxy_IRuleWithRestrictedImportsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithRestrictedImportsOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithRestrictedImportsOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithRestrictedImportsOptions) Options() IRestrictedImportsOptions {
	var returns IRestrictedImportsOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithRestrictedImportsOptions)SetOptions(val IRestrictedImportsOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

