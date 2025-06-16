package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type INoLabelWithoutControlOptions interface {
	// Array of component names that should be considered the same as an `input` element.
	// Experimental.
	InputComponents() *[]*string
	// Experimental.
	SetInputComponents(i *[]*string)
	// Array of attributes that should be treated as the `label` accessible text content.
	// Experimental.
	LabelAttributes() *[]*string
	// Experimental.
	SetLabelAttributes(l *[]*string)
	// Array of component names that should be considered the same as a `label` element.
	// Experimental.
	LabelComponents() *[]*string
	// Experimental.
	SetLabelComponents(l *[]*string)
}

// The jsii proxy for INoLabelWithoutControlOptions
type jsiiProxy_INoLabelWithoutControlOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_INoLabelWithoutControlOptions) InputComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INoLabelWithoutControlOptions)SetInputComponents(val *[]*string) {
	_jsii_.Set(
		j,
		"inputComponents",
		val,
	)
}

func (j *jsiiProxy_INoLabelWithoutControlOptions) LabelAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INoLabelWithoutControlOptions)SetLabelAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"labelAttributes",
		val,
	)
}

func (j *jsiiProxy_INoLabelWithoutControlOptions) LabelComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INoLabelWithoutControlOptions)SetLabelComponents(val *[]*string) {
	_jsii_.Set(
		j,
		"labelComponents",
		val,
	)
}

