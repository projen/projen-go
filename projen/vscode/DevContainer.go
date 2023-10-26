package vscode

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/vscode/internal"
)

// A development environment running VSCode in a container;
//
// used by GitHub
// codespaces.
// Experimental.
type DevContainer interface {
	projen.Component
	IDevContainerEnvironment
	// Direct access to the devcontainer configuration (escape hatch).
	// Experimental.
	Config() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Add a custom Docker image or Dockerfile for the container.
	// Experimental.
	AddDockerImage(image projen.DevEnvironmentDockerImage)
	// Adds a list of VSCode features that should be automatically installed in the container.
	// Experimental.
	AddFeatures(features ...*DevContainerFeature)
	// Adds ports that should be exposed (forwarded) from the container.
	// Experimental.
	AddPorts(ports ...*string)
	// Adds tasks to run when the container starts.
	//
	// Tasks will be run in sequence.
	// Experimental.
	AddTasks(tasks ...projen.Task)
	// Adds a list of VSCode extensions that should be automatically installed in the container.
	// Experimental.
	AddVscodeExtensions(extensions ...*string)
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

// The jsii proxy struct for DevContainer
type jsiiProxy_DevContainer struct {
	internal.Type__projenComponent
	jsiiProxy_IDevContainerEnvironment
}

func (j *jsiiProxy_DevContainer) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevContainer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevContainer) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewDevContainer(project projen.Project, options *DevContainerOptions) DevContainer {
	_init_.Initialize()

	if err := validateNewDevContainerParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevContainer{}

	_jsii_.Create(
		"projen.vscode.DevContainer",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewDevContainer_Override(d DevContainer, project projen.Project, options *DevContainerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.vscode.DevContainer",
		[]interface{}{project, options},
		d,
	)
}

// Test whether the given construct is a component.
// Experimental.
func DevContainer_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevContainer_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.DevContainer",
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
func DevContainer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevContainer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.vscode.DevContainer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevContainer) AddDockerImage(image projen.DevEnvironmentDockerImage) {
	if err := d.validateAddDockerImageParameters(image); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addDockerImage",
		[]interface{}{image},
	)
}

func (d *jsiiProxy_DevContainer) AddFeatures(features ...*DevContainerFeature) {
	if err := d.validateAddFeaturesParameters(&features); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range features {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addFeatures",
		args,
	)
}

func (d *jsiiProxy_DevContainer) AddPorts(ports ...*string) {
	args := []interface{}{}
	for _, a := range ports {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addPorts",
		args,
	)
}

func (d *jsiiProxy_DevContainer) AddTasks(tasks ...projen.Task) {
	args := []interface{}{}
	for _, a := range tasks {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addTasks",
		args,
	)
}

func (d *jsiiProxy_DevContainer) AddVscodeExtensions(extensions ...*string) {
	args := []interface{}{}
	for _, a := range extensions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addVscodeExtensions",
		args,
	)
}

func (d *jsiiProxy_DevContainer) PostSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"postSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevContainer) PreSynthesize() {
	_jsii_.InvokeVoid(
		d,
		"preSynthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevContainer) Synthesize() {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevContainer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

