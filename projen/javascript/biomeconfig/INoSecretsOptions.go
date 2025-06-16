package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type INoSecretsOptions interface {
	// Set entropy threshold (default is 41).
	// Experimental.
	EntropyThreshold() *float64
	// Experimental.
	SetEntropyThreshold(e *float64)
}

// The jsii proxy for INoSecretsOptions
type jsiiProxy_INoSecretsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_INoSecretsOptions) EntropyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"entropyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INoSecretsOptions)SetEntropyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"entropyThreshold",
		val,
	)
}

