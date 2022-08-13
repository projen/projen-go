// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// API for resolving tokens when synthesizing file content.
// Experimental.
type IResolver interface {
	// Given a value (object/string/array/whatever, looks up any functions inside the object and returns an object where all functions are called.
	// Experimental.
	Resolve(value interface{}, options *ResolveOptions) interface{}
}

// The jsii proxy for IResolver
type jsiiProxy_IResolver struct {
	_ byte // padding
}

func (i *jsiiProxy_IResolver) Resolve(value interface{}, options *ResolveOptions) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{value, options},
		&returns,
	)

	return returns
}

