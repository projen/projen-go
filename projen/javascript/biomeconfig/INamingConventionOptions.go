package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Rule's options.
// Experimental.
type INamingConventionOptions interface {
	// Custom conventions.
	// Experimental.
	Conventions() *[]IConvention
	// Experimental.
	SetConventions(c *[]IConvention)
	// Allowed cases for _TypeScript_ `enum` member names.
	// Experimental.
	EnumMemberCase() *string
	// Experimental.
	SetEnumMemberCase(e *string)
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

// The jsii proxy for INamingConventionOptions
type jsiiProxy_INamingConventionOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_INamingConventionOptions) Conventions() *[]IConvention {
	var returns *[]IConvention
	_jsii_.Get(
		j,
		"conventions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamingConventionOptions)SetConventions(val *[]IConvention) {
	_jsii_.Set(
		j,
		"conventions",
		val,
	)
}

func (j *jsiiProxy_INamingConventionOptions) EnumMemberCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enumMemberCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamingConventionOptions)SetEnumMemberCase(val *string) {
	_jsii_.Set(
		j,
		"enumMemberCase",
		val,
	)
}

func (j *jsiiProxy_INamingConventionOptions) RequireAscii() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"requireAscii",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamingConventionOptions)SetRequireAscii(val *bool) {
	_jsii_.Set(
		j,
		"requireAscii",
		val,
	)
}

func (j *jsiiProxy_INamingConventionOptions) StrictCase() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"strictCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamingConventionOptions)SetStrictCase(val *bool) {
	_jsii_.Set(
		j,
		"strictCase",
		val,
	)
}

