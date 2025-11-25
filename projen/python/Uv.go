package python

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/python/internal"
)

// Manage project dependencies, virtual environments, and packaging through uv.
// Experimental.
type Uv interface {
	projen.Component
	IPythonDeps
	IPythonEnv
	IPythonPackaging
	// The `pyproject.toml` file.
	// Experimental.
	File() PyprojectTomlFile
	// A task that installs and updates dependencies.
	// Experimental.
	InstallCiTask() projen.Task
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Uv
type jsiiProxy_Uv struct {
	internal.Type__projenComponent
	jsiiProxy_IPythonDeps
	jsiiProxy_IPythonEnv
	jsiiProxy_IPythonPackaging
}

func (j *jsiiProxy_Uv) File() PyprojectTomlFile {
	var returns PyprojectTomlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Uv) InstallCiTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installCiTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Uv) InstallTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"installTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Uv) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Uv) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Uv) PublishTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Uv) PublishTestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"publishTestTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewUv(scope constructs.IConstruct, options *UvOptions) Uv {
	_init_.Initialize()

	if err := validateNewUvParameters(scope, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Uv{}

	_jsii_.Create(
		"projen.python.Uv",
		[]interface{}{scope, options},
		&j,
	)

	return &j
}

// Experimental.
func NewUv_Override(u Uv, scope constructs.IConstruct, options *UvOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.python.Uv",
		[]interface{}{scope, options},
		u,
	)
}

// Test whether the given construct is a component.
// Experimental.
func Uv_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUv_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Uv",
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
func Uv_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUv_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.python.Uv",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_Uv) AddDependency(spec *string) {
	if err := u.validateAddDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addDependency",
		[]interface{}{spec},
	)
}

func (u *jsiiProxy_Uv) AddDevDependency(spec *string) {
	if err := u.validateAddDevDependencyParameters(spec); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addDevDependency",
		[]interface{}{spec},
	)
}

func (u *jsiiProxy_Uv) InstallDependencies() {
	_jsii_.InvokeVoid(
		u,
		"installDependencies",
		nil, // no parameters
	)
}

func (u *jsiiProxy_Uv) PostSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"postSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_Uv) PreSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"preSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_Uv) SetupEnvironment() {
	_jsii_.InvokeVoid(
		u,
		"setupEnvironment",
		nil, // no parameters
	)
}

func (u *jsiiProxy_Uv) Synthesize() {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_Uv) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

