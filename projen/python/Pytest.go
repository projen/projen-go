package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Experimental.
type Pytest interface {
	projen.Component
	// Experimental.
	Project() projen.Project
	// Experimental.
	Testdir() *string
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

// The jsii proxy struct for Pytest
type jsiiProxy_Pytest struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Pytest) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pytest) Testdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"testdir",
		&returns,
	)
	return returns
}


// Experimental.
func NewPytest(project projen.Project, options *PytestOptions) Pytest {
	_init_.Initialize()

	if err := validateNewPytestParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pytest{}

	_jsii_.Create(
		"projen.python.Pytest",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPytest_Override(p Pytest, project projen.Project, options *PytestOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Pytest",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_Pytest) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pytest) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pytest) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

