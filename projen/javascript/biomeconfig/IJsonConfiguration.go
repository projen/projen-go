package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options applied to JSON files.
// Experimental.
type IJsonConfiguration interface {
	// Assists options.
	// Experimental.
	Assists() IJsonAssists
	// Experimental.
	SetAssists(a IJsonAssists)
	// Formatting options.
	// Experimental.
	Formatter() IJsonFormatter
	// Experimental.
	SetFormatter(f IJsonFormatter)
	// Linting options.
	// Experimental.
	Linter() IJsonLinter
	// Experimental.
	SetLinter(l IJsonLinter)
	// Parsing options.
	// Experimental.
	Parser() IJsonParser
	// Experimental.
	SetParser(p IJsonParser)
}

// The jsii proxy for IJsonConfiguration
type jsiiProxy_IJsonConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IJsonConfiguration) Assists() IJsonAssists {
	var returns IJsonAssists
	_jsii_.Get(
		j,
		"assists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonConfiguration)SetAssists(val IJsonAssists) {
	_jsii_.Set(
		j,
		"assists",
		val,
	)
}

func (j *jsiiProxy_IJsonConfiguration) Formatter() IJsonFormatter {
	var returns IJsonFormatter
	_jsii_.Get(
		j,
		"formatter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonConfiguration)SetFormatter(val IJsonFormatter) {
	_jsii_.Set(
		j,
		"formatter",
		val,
	)
}

func (j *jsiiProxy_IJsonConfiguration) Linter() IJsonLinter {
	var returns IJsonLinter
	_jsii_.Get(
		j,
		"linter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonConfiguration)SetLinter(val IJsonLinter) {
	_jsii_.Set(
		j,
		"linter",
		val,
	)
}

func (j *jsiiProxy_IJsonConfiguration) Parser() IJsonParser {
	var returns IJsonParser
	_jsii_.Get(
		j,
		"parser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonConfiguration)SetParser(val IJsonParser) {
	_jsii_.Set(
		j,
		"parser",
		val,
	)
}

