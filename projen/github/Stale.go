package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
)

// Warns and then closes issues and PRs that have had no activity for a specified amount of time.
//
// The default configuration will:
//
//  * Add a "Stale" label to pull requests after 14 days and closed after 2 days
//  * Add a "Stale" label to issues after 60 days and closed after 7 days
//  * If a comment is added, the label will be removed and timer is restarted.
// See: https://github.com/actions/stale
//
// Experimental.
type Stale interface {
	projen.Component
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
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

// The jsii proxy struct for Stale
type jsiiProxy_Stale struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Stale) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stale) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewStale(github GitHub, options *StaleOptions) Stale {
	_init_.Initialize()

	if err := validateNewStaleParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Stale{}

	_jsii_.Create(
		"projen.github.Stale",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewStale_Override(s Stale, github GitHub, options *StaleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.Stale",
		[]interface{}{github, options},
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Stale_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStale_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.Stale",
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
func Stale_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStale_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.Stale",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stale) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stale) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stale) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stale) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

