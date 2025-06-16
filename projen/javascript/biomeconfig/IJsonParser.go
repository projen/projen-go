package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the JSON parser behaves.
// Experimental.
type IJsonParser interface {
	// Allow parsing comments in `.json` files.
	// Experimental.
	AllowComments() *bool
	// Experimental.
	SetAllowComments(a *bool)
	// Allow parsing trailing commas in `.json` files.
	// Experimental.
	AllowTrailingCommas() *bool
	// Experimental.
	SetAllowTrailingCommas(a *bool)
}

// The jsii proxy for IJsonParser
type jsiiProxy_IJsonParser struct {
	_ byte // padding
}

func (j *jsiiProxy_IJsonParser) AllowComments() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowComments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonParser)SetAllowComments(val *bool) {
	_jsii_.Set(
		j,
		"allowComments",
		val,
	)
}

func (j *jsiiProxy_IJsonParser) AllowTrailingCommas() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowTrailingCommas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonParser)SetAllowTrailingCommas(val *bool) {
	_jsii_.Set(
		j,
		"allowTrailingCommas",
		val,
	)
}

