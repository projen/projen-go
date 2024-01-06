package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
	"github.com/projen/projen-go/projen/github/workflows"
)

// The primary or initial job of a TaskWorkflow.
// Experimental.
type TaskWorkflowJob interface {
	projen.Component
	// Experimental.
	Concurrency() interface{}
	// Experimental.
	Container() *workflows.ContainerOptions
	// Experimental.
	ContinueOnError() *bool
	// Experimental.
	Defaults() *workflows.JobDefaults
	// Experimental.
	Env() *map[string]*string
	// Experimental.
	Environment() interface{}
	// Experimental.
	If() *string
	// Experimental.
	Name() *string
	// Experimental.
	Needs() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Outputs() *map[string]*workflows.JobStepOutput
	// Experimental.
	Permissions() *workflows.JobPermissions
	// Experimental.
	Project() projen.Project
	// Experimental.
	RunsOn() *[]*string
	// Experimental.
	RunsOnGroup() *projen.GroupRunnerOptions
	// Experimental.
	Services() *map[string]*workflows.ContainerOptions
	// Experimental.
	Steps() *[]*workflows.JobStep
	// Experimental.
	Strategy() *workflows.JobStrategy
	// Experimental.
	TimeoutMinutes() *float64
	// Experimental.
	Tools() *workflows.Tools
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

// The jsii proxy struct for TaskWorkflowJob
type jsiiProxy_TaskWorkflowJob struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_TaskWorkflowJob) Concurrency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Container() *workflows.ContainerOptions {
	var returns *workflows.ContainerOptions
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) ContinueOnError() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"continueOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Defaults() *workflows.JobDefaults {
	var returns *workflows.JobDefaults
	_jsii_.Get(
		j,
		"defaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) If() *string {
	var returns *string
	_jsii_.Get(
		j,
		"if",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Needs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"needs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Outputs() *map[string]*workflows.JobStepOutput {
	var returns *map[string]*workflows.JobStepOutput
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Permissions() *workflows.JobPermissions {
	var returns *workflows.JobPermissions
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) RunsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"runsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) RunsOnGroup() *projen.GroupRunnerOptions {
	var returns *projen.GroupRunnerOptions
	_jsii_.Get(
		j,
		"runsOnGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Services() *map[string]*workflows.ContainerOptions {
	var returns *map[string]*workflows.ContainerOptions
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Steps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Strategy() *workflows.JobStrategy {
	var returns *workflows.JobStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflowJob) Tools() *workflows.Tools {
	var returns *workflows.Tools
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}


// Experimental.
func NewTaskWorkflowJob(scope constructs.IConstruct, task projen.Task, options *TaskWorkflowJobOptions) TaskWorkflowJob {
	_init_.Initialize()

	if err := validateNewTaskWorkflowJobParameters(scope, task, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaskWorkflowJob{}

	_jsii_.Create(
		"projen.github.TaskWorkflowJob",
		[]interface{}{scope, task, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTaskWorkflowJob_Override(t TaskWorkflowJob, scope constructs.IConstruct, task projen.Task, options *TaskWorkflowJobOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.TaskWorkflowJob",
		[]interface{}{scope, task, options},
		t,
	)
}

// Test whether the given construct is a component.
// Experimental.
func TaskWorkflowJob_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaskWorkflowJob_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.TaskWorkflowJob",
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
func TaskWorkflowJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaskWorkflowJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.TaskWorkflowJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskWorkflowJob) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskWorkflowJob) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskWorkflowJob) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskWorkflowJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

