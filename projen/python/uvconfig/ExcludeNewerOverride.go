package uvconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Experimental.
type ExcludeNewerOverride interface {
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for ExcludeNewerOverride
type jsiiProxy_ExcludeNewerOverride struct {
	_ byte // padding
}

func (j *jsiiProxy_ExcludeNewerOverride) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func ExcludeNewerOverride_FromBoolean(value *bool) ExcludeNewerOverride {
	_init_.Initialize()

	if err := validateExcludeNewerOverride_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns ExcludeNewerOverride

	_jsii_.StaticInvoke(
		"projen.python.uvConfig.ExcludeNewerOverride",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func ExcludeNewerOverride_FromString(value *string) ExcludeNewerOverride {
	_init_.Initialize()

	if err := validateExcludeNewerOverride_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ExcludeNewerOverride

	_jsii_.StaticInvoke(
		"projen.python.uvConfig.ExcludeNewerOverride",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

