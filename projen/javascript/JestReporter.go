package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Experimental.
type JestReporter interface {
}

// The jsii proxy struct for JestReporter
type jsiiProxy_JestReporter struct {
	_ byte // padding
}

// Experimental.
func NewJestReporter(name *string, options *map[string]interface{}) JestReporter {
	_init_.Initialize()

	if err := validateNewJestReporterParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_JestReporter{}

	_jsii_.Create(
		"projen.javascript.JestReporter",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJestReporter_Override(j JestReporter, name *string, options *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.JestReporter",
		[]interface{}{name, options},
		j,
	)
}

