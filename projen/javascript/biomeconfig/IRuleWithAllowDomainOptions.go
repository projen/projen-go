package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithAllowDomainOptions interface {
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
	Options() IAllowDomainOptions
	// Experimental.
	SetOptions(o IAllowDomainOptions)
}

// The jsii proxy for IRuleWithAllowDomainOptions
type jsiiProxy_IRuleWithAllowDomainOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithAllowDomainOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithAllowDomainOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithAllowDomainOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithAllowDomainOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithAllowDomainOptions) Options() IAllowDomainOptions {
	var returns IAllowDomainOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithAllowDomainOptions)SetOptions(val IAllowDomainOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

