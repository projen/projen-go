package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IOverrideFormatterConfiguration interface {
	// The attribute position style.
	// Experimental.
	AttributePosition() *string
	// Experimental.
	SetAttributePosition(a *string)
	// Whether to insert spaces around brackets in object literals.
	//
	// Defaults to true.
	// Experimental.
	BracketSpacing() *bool
	// Experimental.
	SetBracketSpacing(b *bool)
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// Stores whether formatting should be allowed to proceed if a given file has syntax errors.
	// Experimental.
	FormatWithErrors() *bool
	// Experimental.
	SetFormatWithErrors(f *bool)
	// The size of the indentation, 2 by default (deprecated, use `indent-width`).
	// Experimental.
	IndentSize() *float64
	// Experimental.
	SetIndentSize(i *float64)
	// The indent style.
	// Experimental.
	IndentStyle() *string
	// Experimental.
	SetIndentStyle(i *string)
	// The size of the indentation, 2 by default.
	// Experimental.
	IndentWidth() *float64
	// Experimental.
	SetIndentWidth(i *float64)
	// The type of line ending.
	// Experimental.
	LineEnding() *string
	// Experimental.
	SetLineEnding(l *string)
	// What's the max width of a line.
	//
	// Defaults to 80.
	// Experimental.
	LineWidth() *float64
	// Experimental.
	SetLineWidth(l *float64)
}

// The jsii proxy for IOverrideFormatterConfiguration
type jsiiProxy_IOverrideFormatterConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) AttributePosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributePosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetAttributePosition(val *string) {
	_jsii_.Set(
		j,
		"attributePosition",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) BracketSpacing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bracketSpacing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetBracketSpacing(val *bool) {
	_jsii_.Set(
		j,
		"bracketSpacing",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) FormatWithErrors() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"formatWithErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetFormatWithErrors(val *bool) {
	_jsii_.Set(
		j,
		"formatWithErrors",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) IndentSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetIndentSize(val *float64) {
	_jsii_.Set(
		j,
		"indentSize",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) IndentStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetIndentStyle(val *string) {
	_jsii_.Set(
		j,
		"indentStyle",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) IndentWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetIndentWidth(val *float64) {
	_jsii_.Set(
		j,
		"indentWidth",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) LineEnding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineEnding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetLineEnding(val *string) {
	_jsii_.Set(
		j,
		"lineEnding",
		val,
	)
}

func (j *jsiiProxy_IOverrideFormatterConfiguration) LineWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOverrideFormatterConfiguration)SetLineWidth(val *float64) {
	_jsii_.Set(
		j,
		"lineWidth",
		val,
	)
}

