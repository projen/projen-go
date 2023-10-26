package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Produces a file with the given contents but only once, if the file doesn't already exist.
//
// Use this for creating example code files or other resources.
// Experimental.
type SampleFile interface {
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

// The jsii proxy struct for SampleFile
type jsiiProxy_SampleFile struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_SampleFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SampleFile) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Creates a new SampleFile object.
// Experimental.
func NewSampleFile(project Project, filePath *string, options *SampleFileOptions) SampleFile {
	_init_.Initialize()

	if err := validateNewSampleFileParameters(project, filePath, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SampleFile{}

	_jsii_.Create(
		"projen.SampleFile",
		[]interface{}{project, filePath, options},
		&j,
	)

	return &j
}

// Creates a new SampleFile object.
// Experimental.
func NewSampleFile_Override(s SampleFile, project Project, filePath *string, options *SampleFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.SampleFile",
		[]interface{}{project, filePath, options},
		s,
	)
}

// Test whether the given construct is a component.
// Experimental.
func SampleFile_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSampleFile_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SampleFile",
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
func SampleFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSampleFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.SampleFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SampleFile) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleFile) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleFile) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SampleFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

