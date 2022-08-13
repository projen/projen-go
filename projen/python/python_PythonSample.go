package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Python code sample.
// Experimental.
type PythonSample interface {
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

// The jsii proxy struct for PythonSample
type jsiiProxy_PythonSample struct {
	internal.Type__projenComponent
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

