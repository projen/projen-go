package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Manage project dependencies, virtual environments, and packaging through the poetry CLI tool.
// Experimental.
type Poetry interface {
	projen.Component
	IPythonDeps
	IPythonEnv
	IPythonPackaging
	// A task that installs and updates dependencies.
	// Experimental.
	InstallTask() projen.Task
	// Experimental.
	Project() projen.Project
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
	// A task that uploads the package to the Test PyPI repository.
	// Experimental.
	PublishTestTask() projen.Task
	// Adds a runtime dependency.
	// Experimental.
	AddDependency(spec *string)
	// Adds a dev dependency.
	// Experimental.
	AddDevDependency(spec *string)
	// Installs dependencies (called during post-synthesis).
	// Experimental.
	InstallDependencies()
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Initializes the virtual environment if it doesn't exist (called during post-synthesis).
	// Experimental.
	SetupEnvironment()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Poetry
type jsiiProxy_Poetry struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonDeps
	jsiiProxy_IPythonEnv
	jsiiProxy_IPythonPackaging
}

func (j *jsiiProxy_Poetry) InstallTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poetry) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poetry) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Poetry) PublishTestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTestTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewPoetry(project projen.Project, options *PythonPackagingOptions) Poetry {
	_init_.Initialize()

	j := jsiiProxy_Poetry{}

	_jsii_.Create(
		"projen.python.Poetry",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPoetry_Override(p Poetry, project projen.Project, options *PythonPackagingOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Poetry",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_Poetry) AddDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Poetry) AddDevDependency(spec *string) {
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Poetry) InstallDependencies() {
	_jsii_.InvokeVoid(
		p,
		"installDependencies",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) SetupEnvironment() {
	_jsii_.InvokeVoid(
		p,
		"setupEnvironment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Poetry) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

