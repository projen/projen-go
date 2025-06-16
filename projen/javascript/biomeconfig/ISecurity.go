package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type ISecurity interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Prevent the usage of dangerous JSX props.
	// Experimental.
	NoDangerouslySetInnerHtml() interface{}
	// Experimental.
	SetNoDangerouslySetInnerHtml(n interface{})
	// Report when a DOM element or a component uses both children and dangerouslySetInnerHTML prop.
	// Experimental.
	NoDangerouslySetInnerHtmlWithChildren() interface{}
	// Experimental.
	SetNoDangerouslySetInnerHtmlWithChildren(n interface{})
	// Disallow the use of global eval().
	// Experimental.
	NoGlobalEval() interface{}
	// Experimental.
	SetNoGlobalEval(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
}

// The jsii proxy for ISecurity
type jsiiProxy_ISecurity struct {
	_ byte // padding
}

func (j *jsiiProxy_ISecurity) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurity)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_ISecurity) NoDangerouslySetInnerHtml() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDangerouslySetInnerHtml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurity)SetNoDangerouslySetInnerHtml(val interface{}) {
	if err := j.validateSetNoDangerouslySetInnerHtmlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDangerouslySetInnerHtml",
		val,
	)
}

func (j *jsiiProxy_ISecurity) NoDangerouslySetInnerHtmlWithChildren() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDangerouslySetInnerHtmlWithChildren",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurity)SetNoDangerouslySetInnerHtmlWithChildren(val interface{}) {
	if err := j.validateSetNoDangerouslySetInnerHtmlWithChildrenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDangerouslySetInnerHtmlWithChildren",
		val,
	)
}

func (j *jsiiProxy_ISecurity) NoGlobalEval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noGlobalEval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurity)SetNoGlobalEval(val interface{}) {
	if err := j.validateSetNoGlobalEvalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noGlobalEval",
		val,
	)
}

func (j *jsiiProxy_ISecurity) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurity)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

