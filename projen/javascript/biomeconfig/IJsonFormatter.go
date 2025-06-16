package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IJsonFormatter interface {
	// Control the formatter for JSON (and its super languages) files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// The size of the indentation applied to JSON (and its super languages) files.
	//
	// Default to 2.
	// Experimental.
	IndentSize() *float64
	// Experimental.
	SetIndentSize(i *float64)
	// The indent style applied to JSON (and its super languages) files.
	// Experimental.
	IndentStyle() *string
	// Experimental.
	SetIndentStyle(i *string)
	// The size of the indentation applied to JSON (and its super languages) files.
	//
	// Default to 2.
	// Experimental.
	IndentWidth() *float64
	// Experimental.
	SetIndentWidth(i *float64)
	// The type of line ending applied to JSON (and its super languages) files.
	// Experimental.
	LineEnding() *string
	// Experimental.
	SetLineEnding(l *string)
	// What's the max width of a line applied to JSON (and its super languages) files.
	//
	// Defaults to 80.
	// Experimental.
	LineWidth() *float64
	// Experimental.
	SetLineWidth(l *float64)
	// Print trailing commas wherever possible in multi-line comma-separated syntactic structures.
	//
	// Defaults to "none".
	// Experimental.
	TrailingCommas() *string
	// Experimental.
	SetTrailingCommas(t *string)
}

// The jsii proxy for IJsonFormatter
type jsiiProxy_IJsonFormatter struct {
	_ byte // padding
}

func (j *jsiiProxy_IJsonFormatter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonFormatter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IJsonFormatter) IndentSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonFormatter)SetIndentSize(val *float64) {
	_jsii_.Set(
		j,
		"indentSize",
		val,
	)
}

func (j *jsiiProxy_IJsonFormatter) IndentStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonFormatter)SetIndentStyle(val *string) {
	_jsii_.Set(
		j,
		"indentStyle",
		val,
	)
}

func (j *jsiiProxy_IJsonFormatter) IndentWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonFormatter)SetIndentWidth(val *float64) {
	_jsii_.Set(
		j,
		"indentWidth",
		val,
	)
}

func (j *jsiiProxy_IJsonFormatter) LineEnding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineEnding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonFormatter)SetLineEnding(val *string) {
	_jsii_.Set(
		j,
		"lineEnding",
		val,
	)
}

func (j *jsiiProxy_IJsonFormatter) LineWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonFormatter)SetLineWidth(val *float64) {
	_jsii_.Set(
		j,
		"lineWidth",
		val,
	)
}

func (j *jsiiProxy_IJsonFormatter) TrailingCommas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailingCommas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonFormatter)SetTrailingCommas(val *string) {
	_jsii_.Set(
		j,
		"trailingCommas",
		val,
	)
}

