package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Represents configuration of a pyproject.toml file for a Poetry project.
// See: https://python-poetry.org/docs/pyproject/
//
// Experimental.
type PoetryPyproject interface {
	projen.Component
	// Experimental.
	File() projen.TomlFile
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

// The jsii proxy struct for PoetryPyproject
type jsiiProxy_PoetryPyproject struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_PoetryPyproject) File() projen.TomlFile {
	var returns projen.TomlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PoetryPyproject) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPoetryPyproject(project projen.Project, options *PoetryPyprojectOptions) PoetryPyproject {
	_init_.Initialize()

	if err := validateNewPoetryPyprojectParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PoetryPyproject{}

	_jsii_.Create(
		"projen.python.PoetryPyproject",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPoetryPyproject_Override(p PoetryPyproject, project projen.Project, options *PoetryPyprojectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.PoetryPyproject",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_PoetryPyproject) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PoetryPyproject) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PoetryPyproject) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

