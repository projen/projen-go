package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Options for the rule `noRestrictedImports`.
// Experimental.
type IRestrictedImportsOptions interface {
	// A list of names that should trigger the rule.
	// Experimental.
	Paths() *map[string]*string
	// Experimental.
	SetPaths(p *map[string]*string)
}

// The jsii proxy for IRestrictedImportsOptions
type jsiiProxy_IRestrictedImportsOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_IRestrictedImportsOptions) Paths() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"paths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRestrictedImportsOptions)SetPaths(val *map[string]*string) {
	_jsii_.Set(
		j,
		"paths",
		val,
	)
}

