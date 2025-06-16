package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IOverrideOrganizeImportsConfiguration interface {
	// if `false`, it disables the feature and the linter won't be executed.
	//
	// `true` by default.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
}

// The jsii proxy for IOverrideOrganizeImportsConfiguration
type jsiiProxy_IOverrideOrganizeImportsConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IOverrideOrganizeImportsConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideOrganizeImportsConfiguration)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

