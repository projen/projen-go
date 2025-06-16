package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithUtilityClassSortingOptions interface {
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
	Options() IUtilityClassSortingOptions
	// Experimental.
	SetOptions(o IUtilityClassSortingOptions)
}

// The jsii proxy for IRuleWithUtilityClassSortingOptions
type jsiiProxy_IRuleWithUtilityClassSortingOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithUtilityClassSortingOptions) Fix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUtilityClassSortingOptions)SetFix(val *string) {
	_jsii_.Set(
		j,
		"fix",
		val,
	)
}

func (j *jsiiProxy_IRuleWithUtilityClassSortingOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUtilityClassSortingOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithUtilityClassSortingOptions) Options() IUtilityClassSortingOptions {
	var returns IUtilityClassSortingOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithUtilityClassSortingOptions)SetOptions(val IUtilityClassSortingOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

