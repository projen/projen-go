package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IAllowDomainOptions interface {
	// List of domains to allow `target="_blank"` without `rel="noreferrer"`.
	// Experimental.
	AllowDomains() *[]*string
	// Experimental.
	SetAllowDomains(a *[]*string)
}

// The jsii proxy for IAllowDomainOptions
type jsiiProxy_IAllowDomainOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IAllowDomainOptions) AllowDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAllowDomainOptions)SetAllowDomains(val *[]*string) {
	_jsii_.Set(
		j,
		"allowDomains",
		val,
	)
}

