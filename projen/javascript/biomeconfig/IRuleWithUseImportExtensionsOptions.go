package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithUseImportExtensionsOptions interface {
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
	Options() IUseImportExtensionsOptions
	// Experimental.
	SetOptions(o IUseImportExtensionsOptions)
}

// The jsii proxy for IRuleWithUseImportExtensionsOptions
type jsiiProxy_IRuleWithUseImportExtensionsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithUseImportExtensionsOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseImportExtensionsOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithUseImportExtensionsOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseImportExtensionsOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithUseImportExtensionsOptions) Options() IUseImportExtensionsOptions {
	var returns IUseImportExtensionsOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseImportExtensionsOptions)SetOptions(val IUseImportExtensionsOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

