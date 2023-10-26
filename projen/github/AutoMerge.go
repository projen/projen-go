package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Sets up mergify to merging approved pull requests.
//
// If `buildJob` is specified, the specified GitHub workflow job ID is required
// to succeed in order for the PR to be merged.
//
// `approvedReviews` specified the number of code review approvals required for
// the PR to be merged.
// Experimental.
type AutoMerge interface {
	projen.Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Adds conditions to the auto merge rule.
	// Experimental.
	AddConditions(conditions ...*string)
	// Adds conditions that will be rendered only during synthesis.
	// Experimental.
	AddConditionsLater(later IAddConditionsLater)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoMerge
type jsiiProxy_AutoMerge struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_AutoMerge) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoMerge) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoMerge(github GitHub, options *AutoMergeOptions) AutoMerge {
	_init_.Initialize()

	if err := validateNewAutoMergeParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoMerge{}

	_jsii_.Create(
		"projen.github.AutoMerge",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoMerge_Override(a AutoMerge, github GitHub, options *AutoMergeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.AutoMerge",
		[]interface{}{github, options},
		a,
	)
}

// Test whether the given construct is a component.
// Experimental.
func AutoMerge_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoMerge_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.AutoMerge",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AutoMerge_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoMerge_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.AutoMerge",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoMerge) AddConditions(conditions ...*string) {
	args := []interface{}{}
	for _, a := range conditions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addConditions",
		args,
	)
}

func (a *jsiiProxy_AutoMerge) AddConditionsLater(later IAddConditionsLater) {
	if err := a.validateAddConditionsLaterParameters(later); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addConditionsLater",
		[]interface{}{later},
	)
}

func (a *jsiiProxy_AutoMerge) PostSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"postSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoMerge) PreSynthesize() {
	_jsii_.InvokeVoid(
		a,
		"preSynthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoMerge) Synthesize() {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoMerge) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

