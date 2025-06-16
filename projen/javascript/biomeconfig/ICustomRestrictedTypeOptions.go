package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ICustomRestrictedTypeOptions interface {
	// Experimental.
	Message() *string
	// Experimental.
	SetMessage(m *string)
	// Experimental.
	Use() *string
	// Experimental.
	SetUse(u *string)
}

// The jsii proxy for ICustomRestrictedTypeOptions
type jsiiProxy_ICustomRestrictedTypeOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_ICustomRestrictedTypeOptions) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomRestrictedTypeOptions)SetMessage(val *string) {
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_ICustomRestrictedTypeOptions) Use() *string {
	var returns *string
	_jsii_.Get(
		j,
		"use",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomRestrictedTypeOptions)SetUse(val *string) {
	_jsii_.Set(
		j,
		"use",
		val,
	)
}

