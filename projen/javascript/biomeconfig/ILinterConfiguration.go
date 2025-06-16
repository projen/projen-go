package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ILinterConfiguration interface {
	// if `false`, it disables the feature and the linter won't be executed.
	//
	// `true` by default.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// A list of Unix shell style patterns.
	//
	// The formatter will ignore files/folders that will match these patterns.
	// Experimental.
	Ignore() *[]*string
	// Experimental.
	SetIgnore(i *[]*string)
	// A list of Unix shell style patterns.
	//
	// The formatter will include files/folders that will match these patterns.
	// Experimental.
	Include() *[]*string
	// Experimental.
	SetInclude(i *[]*string)
	// List of rules.
	// Experimental.
	Rules() IRules
	// Experimental.
	SetRules(r IRules)
}

// The jsii proxy for ILinterConfiguration
type jsiiProxy_ILinterConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_ILinterConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinterConfiguration)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ILinterConfiguration) Ignore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinterConfiguration)SetIgnore(val *[]*string) {
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_ILinterConfiguration) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinterConfiguration)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_ILinterConfiguration) Rules() IRules {
	var returns IRules
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILinterConfiguration)SetRules(val IRules) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

