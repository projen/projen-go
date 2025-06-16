package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Rule's options.
// Experimental.
type IFilenamingConventionOptions interface {
	// Allowed cases for file names.
	// Experimental.
	FilenameCases() *[]*string
	// Experimental.
	SetFilenameCases(f *[]*string)
	// If `false`, then non-ASCII characters are allowed.
	// Experimental.
	RequireAscii() *bool
	// Experimental.
	SetRequireAscii(r *bool)
	// If `false`, then consecutive uppercase are allowed in _camel_ and _pascal_ cases.
	//
	// This does not affect other [Case].
	// Experimental.
	StrictCase() *bool
	// Experimental.
	SetStrictCase(s *bool)
}

// The jsii proxy for IFilenamingConventionOptions
type jsiiProxy_IFilenamingConventionOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IFilenamingConventionOptions) FilenameCases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filenameCases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilenamingConventionOptions)SetFilenameCases(val *[]*string) {
	_jsii_.Set(
		j,
		"filenameCases",
		val,
	)
}

func (j *jsiiProxy_IFilenamingConventionOptions) RequireAscii() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"requireAscii",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilenamingConventionOptions)SetRequireAscii(val *bool) {
	_jsii_.Set(
		j,
		"requireAscii",
		val,
	)
}

func (j *jsiiProxy_IFilenamingConventionOptions) StrictCase() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"strictCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilenamingConventionOptions)SetStrictCase(val *bool) {
	_jsii_.Set(
		j,
		"strictCase",
		val,
	)
}

