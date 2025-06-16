package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IOrganizeImports interface {
	// Enables the organization of imports.
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
}

// The jsii proxy for IOrganizeImports
type jsiiProxy_IOrganizeImports struct {
	_ byte // padding
}

func (j *jsiiProxy_IOrganizeImports) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizeImports)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IOrganizeImports) Ignore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizeImports)SetIgnore(val *[]*string) {
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_IOrganizeImports) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOrganizeImports)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

