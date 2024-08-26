package cdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/cdk/internal"
)

// Creates a markdown file based on the jsii manifest: - Adds a `docgen` script to package.json - Runs `jsii-docgen` after compilation - Enforces that markdown file is checked in.
// Experimental.
type JsiiDocgen interface {
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

// The jsii proxy struct for JsiiDocgen
type jsiiProxy_JsiiDocgen struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_JsiiDocgen) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsiiDocgen) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewJsiiDocgen(scope constructs.IConstruct, options *JsiiDocgenOptions) JsiiDocgen {
	_init_.Initialize()

	if err := validateNewJsiiDocgenParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_JsiiDocgen{}

	_jsii_.Create(
		"projen.cdk.JsiiDocgen",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJsiiDocgen_Override(j JsiiDocgen, scope constructs.IConstruct, options *JsiiDocgenOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.cdk.JsiiDocgen",
		[]interface{}{scope, options},
		j,
	)
}

// Test whether the given construct is a component.
// Experimental.
func JsiiDocgen_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJsiiDocgen_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk.JsiiDocgen",
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
func JsiiDocgen_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJsiiDocgen_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.cdk.JsiiDocgen",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JsiiDocgen) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JsiiDocgen) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JsiiDocgen) Synthesize() {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JsiiDocgen) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

