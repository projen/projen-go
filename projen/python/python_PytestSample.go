package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Python test code sample.
// Experimental.
type PytestSample interface {
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

// The jsii proxy struct for PytestSample
type jsiiProxy_PytestSample struct {
	internal.Type__projenComponent
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

