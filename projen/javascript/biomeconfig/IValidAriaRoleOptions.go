package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IValidAriaRoleOptions interface {
	// Experimental.
	AllowInvalidRoles() *[]*string
	// Experimental.
	SetAllowInvalidRoles(a *[]*string)
	// Experimental.
	IgnoreNonDom() *bool
	// Experimental.
	SetIgnoreNonDom(i *bool)
}

// The jsii proxy for IValidAriaRoleOptions
type jsiiProxy_IValidAriaRoleOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IValidAriaRoleOptions) AllowInvalidRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowInvalidRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IValidAriaRoleOptions)SetAllowInvalidRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"allowInvalidRoles",
		val,
	)
}

func (j *jsiiProxy_IValidAriaRoleOptions) IgnoreNonDom() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignoreNonDom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IValidAriaRoleOptions)SetIgnoreNonDom(val *bool) {
	_jsii_.Set(
		j,
		"ignoreNonDom",
		val,
	)
}

