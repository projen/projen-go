package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// If this is enabled, locally available packages are linked to node_modules instead of being downloaded from the registry.
// Experimental.
type PnpmWorkspaceYamlSchemaLinkWorkspacePackages interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for PnpmWorkspaceYamlSchemaLinkWorkspacePackages
type jsiiProxy_PnpmWorkspaceYamlSchemaLinkWorkspacePackages struct {
	_ byte // padding
}

func (j *jsiiProxy_PnpmWorkspaceYamlSchemaLinkWorkspacePackages) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func PnpmWorkspaceYamlSchemaLinkWorkspacePackages_FromBoolean(value *bool) PnpmWorkspaceYamlSchemaLinkWorkspacePackages {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaLinkWorkspacePackages_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaLinkWorkspacePackages

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaLinkWorkspacePackages",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func PnpmWorkspaceYamlSchemaLinkWorkspacePackages_FromString(value *string) PnpmWorkspaceYamlSchemaLinkWorkspacePackages {
	_init_.Initialize()

	if err := validatePnpmWorkspaceYamlSchemaLinkWorkspacePackages_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns PnpmWorkspaceYamlSchemaLinkWorkspacePackages

	_jsii_.StaticInvoke(
		"projen.javascript.PnpmWorkspaceYamlSchemaLinkWorkspacePackages",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

