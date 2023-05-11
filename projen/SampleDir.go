package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Renders the given files into the directory if the directory does not exist.
//
// Use this to create sample code files.
// Experimental.
type SampleDir interface {
	Component
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
}

// The jsii proxy struct for SampleDir
type jsiiProxy_SampleDir struct {
	jsiiProxy_Component
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

