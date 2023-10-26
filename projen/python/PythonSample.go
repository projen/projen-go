package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Python code sample.
// Experimental.
type PythonSample interface {
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

// The jsii proxy struct for PythonSample
type jsiiProxy_PythonSample struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PythonSample) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonSample) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonSample(project projen.Project, options *PythonSampleOptions) PythonSample {
	_init_.Initialize()

	if err := validateNewPythonSampleParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PythonSample{}

	_jsii_.Create(
		"projen.python.PythonSample",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonSample_Override(p PythonSample, project projen.Project, options *PythonSampleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.PythonSample",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func PythonSample_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePythonSample_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.PythonSample",
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
func PythonSample_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePythonSample_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.PythonSample",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonSample) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonSample) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonSample) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonSample) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

