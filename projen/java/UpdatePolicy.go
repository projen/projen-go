package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Experimental.
type UpdatePolicy interface {
}

// The jsii proxy struct for UpdatePolicy
type jsiiProxy_UpdatePolicy struct {
	_ byte // padding
}

// Experimental.
func NewUpdatePolicy() UpdatePolicy {
	_init_.Initialize()

	j := jsiiProxy_UpdatePolicy{}

	_jsii_.Create(
		"projen.java.UpdatePolicy",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewUpdatePolicy_Override(u UpdatePolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.UpdatePolicy",
		nil, // no parameters
		u,
	)
}

// Updates at an interval of X minutes.
// Experimental.
func UpdatePolicy_Interval(minutes *float64) *string {
	_init_.Initialize()

	if err := validateUpdatePolicy_IntervalParameters(minutes); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"projen.java.UpdatePolicy",
		"interval",
		[]interface{}{minutes},
		&returns,
	)

	return returns
}

func UpdatePolicy_ALWAYS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.java.UpdatePolicy",
		"ALWAYS",
		&returns,
	)
	return returns
}

func UpdatePolicy_DAILY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.java.UpdatePolicy",
		"DAILY",
		&returns,
	)
	return returns
}

func UpdatePolicy_NEVER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.java.UpdatePolicy",
		"NEVER",
		&returns,
	)
	return returns
}

