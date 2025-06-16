package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options applied to CSS files.
// Experimental.
type ICssConfiguration interface {
	// CSS assists options.
	// Experimental.
	Assists() ICssAssists
	// Experimental.
	SetAssists(a ICssAssists)
	// CSS formatter options.
	// Experimental.
	Formatter() ICssFormatter
	// Experimental.
	SetFormatter(f ICssFormatter)
	// CSS linter options.
	// Experimental.
	Linter() ICssLinter
	// Experimental.
	SetLinter(l ICssLinter)
	// CSS parsing options.
	// Experimental.
	Parser() ICssParser
	// Experimental.
	SetParser(p ICssParser)
}

// The jsii proxy for ICssConfiguration
type jsiiProxy_ICssConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_ICssConfiguration) Assists() ICssAssists {
	var returns ICssAssists
	_jsii_.Get(
		j,
		"assists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssConfiguration)SetAssists(val ICssAssists) {
	_jsii_.Set(
		j,
		"assists",
		val,
	)
}

func (j *jsiiProxy_ICssConfiguration) Formatter() ICssFormatter {
	var returns ICssFormatter
	_jsii_.Get(
		j,
		"formatter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssConfiguration)SetFormatter(val ICssFormatter) {
	_jsii_.Set(
		j,
		"formatter",
		val,
	)
}

func (j *jsiiProxy_ICssConfiguration) Linter() ICssLinter {
	var returns ICssLinter
	_jsii_.Get(
		j,
		"linter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssConfiguration)SetLinter(val ICssLinter) {
	_jsii_.Set(
		j,
		"linter",
		val,
	)
}

func (j *jsiiProxy_ICssConfiguration) Parser() ICssParser {
	var returns ICssParser
	_jsii_.Get(
		j,
		"parser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICssConfiguration)SetParser(val ICssParser) {
	_jsii_.Set(
		j,
		"parser",
		val,
	)
}

