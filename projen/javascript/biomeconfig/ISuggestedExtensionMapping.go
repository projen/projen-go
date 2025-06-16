package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ISuggestedExtensionMapping interface {
	// Extension that should be used for component file imports.
	// Experimental.
	Component() *string
	// Experimental.
	SetComponent(c *string)
	// Extension that should be used for module imports.
	// Experimental.
	Module() *string
	// Experimental.
	SetModule(m *string)
}

// The jsii proxy for ISuggestedExtensionMapping
type jsiiProxy_ISuggestedExtensionMapping struct {
	_ byte // padding
}

func (j *jsiiProxy_ISuggestedExtensionMapping) Component() *string {
	var returns *string
	_jsii_.Get(
		j,
		"component",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuggestedExtensionMapping)SetComponent(val *string) {
	_jsii_.Set(
		j,
		"component",
		val,
	)
}

func (j *jsiiProxy_ISuggestedExtensionMapping) Module() *string {
	var returns *string
	_jsii_.Get(
		j,
		"module",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuggestedExtensionMapping)SetModule(val *string) {
	_jsii_.Set(
		j,
		"module",
		val,
	)
}

