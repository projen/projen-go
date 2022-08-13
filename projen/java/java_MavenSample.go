package java

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/java/internal"
)

// Java code sample.
// Experimental.
type MavenSample interface {
	projen.Component
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
}

// The jsii proxy struct for MavenSample
type jsiiProxy_MavenSample struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_MavenSample) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewMavenSample(project projen.Project, options *MavenSampleOptions) MavenSample {
	_init_.Initialize()

	j := jsiiProxy_MavenSample{}

	_jsii_.Create(
		"projen.java.MavenSample",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewMavenSample_Override(m MavenSample, project projen.Project, options *MavenSampleOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.java.MavenSample",
		[]interface{}{project, options},
		m,
	)
}

func (m *jsiiProxy_MavenSample) PostSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"postSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenSample) PreSynthesize() {
	_jsii_.InvokeVoid(
		m,
		"preSynthesize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MavenSample) Synthesize() {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		nil, // no parameters
	)
}

