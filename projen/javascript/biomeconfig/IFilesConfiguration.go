package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The configuration of the filesystem.
// Experimental.
type IFilesConfiguration interface {
	// A list of Unix shell style patterns.
	//
	// Biome will ignore files/folders that will match these patterns.
	// Experimental.
	Ignore() *[]*string
	// Experimental.
	SetIgnore(i *[]*string)
	// Tells Biome to not emit diagnostics when handling files that doesn't know.
	// Experimental.
	IgnoreUnknown() *bool
	// Experimental.
	SetIgnoreUnknown(i *bool)
	// A list of Unix shell style patterns.
	//
	// Biome will handle only those files/folders that will match these patterns.
	// Experimental.
	Include() *[]*string
	// Experimental.
	SetInclude(i *[]*string)
	// The maximum allowed size for source code files in bytes.
	//
	// Files above this limit will be ignored for performance reasons. Defaults to 1 MiB
	// Experimental.
	MaxSize() *float64
	// Experimental.
	SetMaxSize(m *float64)
}

// The jsii proxy for IFilesConfiguration
type jsiiProxy_IFilesConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IFilesConfiguration) Ignore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilesConfiguration)SetIgnore(val *[]*string) {
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_IFilesConfiguration) IgnoreUnknown() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignoreUnknown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilesConfiguration)SetIgnoreUnknown(val *bool) {
	_jsii_.Set(
		j,
		"ignoreUnknown",
		val,
	)
}

func (j *jsiiProxy_IFilesConfiguration) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilesConfiguration)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_IFilesConfiguration) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilesConfiguration)SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

