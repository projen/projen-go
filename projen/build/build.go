package build

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build/internal"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Options for `BuildWorkflow.addPostBuildJobCommands`.
// Experimental.
type AddPostBuildJobCommandsOptions struct {
	// Check out the repository at the pull request branch before commands are run.
	// Experimental.
	CheckoutRepo *bool `json:"checkoutRepo" yaml:"checkoutRepo"`
	// Install project dependencies before running commands. `checkoutRepo` must also be set to true.
	//
	// Currently only supported for `NodeProject`.
	// Experimental.
	InstallDeps *bool `json:"installDeps" yaml:"installDeps"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
	// Tools that should be installed before the commands are run.
	// Experimental.
	Tools *workflows.Tools `json:"tools" yaml:"tools"`
}

// Options for `BuildWorkflow.addPostBuildJobTask`.
// Experimental.
type AddPostBuildJobTaskOptions struct {
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
	// Tools that should be installed before the task is run.
	// Experimental.
	Tools *workflows.Tools `json:"tools" yaml:"tools"`
}

// Experimental.
type BuildWorkflow interface {
	projen.Component
	// Returns a list of job IDs that are part of the build.
	// Experimental.
	BuildJobIds() *[]*string
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

func (b *jsiiProxy_BuildWorkflow) AddPostBuildJob(id *string, job *workflows.Job) {
	_jsii_.InvokeVoid(
		b,
		"addPostBuildJob",
		[]interface{}{id, job},
	)
}

func (b *jsiiProxy_BuildWorkflow) AddPostBuildJobCommands(id *string, commands *[]*string, options *AddPostBuildJobCommandsOptions) {
	_jsii_.InvokeVoid(
		b,
		"addPostBuildJobCommands",
		[]interface{}{id, commands, options},
	)
}

func (b *jsiiProxy_BuildWorkflow) AddPostBuildJobTask(task projen.Task, options *AddPostBuildJobTaskOptions) {
	_jsii_.InvokeVoid(
		b,
		"addPostBuildJobTask",
		[]interface{}{task, options},
	)
}

func (b *jsiiProxy_BuildWorkflow) AddPostBuildSteps(steps ...*workflows.JobStep) {
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

// Experimental.
type BuildWorkflowOptions struct {
	// A name of a directory that includes build artifacts.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// The task to execute in order to build the project.
	// Experimental.
	BuildTask projen.Task `json:"buildTask" yaml:"buildTask"`
	// The container image to use for builds.
	// Experimental.
	ContainerImage *string `json:"containerImage" yaml:"containerImage"`
	// Build environment variables.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// Git identity to use for the workflow.
	// Experimental.
	GitIdentity *github.GitIdentity `json:"gitIdentity" yaml:"gitIdentity"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means that any files synthesized by projen or e.g. test snapshots will
	// always be up-to-date before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	//
	// This is enabled by default only if `githubTokenSecret` is set. Otherwise it
	// is disabled, which implies that file changes that happen during build will
	// not be pushed back to the branch.
	// Experimental.
	MutableBuild *bool `json:"mutableBuild" yaml:"mutableBuild"`
	// Steps to execute after build.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `json:"postBuildSteps" yaml:"postBuildSteps"`
	// Steps to execute before the build.
	// Experimental.
	PreBuildSteps *[]*workflows.JobStep `json:"preBuildSteps" yaml:"preBuildSteps"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
	// Build workflow triggers.
	// Experimental.
	WorkflowTriggers *workflows.Triggers `json:"workflowTriggers" yaml:"workflowTriggers"`
}

