package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options that changes how the GraphQL formatter behaves.
// Experimental.
type IGraphqlFormatter interface {
	// Whether to insert spaces around brackets in object literals.
	//
	// Defaults to true.
	// Experimental.
	BracketSpacing() *bool
	// Experimental.
	SetBracketSpacing(b *bool)
	// Control the formatter for GraphQL files.
	// Experimental.
	Enabled() *bool
	// Experimental.
	SetEnabled(e *bool)
	// The indent style applied to GraphQL files.
	// Experimental.
	IndentStyle() *string
	// Experimental.
	SetIndentStyle(i *string)
	// The size of the indentation applied to GraphQL files.
	//
	// Default to 2.
	// Experimental.
	IndentWidth() *float64
	// Experimental.
	SetIndentWidth(i *float64)
	// The type of line ending applied to GraphQL files.
	// Experimental.
	LineEnding() *string
	// Experimental.
	SetLineEnding(l *string)
	// What's the max width of a line applied to GraphQL files.
	//
	// Defaults to 80.
	// Experimental.
	LineWidth() *float64
	// Experimental.
	SetLineWidth(l *float64)
	// The type of quotes used in GraphQL code.
	//
	// Defaults to double.
	// Experimental.
	QuoteStyle() *string
	// Experimental.
	SetQuoteStyle(q *string)
}

// The jsii proxy for IGraphqlFormatter
type jsiiProxy_IGraphqlFormatter struct {
	_ byte // padding
}

func (j *jsiiProxy_IGraphqlFormatter) BracketSpacing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bracketSpacing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlFormatter)SetBracketSpacing(val *bool) {
	_jsii_.Set(
		j,
		"bracketSpacing",
		val,
	)
}

func (j *jsiiProxy_IGraphqlFormatter) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlFormatter)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IGraphqlFormatter) IndentStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlFormatter)SetIndentStyle(val *string) {
	_jsii_.Set(
		j,
		"indentStyle",
		val,
	)
}

func (j *jsiiProxy_IGraphqlFormatter) IndentWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indentWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlFormatter)SetIndentWidth(val *float64) {
	_jsii_.Set(
		j,
		"indentWidth",
		val,
	)
}

func (j *jsiiProxy_IGraphqlFormatter) LineEnding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lineEnding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlFormatter)SetLineEnding(val *string) {
	_jsii_.Set(
		j,
		"lineEnding",
		val,
	)
}

func (j *jsiiProxy_IGraphqlFormatter) LineWidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlFormatter)SetLineWidth(val *float64) {
	_jsii_.Set(
		j,
		"lineWidth",
		val,
	)
}

func (j *jsiiProxy_IGraphqlFormatter) QuoteStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlFormatter)SetQuoteStyle(val *string) {
	_jsii_.Set(
		j,
		"quoteStyle",
		val,
	)
}

