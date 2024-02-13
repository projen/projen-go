package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// A task that installs dependencies (honouring the lockfile).
	// Experimental.
	InstallCiTask() projen.Task
	// A task that updates the lockfile and installs dependencies.
	// Experimental.
	InstallTask() projen.Task
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// A task that uploads the package to a package repository.
	// Experimental.
	PublishTask() projen.Task
	// A task that uploads the package to the Test PyPI repository.
	// Experimental.
	PublishTestTask() projen.Task
	// A task that for upgrades dependencies.
	// Experimental.
	UpgradeTask() projen.Task
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Poetry
type jsiiProxy_Poetry struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonDeps
	jsiiProxy_IPythonEnv
	jsiiProxy_IPythonPackaging
}

func (j *jsiiProxy_Poetry) InstallCiTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installCiTask",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_Poetry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

func (j *jsiiProxy_Poetry) UpgradeTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"upgradeTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewPoetry(project projen.Project, options *PoetryOptions) Poetry {
	_init_.Initialize()

	if err := validateNewPoetryParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Poetry{}

	_jsii_.Create(
		"projen.python.Poetry",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPoetry_Override(p Poetry, project projen.Project, options *PoetryOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Poetry",
		[]interface{}{project, options},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Poetry_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePoetry_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Poetry",
		"isComponent",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func Poetry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePoetry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Poetry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Poetry) AddDependency(spec *string) {
	if err := p.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{spec},
	)
}

func (p *jsiiProxy_Poetry) AddDevDependency(spec *string) {
	if err := p.validateAddDevDependencyParameters(spec); err != nil {
		panic(err)
	}
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

func (p *jsiiProxy_Poetry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

