package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IConvention interface {
	// String cases to enforce.
	// Experimental.
	Formats() *[]*string
	// Experimental.
	SetFormats(f *[]*string)
	// Regular expression to enforce.
	// Experimental.
	Match() *string
	// Experimental.
	SetMatch(m *string)
	// Declarations concerned by this convention.
	// Experimental.
	Selector() ISelector
	// Experimental.
	SetSelector(s ISelector)
}

// The jsii proxy for IConvention
type jsiiProxy_IConvention struct {
	_ byte // padding
}

func (j *jsiiProxy_IConvention) Formats() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"formats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConvention)SetFormats(val *[]*string) {
	_jsii_.Set(
		j,
		"formats",
		val,
	)
}

func (j *jsiiProxy_IConvention) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConvention)SetMatch(val *string) {
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_IConvention) Selector() ISelector {
	var returns ISelector
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConvention)SetSelector(val ISelector) {
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

