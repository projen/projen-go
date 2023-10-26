package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Python test code sample.
// Experimental.
type PytestSample interface {
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

// The jsii proxy struct for PytestSample
type jsiiProxy_PytestSample struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PytestSample) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PytestSample) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPytestSample(project projen.Project, options *PytestSampleOptions) PytestSample {
	_init_.Initialize()

	if err := validateNewPytestSampleParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PytestSample{}

	_jsii_.Create(
		"projen.python.PytestSample",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPytestSample_Override(p PytestSample, project projen.Project, options *PytestSampleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.PytestSample",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func PytestSample_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePytestSample_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.PytestSample",
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
func PytestSample_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePytestSample_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.PytestSample",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PytestSample) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PytestSample) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PytestSample) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PytestSample) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

