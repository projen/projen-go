package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithUseExhaustiveDependenciesOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() IUseExhaustiveDependenciesOptions
	// Experimental.
	SetOptions(o IUseExhaustiveDependenciesOptions)
}

// The jsii proxy for IRuleWithUseExhaustiveDependenciesOptions
type jsiiProxy_IRuleWithUseExhaustiveDependenciesOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithUseExhaustiveDependenciesOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseExhaustiveDependenciesOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithUseExhaustiveDependenciesOptions) Options() IUseExhaustiveDependenciesOptions {
	var returns IUseExhaustiveDependenciesOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUseExhaustiveDependenciesOptions)SetOptions(val IUseExhaustiveDependenciesOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

