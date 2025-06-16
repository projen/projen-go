package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IAssistsConfiguration interface {
	// Whether Biome should fail in CLI if the assists were not applied to the code.
	// Experimental.
	Actions() IActions
	// Experimental.
	SetActions(a IActions)
	// Whether Biome should enable assists via LSP.
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

// The jsii proxy for IAssistsConfiguration
type jsiiProxy_IAssistsConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IAssistsConfiguration) Actions() IActions {
	var returns IActions
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssistsConfiguration)SetActions(val IActions) {
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_IAssistsConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssistsConfiguration)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IAssistsConfiguration) Ignore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssistsConfiguration)SetIgnore(val *[]*string) {
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_IAssistsConfiguration) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssistsConfiguration)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

