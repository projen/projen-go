package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Container for `TypescriptConfig` `tsconfig.json` base configuration(s). Extending from more than one base config file requires TypeScript 5.0+.
// Experimental.
type TypescriptConfigExtends interface {
	// Experimental.
	ToJSON() *[]*string
}

// The jsii proxy struct for TypescriptConfigExtends
type jsiiProxy_TypescriptConfigExtends struct {
	_ byte // padding
}

// Factory for creation from array of file paths.
// Experimental.
func TypescriptConfigExtends_FromPaths(paths *[]*string) TypescriptConfigExtends {
	_init_.Initialize()

	if err := validateTypescriptConfigExtends_FromPathsParameters(paths); err != nil {
		panic(err)
	}
	var returns TypescriptConfigExtends

	_jsii_.StaticInvoke(
		"projen.javascript.TypescriptConfigExtends",
		"fromPaths",
		[]interface{}{paths},
		&returns,
	)

	return returns
}

// Factory for creation from array of other `TypescriptConfig` instances.
// Experimental.
func TypescriptConfigExtends_FromTypescriptConfigs(configs *[]TypescriptConfig) TypescriptConfigExtends {
	_init_.Initialize()

	if err := validateTypescriptConfigExtends_FromTypescriptConfigsParameters(configs); err != nil {
		panic(err)
	}
	var returns TypescriptConfigExtends

	_jsii_.StaticInvoke(
		"projen.javascript.TypescriptConfigExtends",
		"fromTypescriptConfigs",
		[]interface{}{configs},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TypescriptConfigExtends) ToJSON() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

