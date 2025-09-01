package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
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
	// The concurrency configuration of the workflow.
	//
	// undefined means no concurrency limitations.
	// Experimental.
	Concurrency() *ConcurrencyOptions
	// Additional environment variables to set for the workflow.
	// Experimental.
	Env() *map[string]*string
	// The workflow YAML file.
	//
	// May not exist if `workflowsEnabled` is false on `GitHub`.
	// Experimental.
	File() projen.YamlFile
	// All current jobs of the workflow.
	//
	// This is a read-only copy, use the respective helper methods to add, update or remove jobs.
	// Experimental.
	Jobs() *map[string]interface{}
	// The name of the workflow.
	//
	// GitHub displays the names of your workflows under your repository's
	// "Actions" tab.
	// See: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#name
	//
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Get a single job from the workflow.
	// Experimental.
	GetJob(id *string) interface{}
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
	// Removes a single job to the workflow.
	// Experimental.
	RemoveJob(id *string)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Updates a single job to the workflow.
	// Experimental.
	UpdateJob(id *string, job interface{})
	// Updates jobs for this workflow Does a complete replace, it does not try to merge the jobs.
	// Experimental.
	UpdateJobs(jobs *map[string]interface{})
}

// The jsii proxy struct for GithubWorkflow
type jsiiProxy_GithubWorkflow struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_GithubWorkflow) Concurrency() *ConcurrencyOptions {
	var returns *ConcurrencyOptions
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubWorkflow) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
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

func (j *jsiiProxy_GithubWorkflow) Jobs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"jobs",
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

func (j *jsiiProxy_GithubWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

// Test whether the given construct is a component.
// Experimental.
func GithubWorkflow_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGithubWorkflow_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.GithubWorkflow",
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
func GithubWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGithubWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.GithubWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

func (g *jsiiProxy_GithubWorkflow) GetJob(id *string) interface{} {
	if err := g.validateGetJobParameters(id); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"getJob",
		[]interface{}{id},
		&returns,
	)

	return returns
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

func (g *jsiiProxy_GithubWorkflow) RemoveJob(id *string) {
	if err := g.validateRemoveJobParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"removeJob",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GithubWorkflow) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GithubWorkflow) UpdateJob(id *string, job interface{}) {
	if err := g.validateUpdateJobParameters(id, job); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"updateJob",
		[]interface{}{id, job},
	)
}

func (g *jsiiProxy_GithubWorkflow) UpdateJobs(jobs *map[string]interface{}) {
	if err := g.validateUpdateJobsParameters(jobs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"updateJobs",
		[]interface{}{jobs},
	)
}

