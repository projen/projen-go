package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Formatting options specific to the JavaScript files.
// Experimental.
type IJavascriptFormatter interface {
	// Whether to add non-necessary parentheses to arrow functions.
	//
	// Defaults to "always".
	// Experimental.
	ArrowParentheses() *string
	// Experimental.
	SetArrowParentheses(a *string)
	// The attribute position style in jsx elements.
	//
	// Defaults to auto.
	// Experimental.
	AttributePosition() *string
	// Experimental.
	SetAttributePosition(a *string)
	// Whether to hug the closing bracket of multiline HTML/JSX tags to the end of the last line, rather than being alone on the following line.
	//
	// Defaults to false.
	// Experimental.
	BracketSameLine() *bool
	// Experimental.
	SetBracketSameLine(b *bool)
	// Whether to insert spaces around brackets in object literals.
	//
	// Defaults to true.
	// Experimental.
	BracketSpacing() *bool
	// Experimental.
	SetBracketSpacing(b *bool)
	// Control the formatter for JavaScript (and its super languages) files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// The size of the indentation applied to JavaScript (and its super languages) files.
	//
	// Default to 2.
	// Experimental.
	IndentSize() *float64
	// Experimental.
	SetIndentSize(i *float64)
	// The indent style applied to JavaScript (and its super languages) files.
	// Experimental.
	IndentStyle() *string
	// Experimental.
	SetIndentStyle(i *string)
	// The size of the indentation applied to JavaScript (and its super languages) files.
	//
	// Default to 2.
	// Experimental.
	IndentWidth() *float64
	// Experimental.
	SetIndentWidth(i *float64)
	// The type of quotes used in JSX.
	//
	// Defaults to double.
	// Experimental.
	JsxQuoteStyle() *string
	// Experimental.
	SetJsxQuoteStyle(j *string)
	// The type of line ending applied to JavaScript (and its super languages) files.
	// Experimental.
	LineEnding() *string
	// Experimental.
	SetLineEnding(l *string)
	// What's the max width of a line applied to JavaScript (and its super languages) files.
	//
	// Defaults to 80.
	// Experimental.
	LineWidth() *float64
	// Experimental.
	SetLineWidth(l *float64)
	// When properties in objects are quoted.
	//
	// Defaults to asNeeded.
	// Experimental.
	QuoteProperties() *string
	// Experimental.
	SetQuoteProperties(q *string)
	// The type of quotes used in JavaScript code.
	//
	// Defaults to double.
	// Experimental.
	QuoteStyle() *string
	// Experimental.
	SetQuoteStyle(q *string)
	// Whether the formatter prints semicolons for all statements or only in for statements where it is necessary because of ASI.
	// Experimental.
	Semicolons() *string
	// Experimental.
	SetSemicolons(s *string)
	// Print trailing commas wherever possible in multi-line comma-separated syntactic structures.
	//
	// Defaults to "all".
	// Experimental.
	TrailingComma() *string
	// Experimental.
	SetTrailingComma(t *string)
	// Print trailing commas wherever possible in multi-line comma-separated syntactic structures.
	//
	// Defaults to "all".
	// Experimental.
	TrailingCommas() *string
	// Experimental.
	SetTrailingCommas(t *string)
}

// The jsii proxy for IJavascriptFormatter
type jsiiProxy_IJavascriptFormatter struct {
	_ byte // padding
}

func (j *jsiiProxy_IJavascriptFormatter) ArrowParentheses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arrowParentheses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetArrowParentheses(val *string) {
	_jsii_.Set(
		j,
		"arrowParentheses",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) AttributePosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributePosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetAttributePosition(val *string) {
	_jsii_.Set(
		j,
		"attributePosition",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) BracketSameLine() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bracketSameLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetBracketSameLine(val *bool) {
	_jsii_.Set(
		j,
		"bracketSameLine",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) BracketSpacing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bracketSpacing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetBracketSpacing(val *bool) {
	_jsii_.Set(
		j,
		"bracketSpacing",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) IndentSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetIndentSize(val *float64) {
	_jsii_.Set(
		j,
		"indentSize",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) IndentStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetIndentStyle(val *string) {
	_jsii_.Set(
		j,
		"indentStyle",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) IndentWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetIndentWidth(val *float64) {
	_jsii_.Set(
		j,
		"indentWidth",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) JsxQuoteStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsxQuoteStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetJsxQuoteStyle(val *string) {
	_jsii_.Set(
		j,
		"jsxQuoteStyle",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) LineEnding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineEnding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetLineEnding(val *string) {
	_jsii_.Set(
		j,
		"lineEnding",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) LineWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetLineWidth(val *float64) {
	_jsii_.Set(
		j,
		"lineWidth",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) QuoteProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetQuoteProperties(val *string) {
	_jsii_.Set(
		j,
		"quoteProperties",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) QuoteStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetQuoteStyle(val *string) {
	_jsii_.Set(
		j,
		"quoteStyle",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) Semicolons() *string {
	var returns *string
	_jsii_.Get(
		j,
		"semicolons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetSemicolons(val *string) {
	_jsii_.Set(
		j,
		"semicolons",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) TrailingComma() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailingComma",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetTrailingComma(val *string) {
	_jsii_.Set(
		j,
		"trailingComma",
		val,
	)
}

func (j *jsiiProxy_IJavascriptFormatter) TrailingCommas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailingCommas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptFormatter)SetTrailingCommas(val *string) {
	_jsii_.Set(
		j,
		"trailingCommas",
		val,
	)
}

