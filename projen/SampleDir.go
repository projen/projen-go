package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Renders the given files into the directory if the directory does not exist.
//
// Use this to create sample code files.
// Experimental.
type SampleDir interface {
	Component
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

// The jsii proxy struct for SampleDir
type jsiiProxy_SampleDir struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_SampleDir) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SampleDir) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Create sample files in the given directory if the given directory does not exist.
// Experimental.
func NewSampleDir(project Project, dir *string, options *SampleDirOptions) SampleDir {
	_init_.Initialize()

	if err := validateNewSampleDirParameters(project, dir, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SampleDir{}

	_jsii_.Create(
		"projen.SampleDir",
		[]interface{}{project, dir, options},
		&j,
	)

	return &j
}

// Create sample files in the given directory if the given directory does not exist.
// Experimental.
func NewSampleDir_Override(s SampleDir, project Project, dir *string, options *SampleDirOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SampleDir",
		[]interface{}{project, dir, options},
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func SampleDir_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSampleDir_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SampleDir",
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
func SampleDir_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSampleDir_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SampleDir",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SampleDir) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleDir) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleDir) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleDir) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

