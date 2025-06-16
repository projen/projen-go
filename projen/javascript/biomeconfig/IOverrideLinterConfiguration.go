package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IOverrideLinterConfiguration interface {
	// if `false`, it disables the feature and the linter won't be executed.
	//
	// `true` by default.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// List of rules.
	// Experimental.
	Rules() IRules
	// Experimental.
	SetRules(r IRules)
}

// The jsii proxy for IOverrideLinterConfiguration
type jsiiProxy_IOverrideLinterConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IOverrideLinterConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideLinterConfiguration)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IOverrideLinterConfiguration) Rules() IRules {
	var returns IRules
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideLinterConfiguration)SetRules(val IRules) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

