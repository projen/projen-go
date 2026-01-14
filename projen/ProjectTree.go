package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Generates a `.projen/tree.json` file that provides a snapshot of your project's component hierarchy. This file includes metadata about each component such as file paths, types, and the projen version used.
//
// The tree file is helpful for:
// - Understanding how your project is structured
// - Debugging component relationships
// - Verifying which versions synthesized the project.
// Experimental.
type ProjectTree interface {
	Component
	// Experimental.
	File() JsonFile
	// Experimental.
	SetFile(val JsonFile)
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

// The jsii proxy struct for ProjectTree
type jsiiProxy_ProjectTree struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_ProjectTree) File() JsonFile {
	var returns JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectTree) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectTree) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjectTree(project Project) ProjectTree {
	_init_.Initialize()

	if err := validateNewProjectTreeParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectTree{}

	_jsii_.Create(
		"projen.ProjectTree",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewProjectTree_Override(p ProjectTree, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ProjectTree",
		[]interface{}{project},
		p,
	)
}

func (j *jsiiProxy_ProjectTree)SetFile(val JsonFile) {
	if err := j.validateSetFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"file",
		val,
	)
}

// Test whether the given construct is a component.
// Experimental.
func ProjectTree_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectTree_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjectTree",
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
func ProjectTree_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectTree_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjectTree",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectTree) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectTree) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectTree) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectTree) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

