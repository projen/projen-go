package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// A comma-separated string of domain extensions that a proxy should not be used for, or a boolean.
//
// A null value leaves the setting unset. Since pnpm v12.23.0, any other type fails validation.
// Experimental.
type PnpmWorkspaceYamlSchemaNoProxy interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for PnpmWorkspaceYamlSchemaNoProxy
type jsiiProxy_PnpmWorkspaceYamlSchemaNoProxy struct {
	_ byte // padding
}

func (j *jsiiProxy_PnpmWorkspaceYamlSchemaNoProxy) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func PnpmWorkspaceYamlSchemaNoProxy_FromBoolean(value *bool) PnpmWorkspaceYamlSchemaNoProxy {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaNoProxy_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaNoProxy

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaNoProxy",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func PnpmWorkspaceYamlSchemaNoProxy_FromString(value *string) PnpmWorkspaceYamlSchemaNoProxy {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaNoProxy_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaNoProxy

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaNoProxy",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

