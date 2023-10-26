package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen/internal"
)

// Represents a project component.
// Experimental.
type Component interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() Project
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

// The jsii proxy struct for Component
type jsiiProxy_Component struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Component) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Component) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponent(scope constructs.IConstruct, id *string) Component {
	_init_.Initialize()

	if err := validateNewComponentParameters(scope); err != nil {
		panic(err)
	}
	j := jsiiProxy_Component{}

	_jsii_.Create(
		"projen.Component",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewComponent_Override(c Component, scope constructs.IConstruct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Component",
		[]interface{}{scope, id},
		c,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Component_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComponent_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Component",
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
func Component_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComponent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Component",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Component) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Component) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Component) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Component) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

