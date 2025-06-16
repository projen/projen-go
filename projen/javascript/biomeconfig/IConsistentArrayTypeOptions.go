package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IConsistentArrayTypeOptions interface {
	// Experimental.
	Syntax() *string
	// Experimental.
	SetSyntax(s *string)
}

// The jsii proxy for IConsistentArrayTypeOptions
type jsiiProxy_IConsistentArrayTypeOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IConsistentArrayTypeOptions) Syntax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syntax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsistentArrayTypeOptions)SetSyntax(val *string) {
	_jsii_.Set(
		j,
		"syntax",
		val,
	)
}

