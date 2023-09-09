package gitlab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// id_tokens Definition.
// See: https://docs.gitlab.com/ee/ci/yaml/#id_tokens
//
// Experimental.
type IDToken interface {
	// The required aud sub-keyword is used to configure the aud claim for the JWT.
	// Experimental.
	Aud() interface{}
	// Experimental.
	SetAud(a interface{})
}

// The jsii proxy for IDToken
type jsiiProxy_IDToken struct {
	_ byte // padding
}

func (j *jsiiProxy_IDToken) Aud() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDToken)SetAud(val interface{}) {
	if err := j.validateSetAudParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aud",
		val,
	)
}

