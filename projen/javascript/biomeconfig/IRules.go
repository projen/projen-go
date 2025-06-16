package biomeconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IRules interface {
	// Experimental.
	A11y() IA11y
	// Experimental.
	SetA11y(a IA11y)
	// It enables ALL rules.
	//
	// The rules that belong to `nursery` won't be enabled.
	// Experimental.
	All() *bool
	// Experimental.
	SetAll(a *bool)
	// Experimental.
	Complexity() IComplexity
	// Experimental.
	SetComplexity(c IComplexity)
	// Experimental.
	Correctness() ICorrectness
	// Experimental.
	SetCorrectness(c ICorrectness)
	// Experimental.
	Nursery() INursery
	// Experimental.
	SetNursery(n INursery)
	// Experimental.
	Performance() IPerformance
	// Experimental.
	SetPerformance(p IPerformance)
	// It enables the lint rules recommended by Biome.
	//
	// `true` by default.
	// Experimental.
	Recommended() *bool
	// Experimental.
	SetRecommended(r *bool)
	// Experimental.
	Security() ISecurity
	// Experimental.
	SetSecurity(s ISecurity)
	// Experimental.
	Style() IStyle
	// Experimental.
	SetStyle(s IStyle)
	// Experimental.
	Suspicious() ISuspicious
	// Experimental.
	SetSuspicious(s ISuspicious)
}

// The jsii proxy for IRules
type jsiiProxy_IRules struct {
	_ byte // padding
}

func (j *jsiiProxy_IRules) A11y() IA11y {
	var returns IA11y
	_jsii_.Get(
		j,
		"a11y",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetA11y(val IA11y) {
	_jsii_.Set(
		j,
		"a11y",
		val,
	)
}

func (j *jsiiProxy_IRules) All() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetAll(val *bool) {
	_jsii_.Set(
		j,
		"all",
		val,
	)
}

func (j *jsiiProxy_IRules) Complexity() IComplexity {
	var returns IComplexity
	_jsii_.Get(
		j,
		"complexity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetComplexity(val IComplexity) {
	_jsii_.Set(
		j,
		"complexity",
		val,
	)
}

func (j *jsiiProxy_IRules) Correctness() ICorrectness {
	var returns ICorrectness
	_jsii_.Get(
		j,
		"correctness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetCorrectness(val ICorrectness) {
	_jsii_.Set(
		j,
		"correctness",
		val,
	)
}

func (j *jsiiProxy_IRules) Nursery() INursery {
	var returns INursery
	_jsii_.Get(
		j,
		"nursery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetNursery(val INursery) {
	_jsii_.Set(
		j,
		"nursery",
		val,
	)
}

func (j *jsiiProxy_IRules) Performance() IPerformance {
	var returns IPerformance
	_jsii_.Get(
		j,
		"performance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetPerformance(val IPerformance) {
	_jsii_.Set(
		j,
		"performance",
		val,
	)
}

func (j *jsiiProxy_IRules) Recommended() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"recommended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetRecommended(val *bool) {
	_jsii_.Set(
		j,
		"recommended",
		val,
	)
}

func (j *jsiiProxy_IRules) Security() ISecurity {
	var returns ISecurity
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetSecurity(val ISecurity) {
	_jsii_.Set(
		j,
		"security",
		val,
	)
}

func (j *jsiiProxy_IRules) Style() IStyle {
	var returns IStyle
	_jsii_.Get(
		j,
		"style",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetStyle(val IStyle) {
	_jsii_.Set(
		j,
		"style",
		val,
	)
}

func (j *jsiiProxy_IRules) Suspicious() ISuspicious {
	var returns ISuspicious
	_jsii_.Get(
		j,
		"suspicious",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetSuspicious(val ISuspicious) {
	_jsii_.Set(
		j,
		"suspicious",
		val,
	)
}

