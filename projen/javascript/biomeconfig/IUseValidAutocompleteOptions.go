package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IUseValidAutocompleteOptions interface {
	// `input` like custom components that should be checked.
	// Experimental.
	InputComponents() *[]*string
	// Experimental.
	SetInputComponents(i *[]*string)
}

// The jsii proxy for IUseValidAutocompleteOptions
type jsiiProxy_IUseValidAutocompleteOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IUseValidAutocompleteOptions) InputComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUseValidAutocompleteOptions)SetInputComponents(val *[]*string) {
	_jsii_.Set(
		j,
		"inputComponents",
		val,
	)
}

