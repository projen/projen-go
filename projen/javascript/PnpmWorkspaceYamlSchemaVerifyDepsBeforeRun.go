package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// This setting allows the checking of the state of dependencies before running scripts.
// Experimental.
type PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun
type jsiiProxy_PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun struct {
	_ byte // padding
}

func (j *jsiiProxy_PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromBoolean(value *bool) PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromString(value *string) PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaVerifyDepsBeforeRun_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaVerifyDepsBeforeRun",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

