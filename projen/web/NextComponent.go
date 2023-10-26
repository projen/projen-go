package web

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript"
	"github.com/projen/projen-go/projen/web/internal"
)

// Experimental.
type NextComponent interface {
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

// The jsii proxy struct for NextComponent
type jsiiProxy_NextComponent struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NextComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NextComponent) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNextComponent(project javascript.NodeProject, options *NextComponentOptions) NextComponent {
	_init_.Initialize()

	if err := validateNewNextComponentParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NextComponent{}

	_jsii_.Create(
		"projen.web.NextComponent",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNextComponent_Override(n NextComponent, project javascript.NodeProject, options *NextComponentOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.web.NextComponent",
		[]interface{}{project, options},
		n,
	)
}

// Test whether the given construct is a component.
// Experimental.
func NextComponent_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextComponent_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.web.NextComponent",
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
func NextComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNextComponent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.web.NextComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NextComponent) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NextComponent) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NextComponent) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NextComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

