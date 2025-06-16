package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IConsistentMemberAccessibilityOptions interface {
	// Experimental.
	Accessibility() *string
	// Experimental.
	SetAccessibility(a *string)
}

// The jsii proxy for IConsistentMemberAccessibilityOptions
type jsiiProxy_IConsistentMemberAccessibilityOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IConsistentMemberAccessibilityOptions) Accessibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsistentMemberAccessibilityOptions)SetAccessibility(val *string) {
	_jsii_.Set(
		j,
		"accessibility",
		val,
	)
}

