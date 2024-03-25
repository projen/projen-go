package build

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build/internal"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Experimental.
type BuildWorkflow interface {
	projen.Component
	// Returns a list of job IDs that are part of the build.
	// Experimental.
	BuildJobIds() *[]*string
	// Name of generated github workflow.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Project() projen.Project
	// Adds another job to the build workflow which is executed after the build job succeeded.
	//
	// Jobs are executed _only_ if the build did NOT self mutate. If the build
	// self-mutate, the branch will either be updated or the build will fail (in
	// forks), so there is no point in executing the post-build job.
	// Experimental.
	AddPostBuildJob(id *string, job *workflows.Job)
	// Run a sequence of commands as a job within the build workflow which is executed after the build job succeeded.
	//
	// Jobs are executed _only_ if the build did NOT self mutate. If the build
	// self-mutate, the branch will either be updated or the build will fail (in
	// forks), so there is no point in executing the post-build job.
	// Experimental.
	AddPostBuildJobCommands(id *string, commands *[]*string, options *AddPostBuildJobCommandsOptions)
	// Run a task as a job within the build workflow which is executed after the build job succeeded.
	//
	// The job will have access to build artifacts and will install project
	// dependencies in order to be able to run any commands used in the tasks.
	//
	// Jobs are executed _only_ if the build did NOT self mutate. If the build
	// self-mutate, the branch will either be updated or the build will fail (in
	// forks), so there is no point in executing the post-build job.
	// Experimental.
	AddPostBuildJobTask(task projen.Task, options *AddPostBuildJobTaskOptions)
	// Adds steps that are executed after the build.
	// Experimental.
	AddPostBuildSteps(steps ...*workflows.JobStep)
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

// The jsii proxy struct for BuildWorkflow
type jsiiProxy_BuildWorkflow struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_BuildWorkflow) BuildJobIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"buildJobIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BuildWorkflow) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewBuildWorkflow(project projen.Project, options *BuildWorkflowOptions) BuildWorkflow {
	_init_.Initialize()

	if err := validateNewBuildWorkflowParameters(project, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_BuildWorkflow{}

	_jsii_.Create(
		"projen.build.BuildWorkflow",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewBuildWorkflow_Override(b BuildWorkflow, project projen.Project, options *BuildWorkflowOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.build.BuildWorkflow",
		[]interface{}{project, options},
		b,
	)
}

// Test whether the given construct is a component.
// Experimental.
func BuildWorkflow_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBuildWorkflow_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.build.BuildWorkflow",
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
func BuildWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBuildWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.build.BuildWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BuildWorkflow) AddPostBuildJob(id *string, job *workflows.Job) {
	if err := b.validateAddPostBuildJobParameters(id, job); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addPostBuildJob",
		[]interface{}{id, job},
	)
}

func (b *jsiiProxy_BuildWorkflow) AddPostBuildJobCommands(id *string, commands *[]*string, options *AddPostBuildJobCommandsOptions) {
	if err := b.validateAddPostBuildJobCommandsParameters(id, commands, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addPostBuildJobCommands",
		[]interface{}{id, commands, options},
	)
}

func (b *jsiiProxy_BuildWorkflow) AddPostBuildJobTask(task projen.Task, options *AddPostBuildJobTaskOptions) {
	if err := b.validateAddPostBuildJobTaskParameters(task, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addPostBuildJobTask",
		[]interface{}{task, options},
	)
}

func (b *jsiiProxy_BuildWorkflow) AddPostBuildSteps(steps ...*workflows.JobStep) {
	if err := b.validateAddPostBuildStepsParameters(&steps); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addPostBuildSteps",
		args,
	)
}

func (b *jsiiProxy_BuildWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"postSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"preSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildWorkflow) Synthesize() {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BuildWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

