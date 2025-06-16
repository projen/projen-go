package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithNoRestrictedTypesOptions interface {
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
	Options() INoRestrictedTypesOptions
	// Experimental.
	SetOptions(o INoRestrictedTypesOptions)
}

// The jsii proxy for IRuleWithNoRestrictedTypesOptions
type jsiiProxy_IRuleWithNoRestrictedTypesOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithNoRestrictedTypesOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoRestrictedTypesOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithNoRestrictedTypesOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoRestrictedTypesOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithNoRestrictedTypesOptions) Options() INoRestrictedTypesOptions {
	var returns INoRestrictedTypesOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoRestrictedTypesOptions)SetOptions(val INoRestrictedTypesOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

