package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A set of options applied to the JavaScript files.
// Experimental.
type IJavascriptConfiguration interface {
	// Assists options.
	// Experimental.
	Assists() IJavascriptAssists
	// Experimental.
	SetAssists(a IJavascriptAssists)
	// Formatting options.
	// Experimental.
	Formatter() IJavascriptFormatter
	// Experimental.
	SetFormatter(f IJavascriptFormatter)
	// A list of global bindings that should be ignored by the analyzers.
	//
	// If defined here, they should not emit diagnostics.
	// Experimental.
	Globals() *[]*string
	// Experimental.
	SetGlobals(g *[]*string)
	// Indicates the type of runtime or transformation used for interpreting JSX.
	// Experimental.
	JsxRuntime() *string
	// Experimental.
	SetJsxRuntime(j *string)
	// Linter options.
	// Experimental.
	Linter() IJavascriptLinter
	// Experimental.
	SetLinter(l IJavascriptLinter)
	// Experimental.
	OrganizeImports() IJavascriptOrganizeImports
	// Experimental.
	SetOrganizeImports(o IJavascriptOrganizeImports)
	// Parsing options.
	// Experimental.
	Parser() IJavascriptParser
	// Experimental.
	SetParser(p IJavascriptParser)
}

// The jsii proxy for IJavascriptConfiguration
type jsiiProxy_IJavascriptConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_IJavascriptConfiguration) Assists() IJavascriptAssists {
	var returns IJavascriptAssists
	_jsii_.Get(
		j,
		"assists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptConfiguration)SetAssists(val IJavascriptAssists) {
	_jsii_.Set(
		j,
		"assists",
		val,
	)
}

func (j *jsiiProxy_IJavascriptConfiguration) Formatter() IJavascriptFormatter {
	var returns IJavascriptFormatter
	_jsii_.Get(
		j,
		"formatter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptConfiguration)SetFormatter(val IJavascriptFormatter) {
	_jsii_.Set(
		j,
		"formatter",
		val,
	)
}

func (j *jsiiProxy_IJavascriptConfiguration) Globals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"globals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptConfiguration)SetGlobals(val *[]*string) {
	_jsii_.Set(
		j,
		"globals",
		val,
	)
}

func (j *jsiiProxy_IJavascriptConfiguration) JsxRuntime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsxRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptConfiguration)SetJsxRuntime(val *string) {
	_jsii_.Set(
		j,
		"jsxRuntime",
		val,
	)
}

func (j *jsiiProxy_IJavascriptConfiguration) Linter() IJavascriptLinter {
	var returns IJavascriptLinter
	_jsii_.Get(
		j,
		"linter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptConfiguration)SetLinter(val IJavascriptLinter) {
	_jsii_.Set(
		j,
		"linter",
		val,
	)
}

func (j *jsiiProxy_IJavascriptConfiguration) OrganizeImports() IJavascriptOrganizeImports {
	var returns IJavascriptOrganizeImports
	_jsii_.Get(
		j,
		"organizeImports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptConfiguration)SetOrganizeImports(val IJavascriptOrganizeImports) {
	_jsii_.Set(
		j,
		"organizeImports",
		val,
	)
}

func (j *jsiiProxy_IJavascriptConfiguration) Parser() IJavascriptParser {
	var returns IJavascriptParser
	_jsii_.Get(
		j,
		"parser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJavascriptConfiguration)SetParser(val IJavascriptParser) {
	_jsii_.Set(
		j,
		"parser",
		val,
	)
}

