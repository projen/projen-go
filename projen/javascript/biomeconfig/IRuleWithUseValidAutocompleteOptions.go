package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithUseValidAutocompleteOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IUseValidAutocompleteOptions
	// Experimental.
	SetOptions(o IUseValidAutocompleteOptions)
}

// The jsii proxy for IRuleWithUseValidAutocompleteOptions
type jsiiProxy_IRuleWithUseValidAutocompleteOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithUseValidAutocompleteOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseValidAutocompleteOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithUseValidAutocompleteOptions) Options() IUseValidAutocompleteOptions {
	var returns IUseValidAutocompleteOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseValidAutocompleteOptions)SetOptions(val IUseValidAutocompleteOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

