package cdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

// Base class for locating integration tests in the project's test tree.
// Experimental.
type IntegrationTestAutoDiscoverBase interface {
	AutoDiscoverBase
	// Auto-discovered entry points with paths relative to the project directory.
	// Experimental.
	Entrypoints() *[]*string
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

// The jsii proxy struct for IntegrationTestAutoDiscoverBase
type jsiiProxy_IntegrationTestAutoDiscoverBase struct {
	jsiiProxy_AutoDiscoverBase
}

func (j *jsiiProxy_IntegrationTestAutoDiscoverBase) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestAutoDiscoverBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationTestAutoDiscoverBase) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegrationTestAutoDiscoverBase(project projen.Project, options *IntegrationTestAutoDiscoverBaseOptions) IntegrationTestAutoDiscoverBase {
	_init_.Initialize()

	if err := validateNewIntegrationTestAutoDiscoverBaseParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationTestAutoDiscoverBase{}

	_jsii_.Create(
		"projen.cdk.IntegrationTestAutoDiscoverBase",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegrationTestAutoDiscoverBase_Override(i IntegrationTestAutoDiscoverBase, project projen.Project, options *IntegrationTestAutoDiscoverBaseOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk.IntegrationTestAutoDiscoverBase",
		[]interface{}{project, options},
		i,
	)
}

// Test whether the given construct is a component.
// Experimental.
func IntegrationTestAutoDiscoverBase_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTestAutoDiscoverBase_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk.IntegrationTestAutoDiscoverBase",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func IntegrationTestAutoDiscoverBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationTestAutoDiscoverBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk.IntegrationTestAutoDiscoverBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationTestAutoDiscoverBase) PostSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"postSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscoverBase) PreSynthesize() {
	_jsii_.InvokeVoid(
		i,
		"preSynthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscoverBase) Synthesize() {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationTestAutoDiscoverBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

