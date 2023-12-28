package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen/javascript"
)

// See: https://kulshekhar.github.io/ts-jest/docs/getting-started/options/tsconfig/
//
// Experimental.
type TsJestTsconfig interface {
}

// The jsii proxy struct for TsJestTsconfig
type jsiiProxy_TsJestTsconfig struct {
	_ byte // padding
}

// Uses `tsconfig.json` if found, or the built-in default TypeScript compiler options.
// Experimental.
func TsJestTsconfig_Auto() TsJestTsconfig {
	_init_.Initialize()

	var returns TsJestTsconfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestTsconfig",
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Force` ts-jest` to use its built-in defaults even if there is a `tsconfig.json` in your project.
// Experimental.
func TsJestTsconfig_BuiltInDefaults() TsJestTsconfig {
	_init_.Initialize()

	var returns TsJestTsconfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestTsconfig",
		"builtInDefaults",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Inline compiler options.
// See: TypescriptConfigOptions.
//
// Experimental.
func TsJestTsconfig_Custom(config *javascript.TypescriptConfigOptions) TsJestTsconfig {
	_init_.Initialize()

	if err := validateTsJestTsconfig_CustomParameters(config); err != nil {
		panic(err)
	}
	var returns TsJestTsconfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestTsconfig",
		"custom",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Path to a `tsconfig` file.
//
// The path should be relative to the current working directory where you start Jest from. You can also use `<rootDir>` in the path to start from the project root dir.
// Experimental.
func TsJestTsconfig_FromFile(filePath *string) TsJestTsconfig {
	_init_.Initialize()

	if err := validateTsJestTsconfig_FromFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns TsJestTsconfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestTsconfig",
		"fromFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

