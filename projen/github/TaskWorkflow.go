package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
)

// A GitHub workflow for common build tasks within a project.
// Experimental.
type TaskWorkflow interface {
	GithubWorkflow
	// Experimental.
	ArtifactsDirectory() *string
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	// Experimental.
	Concurrency() *string
	// The workflow YAML file.
	//
	// May not exist if `workflowsEnabled` is false on `GitHub`.
	// Experimental.
	File() projen.YamlFile
	// Experimental.
	JobId() *string
	// The name of the workflow.
	// Experimental.
	Name() *string
	// Experimental.
	Project() projen.Project
	// GitHub API authentication method used by projen workflows.
	// Experimental.
	ProjenCredentials() GithubCredentials
	// The name for workflow runs generated from the workflow.
	//
	// GitHub displays the
	// workflow run name in the list of workflow runs on your repository's
	// "Actions" tab. If `run-name` is omitted or is only whitespace, then the run
	// name is set to event-specific information for the workflow run. For
	// example, for a workflow triggered by a `push` or `pull_request` event, it
	// is set as the commit message.
	//
	// This value can include expressions and can reference `github` and `inputs`
	// contexts.
	// Experimental.
	RunName() *string
	// Experimental.
	SetRunName(val *string)
	// Adds a single job to the workflow.
	// Experimental.
	AddJob(id *string, job interface{})
	// Add jobs to the workflow.
	// Experimental.
	AddJobs(jobs *map[string]interface{})
	// Add events to triggers the workflow.
	// Experimental.
	On(events *workflows.Triggers)
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

// The jsii proxy struct for TaskWorkflow
type jsiiProxy_TaskWorkflow struct {
	jsiiProxy_GithubWorkflow
}

func (j *jsiiProxy_TaskWorkflow) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) Concurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) ProjenCredentials() GithubCredentials {
	var returns GithubCredentials
	_jsii_.Get(
		j,
		"projenCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) RunName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTaskWorkflow(github GitHub, options *TaskWorkflowOptions) TaskWorkflow {
	_init_.Initialize()

	if err := validateNewTaskWorkflowParameters(github, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaskWorkflow{}

	_jsii_.Create(
		"projen.github.TaskWorkflow",
		[]interface{}{github, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTaskWorkflow_Override(t TaskWorkflow, github GitHub, options *TaskWorkflowOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.TaskWorkflow",
		[]interface{}{github, options},
		t,
	)
}

func (j *jsiiProxy_TaskWorkflow)SetRunName(val *string) {
	_jsii_.Set(
		j,
		"runName",
		val,
	)
}

func (t *jsiiProxy_TaskWorkflow) AddJob(id *string, job interface{}) {
	if err := t.validateAddJobParameters(id, job); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addJob",
		[]interface{}{id, job},
	)
}

func (t *jsiiProxy_TaskWorkflow) AddJobs(jobs *map[string]interface{}) {
	if err := t.validateAddJobsParameters(jobs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (t *jsiiProxy_TaskWorkflow) On(events *workflows.Triggers) {
	if err := t.validateOnParameters(events); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"on",
		[]interface{}{events},
	)
}

func (t *jsiiProxy_TaskWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"postSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		t,
		"preSynthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskWorkflow) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

