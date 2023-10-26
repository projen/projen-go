package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a README.md sample file. You are expected to manage this file after creation.
// Experimental.
type SampleReadme interface {
	SampleFile
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

// The jsii proxy struct for SampleReadme
type jsiiProxy_SampleReadme struct {
	jsiiProxy_SampleFile
}

func (j *jsiiProxy_SampleReadme) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SampleReadme) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewSampleReadme(project Project, props *SampleReadmeProps) SampleReadme {
	_init_.Initialize()

	if err := validateNewSampleReadmeParameters(project, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SampleReadme{}

	_jsii_.Create(
		"projen.SampleReadme",
		[]interface{}{project, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSampleReadme_Override(s SampleReadme, project Project, props *SampleReadmeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SampleReadme",
		[]interface{}{project, props},
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func SampleReadme_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSampleReadme_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SampleReadme",
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
func SampleReadme_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSampleReadme_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SampleReadme",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SampleReadme) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleReadme) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleReadme) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleReadme) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

