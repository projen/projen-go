package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ISelector interface {
	// Declaration kind.
	// Experimental.
	Kind() *string
	// Experimental.
	SetKind(k *string)
	// Modifiers used on the declaration.
	// Experimental.
	Modifiers() *[]*string
	// Experimental.
	SetModifiers(m *[]*string)
	// Scope of the declaration.
	// Experimental.
	Scope() *string
	// Experimental.
	SetScope(s *string)
}

// The jsii proxy for ISelector
type jsiiProxy_ISelector struct {
	_ byte // padding
}

func (j *jsiiProxy_ISelector) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISelector)SetKind(val *string) {
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_ISelector) Modifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"modifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISelector)SetModifiers(val *[]*string) {
	_jsii_.Set(
		j,
		"modifiers",
		val,
	)
}

func (j *jsiiProxy_ISelector) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISelector)SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

