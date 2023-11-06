package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Experimental.
type Yarnrc interface {
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

// The jsii proxy struct for Yarnrc
type jsiiProxy_Yarnrc struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Yarnrc) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Yarnrc) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewYarnrc(project projen.Project, version *string, options *YarnrcOptions) Yarnrc {
	_init_.Initialize()

	if err := validateNewYarnrcParameters(project, version, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Yarnrc{}

	_jsii_.Create(
		"projen.javascript.Yarnrc",
		[]interface{}{project, version, options},
		&j,
	)

	return &j
}

// Experimental.
func NewYarnrc_Override(y Yarnrc, project projen.Project, version *string, options *YarnrcOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Yarnrc",
		[]interface{}{project, version, options},
		y,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Yarnrc_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateYarnrc_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Yarnrc",
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
func Yarnrc_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateYarnrc_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.Yarnrc",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (y *jsiiProxy_Yarnrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		y,
		"postSynthesize",
		nil, // no parameters
	)
}

func (y *jsiiProxy_Yarnrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		y,
		"preSynthesize",
		nil, // no parameters
	)
}

func (y *jsiiProxy_Yarnrc) Synthesize() {
	_jsii_.InvokeVoid(
		y,
		"synthesize",
		nil, // no parameters
	)
}

func (y *jsiiProxy_Yarnrc) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		y,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

