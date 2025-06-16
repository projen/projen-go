package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRuleWithNoLabelWithoutControlOptions interface {
	// The severity of the emitted diagnostics by the rule.
	// Experimental.
	Level() *string
	// Experimental.
	SetLevel(l *string)
	// Rule's options.
	// Experimental.
	Options() INoLabelWithoutControlOptions
	// Experimental.
	SetOptions(o INoLabelWithoutControlOptions)
}

// The jsii proxy for IRuleWithNoLabelWithoutControlOptions
type jsiiProxy_IRuleWithNoLabelWithoutControlOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRuleWithNoLabelWithoutControlOptions) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoLabelWithoutControlOptions)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IRuleWithNoLabelWithoutControlOptions) Options() INoLabelWithoutControlOptions {
	var returns INoLabelWithoutControlOptions
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleWithNoLabelWithoutControlOptions)SetOptions(val INoLabelWithoutControlOptions) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

