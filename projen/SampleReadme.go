// CDK for software projects
package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"
)

// Represents a README.md sample file. You are expected to manage this file after creation.
// Experimental.
type SampleReadme interface {
	SampleFile
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

// The jsii proxy struct for SampleReadme
type jsiiProxy_SampleReadme struct {
	jsiiProxy_SampleFile
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

