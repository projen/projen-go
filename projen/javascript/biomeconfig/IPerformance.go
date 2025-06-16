package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A list of rules that belong to this group.
// Experimental.
type IPerformance interface {
	// It enables ALL rules for this group.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Disallow the use of spread (...) syntax on accumulators.
	// Experimental.
	NoAccumulatingSpread() interface{}
	// Experimental.
	SetNoAccumulatingSpread(n interface{})
	// Disallow the use of barrel file.
	// Experimental.
	NoBarrelFile() interface{}
	// Experimental.
	SetNoBarrelFile(n interface{})
	// Disallow the use of the delete operator.
	// Experimental.
	NoDelete() interface{}
	// Experimental.
	SetNoDelete(n interface{})
	// Avoid re-export all.
	// Experimental.
	NoReExportAll() interface{}
	// Experimental.
	SetNoReExportAll(n interface{})
	// It enables the recommended rules for this group.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Require regex literals to be declared at the top level.
	// Experimental.
	UseTopLevelRegex() interface{}
	// Experimental.
	SetUseTopLevelRegex(u interface{})
}

// The jsii proxy for IPerformance
type jsiiProxy_IPerformance struct {
	_ byte // padding
}

func (j *jsiiProxy_IPerformance) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformance)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_IPerformance) NoAccumulatingSpread() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noAccumulatingSpread",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformance)SetNoAccumulatingSpread(val interface{}) {
	if err := j.validateSetNoAccumulatingSpreadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noAccumulatingSpread",
		val,
	)
}

func (j *jsiiProxy_IPerformance) NoBarrelFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noBarrelFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformance)SetNoBarrelFile(val interface{}) {
	if err := j.validateSetNoBarrelFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noBarrelFile",
		val,
	)
}

func (j *jsiiProxy_IPerformance) NoDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformance)SetNoDelete(val interface{}) {
	if err := j.validateSetNoDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noDelete",
		val,
	)
}

func (j *jsiiProxy_IPerformance) NoReExportAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noReExportAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformance)SetNoReExportAll(val interface{}) {
	if err := j.validateSetNoReExportAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noReExportAll",
		val,
	)
}

func (j *jsiiProxy_IPerformance) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformance)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_IPerformance) UseTopLevelRegex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTopLevelRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformance)SetUseTopLevelRegex(val interface{}) {
	if err := j.validateSetUseTopLevelRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTopLevelRegex",
		val,
	)
}

