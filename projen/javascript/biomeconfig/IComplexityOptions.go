package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for the rule `noExcessiveCognitiveComplexity`.
// Experimental.
type IComplexityOptions interface {
	// The maximum complexity score that we allow.
	//
	// Anything higher is considered excessive.
	// Experimental.
	MaxAllowedComplexity() *float64
	// Experimental.
	SetMaxAllowedComplexity(m *float64)
}

// The jsii proxy for IComplexityOptions
type jsiiProxy_IComplexityOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IComplexityOptions) MaxAllowedComplexity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedComplexity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComplexityOptions)SetMaxAllowedComplexity(val *float64) {
	_jsii_.Set(
		j,
		"maxAllowedComplexity",
		val,
	)
}

