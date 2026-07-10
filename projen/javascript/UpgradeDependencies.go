package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript/internal"
)

// Upgrade node project dependencies.
// Experimental.
type UpgradeDependencies interface {
	projen.Component
	// Container definitions for the upgrade workflow.
	// Experimental.
	ContainerOptions() *workflows.ContainerOptions
	// Experimental.
	SetContainerOptions(val *workflows.ContainerOptions)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// A task run after the upgrade task.
	// Experimental.
	PostUpgradeTask() projen.Task
	// Experimental.
	// Returns `NodeProject`, use interface conversion if needed.
	Project() projen.Project
	// The upgrade task.
	// Experimental.
	UpgradeTask() projen.Task
	// The workflows that execute the upgrades.
	//
	// One workflow per branch.
	// Experimental.
	Workflows() *[]github.GithubWorkflow
	// Add steps to execute a successful build.
	// Experimental.
	AddPostBuildSteps(steps ...*workflows.JobStep)
	// Called once, right after `postSynthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// It is also skipped when post-synthesis steps are disabled, e.g. `--no-post` or `PROJEN_DISABLE_POST`.
	// Use it for one-off setup that can be turned off by the user, like running a task to give the user immediate
	// feedback on their new project. Order across components is not guaranteed.
	// Experimental.
	PostProjectCreation(initProject *projen.InitProject)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Called once, right after `synthesize()`, only when the project is created for the first time.
	//
	// It does not run on later `projen` invocations. It only fires for `projen new` (or `Projects.createProject`).
	// Use it for deterministic, one-off file generation. Order across components is not guaranteed.
	// Experimental.
	ProjectCreation(initProject *projen.InitProject)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for UpgradeDependencies
type jsiiProxy_UpgradeDependencies struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_UpgradeDependencies) ContainerOptions() *workflows.ContainerOptions {
	var returns *workflows.ContainerOptions
	_jsii_.Get(
		j,
		"containerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) PostUpgradeTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postUpgradeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) UpgradeTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"upgradeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) Workflows() *[]github.GithubWorkflow {
	var returns *[]github.GithubWorkflow
	_jsii_.Get(
		j,
		"workflows",
		&returns,
	)
	return returns
}


// Experimental.
func NewUpgradeDependencies(project NodeProject, options *UpgradeDependenciesOptions) UpgradeDependencies {
	_init_.Initialize()

	if err := validateNewUpgradeDependenciesParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_UpgradeDependencies{}

	_jsii_.Create(
		"projen.javascript.UpgradeDependencies",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewUpgradeDependencies_Override(u UpgradeDependencies, project NodeProject, options *UpgradeDependenciesOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.UpgradeDependencies",
		[]interface{}{project, options},
		u,
	)
}

func (j *jsiiProxy_UpgradeDependencies)SetContainerOptions(val *workflows.ContainerOptions) {
	if err := j.validateSetContainerOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerOptions",
		val,
	)
}

// Test whether the given construct is a component.
// Experimental.
func UpgradeDependencies_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpgradeDependencies_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.UpgradeDependencies",
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
func UpgradeDependencies_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpgradeDependencies_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.javascript.UpgradeDependencies",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpgradeDependencies) AddPostBuildSteps(steps ...*workflows.JobStep) {
	if err := u.validateAddPostBuildStepsParameters(&steps); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addPostBuildSteps",
		args,
	)
}

func (u *jsiiProxy_UpgradeDependencies) PostProjectCreation(initProject *projen.InitProject) {
	if err := u.validatePostProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"postProjectCreation",
		[]interface{}{initProject},
	)
}

func (u *jsiiProxy_UpgradeDependencies) PostSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"postSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpgradeDependencies) PreSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"preSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpgradeDependencies) ProjectCreation(initProject *projen.InitProject) {
	if err := u.validateProjectCreationParameters(initProject); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"projectCreation",
		[]interface{}{initProject},
	)
}

func (u *jsiiProxy_UpgradeDependencies) Synthesize() {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpgradeDependencies) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpgradeDependencies) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		u,
		"with",
		args,
		&returns,
	)

	return returns
}

