package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Manages packaging through setuptools with a setup.py script.
// Experimental.
type Setuptools interface {
	projen.Component
	IPythonPackaging
	// Experimental.
	Project() projen.Project
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
	// A task that uploads the package to the Test PyPI repository.
	// Experimental.
	PublishTestTask() projen.Task
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

// The jsii proxy struct for Setuptools
type jsiiProxy_Setuptools struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonPackaging
}

func (j *jsiiProxy_Setuptools) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Setuptools) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Setuptools) PublishTestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTestTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewSetuptools(project projen.Project, options *PythonPackagingOptions) Setuptools {
	_init_.Initialize()

	if err := validateNewSetuptoolsParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Setuptools{}

	_jsii_.Create(
		"projen.python.Setuptools",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewSetuptools_Override(s Setuptools, project projen.Project, options *PythonPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Setuptools",
		[]interface{}{project, options},
		s,
	)
}

func (s *jsiiProxy_Setuptools) PostSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"postSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Setuptools) PreSynthesize() {
	_jsii_.InvokeVoid(
		s,
		"preSynthesize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Setuptools) Synthesize() {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		nil, // no parameters
	)
}

