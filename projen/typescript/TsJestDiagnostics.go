package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// See: https://kulshekhar.github.io/ts-jest/docs/getting-started/options/diagnostics/
//
// Experimental.
type TsJestDiagnostics interface {
}

// The jsii proxy struct for TsJestDiagnostics
type jsiiProxy_TsJestDiagnostics struct {
	_ byte // padding
}

// Enable all diagnostics.
// Experimental.
func TsJestDiagnostics_All() TsJestDiagnostics {
	_init_.Initialize()

	var returns TsJestDiagnostics

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestDiagnostics",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Provide a custom diagnostics configuration.
// See: https://kulshekhar.github.io/ts-jest/docs/getting-started/options/diagnostics/
//
// Experimental.
func TsJestDiagnostics_Custom(config *map[string]interface{}) TsJestDiagnostics {
	_init_.Initialize()

	if err := validateTsJestDiagnostics_CustomParameters(config); err != nil {
		panic(err)
	}
	var returns TsJestDiagnostics

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestDiagnostics",
		"custom",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Disable all diagnostics.
// Experimental.
func TsJestDiagnostics_None() TsJestDiagnostics {
	_init_.Initialize()

	var returns TsJestDiagnostics

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestDiagnostics",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

