package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithNoSecretsOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() INoSecretsOptions
	// Experimental.
	SetOptions(o INoSecretsOptions)
}

// The jsii proxy for IRuleWithNoSecretsOptions
type jsiiProxy_IRuleWithNoSecretsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithNoSecretsOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoSecretsOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithNoSecretsOptions) Options() INoSecretsOptions {
	var returns INoSecretsOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoSecretsOptions)SetOptions(val INoSecretsOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

