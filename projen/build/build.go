package build

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build/internal"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Experimental.
type BuildWorkflow interface {
	projen.Component
	BuildJobIds() *[]*string
	Project() projen.Project
	AddPostBuildJob(id *string, job *workflows.Job)
	AddPostBuildSteps(steps ...*workflows.JobStep)
	PostSynthesize()
	PreSynthesize()
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

// Adds another job to the build workflow which is executed after the build job succeeded.
// Experimental.
func (b *jsiiProxy_BuildWorkflow) AddPostBuildJob(id *string, job *workflows.Job) {
	_jsii_.InvokeVoid(
		b,
		"addPostBuildJob",
		[]interface{}{id, job},
	)
}

// Adds steps that are executed after the build.
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (b *jsiiProxy_BuildWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (b *jsiiProxy_BuildWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	ArtifactsDirectory *string `json:"artifactsDirectory"`
	// The task to execute in order to build the project.
	// Experimental.
	BuildTask projen.Task `json:"buildTask"`
	// Enable anti-tamper check.
	// Experimental.
	Antitamper *bool `json:"antitamper"`
	// The container image to use for builds.
	// Experimental.
	ContainerImage *string `json:"containerImage"`
	// Build environment variables.
	// Experimental.
	Env *map[string]*string `json:"env"`
	// Git identity to use for the workflow.
	// Experimental.
	GitIdentity *github.GitIdentity `json:"gitIdentity"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Experimental.
	MutableBuild *bool `json:"mutableBuild"`
	// Steps to execute after build.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `json:"postBuildSteps"`
	// Steps to execute before the build.
	// Experimental.
	PreBuildSteps *[]*workflows.JobStep `json:"preBuildSteps"`
}

