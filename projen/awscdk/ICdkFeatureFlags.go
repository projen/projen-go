package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type ICdkFeatureFlags interface {
	// Experimental.
	Flags() *map[string]interface{}
}

// The jsii proxy for ICdkFeatureFlags
type jsiiProxy_ICdkFeatureFlags struct {
	_ byte // padding
}

func (j *jsiiProxy_ICdkFeatureFlags) Flags() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}

