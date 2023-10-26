package projen

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Manages a standard build process for all projects.
//
// Build spawns these tasks in order:
// 1. default
// 2. pre-compile
// 3. compile
// 4. post-compile
// 5. test
// 6. package
// Experimental.
type ProjectBuild interface {
	Component
	// The task responsible for a full release build.
	// Experimental.
	BuildTask() Task
	// Compiles the code.
	//
	// By default for node.js projects this task is empty.
	// Experimental.
	CompileTask() Task
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The "package" task.
	// Experimental.
	PackageTask() Task
	// Post-compile task.
	// Experimental.
	PostCompileTask() Task
	// Pre-compile task.
	// Experimental.
	PreCompileTask() Task
	// Experimental.
	Project() Project
	// Tests the code.
	// Experimental.
	TestTask() Task
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

// The jsii proxy struct for ProjectBuild
type jsiiProxy_ProjectBuild struct {
	jsiiProxy_Component
}

func (j *jsiiProxy_ProjectBuild) BuildTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) CompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) PackageTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) PostCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) PreCompileTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) Project() Project {
	var returns Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectBuild) TestTask() Task {
	var returns Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjectBuild(project Project) ProjectBuild {
	_init_.Initialize()

	if err := validateNewProjectBuildParameters(project); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectBuild{}

	_jsii_.Create(
		"projen.ProjectBuild",
		[]interface{}{project},
		&j,
	)

	return &j
}

// Experimental.
func NewProjectBuild_Override(p ProjectBuild, project Project) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.ProjectBuild",
		[]interface{}{project},
		p,
	)
}

// Test whether the given construct is a component.
// Experimental.
func ProjectBuild_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectBuild_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjectBuild",
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
func ProjectBuild_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectBuild_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.ProjectBuild",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectBuild) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectBuild) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectBuild) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectBuild) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

