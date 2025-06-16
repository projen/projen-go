package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IHook interface {
	// The "position" of the closure function, starting from zero.
	//
	// For example, for React's `useEffect()` hook, the closure index is 0.
	// Experimental.
	ClosureIndex() *float64
	// Experimental.
	SetClosureIndex(c *float64)
	// The "position" of the array of dependencies, starting from zero.
	//
	// For example, for React's `useEffect()` hook, the dependencies index is 1.
	// Experimental.
	DependenciesIndex() *float64
	// Experimental.
	SetDependenciesIndex(d *float64)
	// The name of the hook.
	// Experimental.
	Name() *string
	// Experimental.
	SetName(n *string)
	// Whether the result of the hook is stable.
	//
	// Set to `true` to mark the identity of the hook's return value as stable, or use a number/an array of numbers to mark the "positions" in the return array as stable.
	//
	// For example, for React's `useRef()` hook the value would be `true`, while for `useState()` it would be `[1]`.
	// Experimental.
	StableResult() interface{}
	// Experimental.
	SetStableResult(s interface{})
}

// The jsii proxy for IHook
type jsiiProxy_IHook struct {
	_ byte // padding
}

func (j *jsiiProxy_IHook) ClosureIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"closureIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHook)SetClosureIndex(val *float64) {
	_jsii_.Set(
		j,
		"closureIndex",
		val,
	)
}

func (j *jsiiProxy_IHook) DependenciesIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dependenciesIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHook)SetDependenciesIndex(val *float64) {
	_jsii_.Set(
		j,
		"dependenciesIndex",
		val,
	)
}

func (j *jsiiProxy_IHook) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHook)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IHook) StableResult() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stableResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHook)SetStableResult(val interface{}) {
	if err := j.validateSetStableResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stableResult",
		val,
	)
}

