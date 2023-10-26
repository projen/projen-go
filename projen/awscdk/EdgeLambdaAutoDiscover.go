package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/awscdk/internal"
	"github.com/projen/projen-go/projen/cdk"
)

// Creates edge lambdas from entry points discovered in the project's source tree.
// Experimental.
type EdgeLambdaAutoDiscover interface {
	cdk.AutoDiscoverBase
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

// The jsii proxy struct for EdgeLambdaAutoDiscover
type jsiiProxy_EdgeLambdaAutoDiscover struct {
	internal.Type__cdkAutoDiscoverBase
}

func (j *jsiiProxy_EdgeLambdaAutoDiscover) Entrypoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entrypoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeLambdaAutoDiscover) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeLambdaAutoDiscover) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewEdgeLambdaAutoDiscover(project projen.Project, options *EdgeLambdaAutoDiscoverOptions) EdgeLambdaAutoDiscover {
	_init_.Initialize()

	if err := validateNewEdgeLambdaAutoDiscoverParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgeLambdaAutoDiscover{}

	_jsii_.Create(
		"projen.awscdk.EdgeLambdaAutoDiscover",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewEdgeLambdaAutoDiscover_Override(e EdgeLambdaAutoDiscover, project projen.Project, options *EdgeLambdaAutoDiscoverOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.awscdk.EdgeLambdaAutoDiscover",
		[]interface{}{project, options},
		e,
	)
}

// Test whether the given construct is a component.
// Experimental.
func EdgeLambdaAutoDiscover_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgeLambdaAutoDiscover_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.EdgeLambdaAutoDiscover",
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
func EdgeLambdaAutoDiscover_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgeLambdaAutoDiscover_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.awscdk.EdgeLambdaAutoDiscover",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgeLambdaAutoDiscover) PostSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"postSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgeLambdaAutoDiscover) PreSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"preSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgeLambdaAutoDiscover) Synthesize() {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgeLambdaAutoDiscover) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

