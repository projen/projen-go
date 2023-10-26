package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
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

func (j *jsiiProxy_PoetryPyproject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

// Test whether the given construct is a component.
// Experimental.
func PoetryPyproject_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePoetryPyproject_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.PoetryPyproject",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PoetryPyproject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePoetryPyproject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.PoetryPyproject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

func (p *jsiiProxy_PoetryPyproject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

