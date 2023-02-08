package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/internal"
	"github.com/projen/projen-go/projen/github/workflows"
)

// Workflow for GitHub.
//
// A workflow is a configurable automated process made up of one or more jobs.
// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions
//
// Experimental.
type GithubWorkflow interface {
	projen.Component
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	// Experimental.
	Concurrency() *string
	// The workflow YAML file.
	//
	// May not exist if `workflowsEnabled` is false on `GitHub`.
	// Experimental.
	File() projen.YamlFile
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

// The jsii proxy struct for GithubWorkflow
type jsiiProxy_GithubWorkflow struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_GithubWorkflow) Concurrency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) ProjenCredentials() GithubCredentials {
	var returns GithubCredentials
	_jsii_.Get(
		j,
		"projenCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) RunName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runName",
		&returns,
	)
	return returns
}


// Experimental.
func NewGithubWorkflow(github GitHub, name *string, options *GithubWorkflowOptions) GithubWorkflow {
	_init_.Initialize()

	if err := validateNewGithubWorkflowParameters(github, name, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_GithubWorkflow{}

	_jsii_.Create(
		"projen.github.GithubWorkflow",
		[]interface{}{github, name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGithubWorkflow_Override(g GithubWorkflow, github GitHub, name *string, options *GithubWorkflowOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.github.GithubWorkflow",
		[]interface{}{github, name, options},
		g,
	)
}

func (j *jsiiProxy_GithubWorkflow)SetRunName(val *string) {
	_jsii_.Set(
		j,
		"runName",
		val,
	)
}

func (g *jsiiProxy_GithubWorkflow) AddJob(id *string, job interface{}) {
	if err := g.validateAddJobParameters(id, job); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addJob",
		[]interface{}{id, job},
	)
}

func (g *jsiiProxy_GithubWorkflow) AddJobs(jobs *map[string]interface{}) {
	if err := g.validateAddJobsParameters(jobs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (g *jsiiProxy_GithubWorkflow) On(events *workflows.Triggers) {
	if err := g.validateOnParameters(events); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"on",
		[]interface{}{events},
	)
}

func (g *jsiiProxy_GithubWorkflow) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubWorkflow) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubWorkflow) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

