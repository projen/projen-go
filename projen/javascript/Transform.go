package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Experimental.
type Transform interface {
}

// The jsii proxy struct for Transform
type jsiiProxy_Transform struct {
	_ byte // padding
}

// Experimental.
func NewTransform(name *string, options interface{}) Transform {
	_init_.Initialize()

	if err := validateNewTransformParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_Transform{}

	_jsii_.Create(
		"projen.javascript.Transform",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTransform_Override(t Transform, name *string, options interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Transform",
		[]interface{}{name, options},
		t,
	)
}

