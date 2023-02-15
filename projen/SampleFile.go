// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Produces a file with the given contents but only once, if the file doesn't already exist.
//
// Use this for creating example code files or other resources.
// Experimental.
type SampleFile interface {
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

// The jsii proxy struct for SampleFile
type jsiiProxy_SampleFile struct {
	jsiiProxy_Component
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

