package cdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Creates a markdown file based on the jsii manifest: - Adds a `docgen` script to package.json - Runs `jsii-docgen` after compilation - Enforces that markdown file is checked in.
// Experimental.
type JsiiDocgen interface {
}

// The jsii proxy struct for JsiiDocgen
type jsiiProxy_JsiiDocgen struct {
	_ byte // padding
}

// Experimental.
func NewJsiiDocgen(project JsiiProject, options *JsiiDocgenOptions) JsiiDocgen {
	_init_.Initialize()

	if err := validateNewJsiiDocgenParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_JsiiDocgen{}

	_jsii_.Create(
		"projen.cdk.JsiiDocgen",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJsiiDocgen_Override(j JsiiDocgen, project JsiiProject, options *JsiiDocgenOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk.JsiiDocgen",
		[]interface{}{project, options},
		j,
	)
}

