package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Defines renovatebot configuration for projen project.
//
// Ignores the versions controlled by Projen.
// Experimental.
type Renovatebot interface {
	Component
	// The file holding the renovatebot configuration.
	// Experimental.
	File() JsonFile
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

// The jsii proxy struct for Renovatebot
type jsiiProxy_Renovatebot struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_Renovatebot) File() JsonFile {
	var returns JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Renovatebot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Renovatebot) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewRenovatebot(project Project, options *RenovatebotOptions) Renovatebot {
	_init_.Initialize()

	if err := validateNewRenovatebotParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Renovatebot{}

	_jsii_.Create(
		"projen.Renovatebot",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewRenovatebot_Override(r Renovatebot, project Project, options *RenovatebotOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.Renovatebot",
		[]interface{}{project, options},
		r,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Renovatebot_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRenovatebot_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Renovatebot",
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
func Renovatebot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRenovatebot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.Renovatebot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Renovatebot) PostSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"postSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Renovatebot) PreSynthesize() {
	_jsii_.InvokeVoid(
		r,
		"preSynthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Renovatebot) Synthesize() {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Renovatebot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

