package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Generic options applied to all files.
// Experimental.
type IFormatterConfiguration interface {
	// The attribute position style in HTMLish languages.
	//
	// By default auto.
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
	// Use any `.editorconfig` files to configure the formatter. Configuration in `biome.json` will override `.editorconfig` configuration. Default: false.
	// Experimental.
	UseEditorconfig() *bool
	// Experimental.
	SetUseEditorconfig(u *bool)
}

// The jsii proxy for IFormatterConfiguration
type jsiiProxy_IFormatterConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IFormatterConfiguration) AttributePosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributePosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetAttributePosition(val *string) {
	_jsii_.Set(
		j,
		"attributePosition",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) BracketSpacing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bracketSpacing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetBracketSpacing(val *bool) {
	_jsii_.Set(
		j,
		"bracketSpacing",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) FormatWithErrors() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"formatWithErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetFormatWithErrors(val *bool) {
	_jsii_.Set(
		j,
		"formatWithErrors",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) Ignore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetIgnore(val *[]*string) {
	_jsii_.Set(
		j,
		"ignore",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) IndentSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetIndentSize(val *float64) {
	_jsii_.Set(
		j,
		"indentSize",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) IndentStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetIndentStyle(val *string) {
	_jsii_.Set(
		j,
		"indentStyle",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) IndentWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetIndentWidth(val *float64) {
	_jsii_.Set(
		j,
		"indentWidth",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) LineEnding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineEnding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetLineEnding(val *string) {
	_jsii_.Set(
		j,
		"lineEnding",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) LineWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetLineWidth(val *float64) {
	_jsii_.Set(
		j,
		"lineWidth",
		val,
	)
}

func (j *jsiiProxy_IFormatterConfiguration) UseEditorconfig() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"useEditorconfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFormatterConfiguration)SetUseEditorconfig(val *bool) {
	_jsii_.Set(
		j,
		"useEditorconfig",
		val,
	)
}

