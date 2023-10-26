package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// Implements JUnit-based testing.
// Experimental.
type Junit interface {
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

// The jsii proxy struct for Junit
type jsiiProxy_Junit struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Junit) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Junit) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewJunit(project projen.Project, options *JunitOptions) Junit {
	_init_.Initialize()

	if err := validateNewJunitParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Junit{}

	_jsii_.Create(
		"projen.java.Junit",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJunit_Override(j Junit, project projen.Project, options *JunitOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.Junit",
		[]interface{}{project, options},
		j,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Junit_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJunit_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.java.Junit",
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
func Junit_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJunit_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.java.Junit",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_Junit) PostSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"postSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Junit) PreSynthesize() {
	_jsii_.InvokeVoid(
		j,
		"preSynthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Junit) Synthesize() {
	_jsii_.InvokeVoid(
		j,
		"synthesize",
		nil, // no parameters
	)
}

func (j *jsiiProxy_Junit) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

