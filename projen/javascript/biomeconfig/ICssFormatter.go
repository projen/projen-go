package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the CSS formatter behaves.
// Experimental.
type ICssFormatter interface {
	// Control the formatter for CSS (and its super languages) files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// The indent style applied to CSS (and its super languages) files.
	// Experimental.
	IndentStyle() *string
	// Experimental.
	SetIndentStyle(i *string)
	// The size of the indentation applied to CSS (and its super languages) files.
	//
	// Default to 2.
	// Experimental.
	IndentWidth() *float64
	// Experimental.
	SetIndentWidth(i *float64)
	// The type of line ending applied to CSS (and its super languages) files.
	// Experimental.
	LineEnding() *string
	// Experimental.
	SetLineEnding(l *string)
	// What's the max width of a line applied to CSS (and its super languages) files.
	//
	// Defaults to 80.
	// Experimental.
	LineWidth() *float64
	// Experimental.
	SetLineWidth(l *float64)
	// The type of quotes used in CSS code.
	//
	// Defaults to double.
	// Experimental.
	QuoteStyle() *string
	// Experimental.
	SetQuoteStyle(q *string)
}

// The jsii proxy for ICssFormatter
type jsiiProxy_ICssFormatter struct {
	_ byte // padding
}

func (j *jsiiProxy_ICssFormatter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssFormatter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ICssFormatter) IndentStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssFormatter)SetIndentStyle(val *string) {
	_jsii_.Set(
		j,
		"indentStyle",
		val,
	)
}

func (j *jsiiProxy_ICssFormatter) IndentWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssFormatter)SetIndentWidth(val *float64) {
	_jsii_.Set(
		j,
		"indentWidth",
		val,
	)
}

func (j *jsiiProxy_ICssFormatter) LineEnding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineEnding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssFormatter)SetLineEnding(val *string) {
	_jsii_.Set(
		j,
		"lineEnding",
		val,
	)
}

func (j *jsiiProxy_ICssFormatter) LineWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssFormatter)SetLineWidth(val *float64) {
	_jsii_.Set(
		j,
		"lineWidth",
		val,
	)
}

func (j *jsiiProxy_ICssFormatter) QuoteStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssFormatter)SetQuoteStyle(val *string) {
	_jsii_.Set(
		j,
		"quoteStyle",
		val,
	)
}

