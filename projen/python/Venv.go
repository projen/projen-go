package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Manages a virtual environment through the Python venv module.
// Experimental.
type Venv interface {
	projen.Component
	IPythonEnv
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
	// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
	// Experimental.
	SetupEnvironment()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Venv
type jsiiProxy_Venv struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonEnv
}

func (j *jsiiProxy_Venv) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Venv) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewVenv(project projen.Project, options *VenvOptions) Venv {
	_init_.Initialize()

	if err := validateNewVenvParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Venv{}

	_jsii_.Create(
		"projen.python.Venv",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewVenv_Override(v Venv, project projen.Project, options *VenvOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Venv",
		[]interface{}{project, options},
		v,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Venv_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVenv_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Venv",
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
func Venv_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVenv_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Venv",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Venv) PostSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"postSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) PreSynthesize() {
	_jsii_.InvokeVoid(
		v,
		"preSynthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) SetupEnvironment() {
	_jsii_.InvokeVoid(
		v,
		"setupEnvironment",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) Synthesize() {
	_jsii_.InvokeVoid(
		v,
		"synthesize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_Venv) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

