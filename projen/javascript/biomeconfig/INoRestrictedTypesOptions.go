package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type INoRestrictedTypesOptions interface {
	// Experimental.
	Types() *map[string]interface{}
	// Experimental.
	SetTypes(t *map[string]interface{})
}

// The jsii proxy for INoRestrictedTypesOptions
type jsiiProxy_INoRestrictedTypesOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_INoRestrictedTypesOptions) Types() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INoRestrictedTypesOptions)SetTypes(val *map[string]interface{}) {
	if err := j.validateSetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

