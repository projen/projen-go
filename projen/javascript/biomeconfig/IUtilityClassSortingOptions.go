package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IUtilityClassSortingOptions interface {
	// Additional attributes that will be sorted.
	// Experimental.
	Attributes() *[]*string
	// Experimental.
	SetAttributes(a *[]*string)
	// Names of the functions or tagged templates that will be sorted.
	// Experimental.
	Functions() *[]*string
	// Experimental.
	SetFunctions(f *[]*string)
}

// The jsii proxy for IUtilityClassSortingOptions
type jsiiProxy_IUtilityClassSortingOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IUtilityClassSortingOptions) Attributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUtilityClassSortingOptions)SetAttributes(val *[]*string) {
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_IUtilityClassSortingOptions) Functions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUtilityClassSortingOptions)SetFunctions(val *[]*string) {
	_jsii_.Set(
		j,
		"functions",
		val,
	)
}

