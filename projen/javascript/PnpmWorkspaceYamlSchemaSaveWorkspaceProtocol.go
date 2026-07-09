package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// This setting controls how dependencies that are linked from the workspace are added to package.json.
// Experimental.
type PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol
type jsiiProxy_PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol_FromBoolean(value *bool) PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaSaveWorkspaceProtocol_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol_FromString(value *string) PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaSaveWorkspaceProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaSaveWorkspaceProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

