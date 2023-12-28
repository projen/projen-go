package typescript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// See: https://kulshekhar.github.io/ts-jest/docs/getting-started/options/babelConfig/
//
// Experimental.
type TsJestBabelConfig interface {
}

// The jsii proxy struct for TsJestBabelConfig
type jsiiProxy_TsJestBabelConfig struct {
	_ byte // padding
}

// Enables Babel processing.
//
// `ts-jest` will try to find an existing Babel configuration and pass it to the `babel-jest` processor.
// Experimental.
func TsJestBabelConfig_AutoDetectConfig() TsJestBabelConfig {
	_init_.Initialize()

	var returns TsJestBabelConfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestBabelConfig",
		"autoDetectConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Inline compiler options.
// See: https://babeljs.io/docs/options
//
// Experimental.
func TsJestBabelConfig_Custom(config *map[string]interface{}) TsJestBabelConfig {
	_init_.Initialize()

	if err := validateTsJestBabelConfig_CustomParameters(config); err != nil {
		panic(err)
	}
	var returns TsJestBabelConfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestBabelConfig",
		"custom",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Disables the use of Babel.
// Experimental.
func TsJestBabelConfig_Disabled() TsJestBabelConfig {
	_init_.Initialize()

	var returns TsJestBabelConfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestBabelConfig",
		"disabled",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Path to a babelrc file.
//
// The path should be relative to the current working directory where you start Jest from. You can also use `<rootDir>` in the path.
// Experimental.
func TsJestBabelConfig_FromFile(filePath *string) TsJestBabelConfig {
	_init_.Initialize()

	if err := validateTsJestBabelConfig_FromFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns TsJestBabelConfig

	_jsii_.StaticInvoke(
		"projen.typescript.TsJestBabelConfig",
		"fromFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

