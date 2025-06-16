package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithFilenamingConventionOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IFilenamingConventionOptions
	// Experimental.
	SetOptions(o IFilenamingConventionOptions)
}

// The jsii proxy for IRuleWithFilenamingConventionOptions
type jsiiProxy_IRuleWithFilenamingConventionOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithFilenamingConventionOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithFilenamingConventionOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithFilenamingConventionOptions) Options() IFilenamingConventionOptions {
	var returns IFilenamingConventionOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithFilenamingConventionOptions)SetOptions(val IFilenamingConventionOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

