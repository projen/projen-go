package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Manages dependencies using a requirements.txt file and the pip CLI tool.
// Experimental.
type Pip interface {
	projen.Component
	IPythonDeps
	// A task that installs and updates dependencies.
	// Experimental.
	InstallTask() projen.Task
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
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
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Pip
type jsiiProxy_Pip struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonDeps
}

func (j *jsiiProxy_Pip) InstallTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pip) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pip) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewPip(project projen.Project, _options *PipOptions) Pip {
	_init_.Initialize()

	if err := validateNewPipParameters(project, _options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pip{}

	_jsii_.Create(
		"projen.python.Pip",
		[]interface{}{project, _options},
		&j,
	)

	return &j
}

// Experimental.
func NewPip_Override(p Pip, project projen.Project, _options *PipOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Pip",
		[]interface{}{project, _options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Pip_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePip_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Pip",
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
func Pip_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePip_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Pip",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pip) AddDependency(spec *string) {
	if err := p.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Pip) AddDevDependency(spec *string) {
	if err := p.validateAddDevDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Pip) InstallDependencies() {
	_jsii_.InvokeVoid(
		p,
		"installDependencies",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pip) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pip) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pip) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pip) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

