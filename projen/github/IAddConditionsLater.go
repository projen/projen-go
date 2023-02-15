package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IAddConditionsLater interface {
	// Experimental.
	Render() *[]*string
}

// The jsii proxy for IAddConditionsLater
type jsiiProxy_IAddConditionsLater struct {
	_ byte // padding
}

func (i *jsiiProxy_IAddConditionsLater) Render() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

