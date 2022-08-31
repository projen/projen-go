package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Adds a simple Typescript documentation generator.
// Experimental.
type TypedocDocgen interface {
}

// The jsii proxy struct for TypedocDocgen
type jsiiProxy_TypedocDocgen struct {
	_ byte // padding
}

// Experimental.
func NewTypedocDocgen(project TypeScriptProject) TypedocDocgen {
	_init_.Initialize()

	if err := validateNewTypedocDocgenParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_TypedocDocgen{}

	_jsii_.Create(
		"projen.typescript.TypedocDocgen",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewTypedocDocgen_Override(t TypedocDocgen, project TypeScriptProject) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.typescript.TypedocDocgen",
		[]interface{}{project},
		t,
	)
}

