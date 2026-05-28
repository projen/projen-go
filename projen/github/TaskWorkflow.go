package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/github/workflows"
)

// A GitHub workflow for common build tasks within a project.
// Experimental.
type TaskWorkflow interface {
	GithubWorkflow
	// Experimental.
	ArtifactsDirectory() *string
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
	// Experimental.
	JobId() *string
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
	// Appends a step to the end of a job's step list.
	// Experimental.
	AppendStep(jobId *string, step *workflows.JobStep)
	// Get a single job from the workflow.
	// Experimental.
	GetJob(id *string) interface{}
	// Gets a single step from a job by step ID.
	//
	// The returned object is frozen and read-only. Use `replaceStep` or
	// `patchStep` to modify a step.
	//
	// Returns: A read-only copy of the step.
	// Experimental.
	GetStep(jobId *string, stepId *string) *workflows.JobStep
	// Inserts a step after an existing step, identified by ID.
	// Experimental.
	InsertStepAfter(jobId *string, referenceStepId *string, step *workflows.JobStep)
	// Inserts a step before an existing step, identified by ID.
	// Experimental.
	InsertStepBefore(jobId *string, referenceStepId *string, step *workflows.JobStep)
	// Add events to triggers the workflow.
	// Experimental.
	On(events *workflows.Triggers)
	// Applies a surgical modification to an existing step.
	//
	// The provided patch is shallow-merged onto the existing step. Fields not
	// present in the patch are preserved unchanged. Use `getStep` to read the
	// current step values before constructing the patch.
	//
	// Example:
	//   // Add an env var without replacing existing ones:
	//   const step = workflow.getStep("build", "install");
	//   workflow.patchStep("build", "install", {
	//     env: { ...step.env, NODE_AUTH_TOKEN: "${{ steps.jfrog.outputs.oidc-token }}" },
	//   });
	//
	// Experimental.
	PatchStep(jobId *string, stepId *string, patch *workflows.JobStep)
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
	// Removes a step from a job by step ID.
	// Experimental.
	RemoveStep(jobId *string, stepId *string)
	// Replaces an existing step in a job, preserving its position.
	// Experimental.
	ReplaceStep(jobId *string, stepId *string, replacementStep *workflows.JobStep)
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

func (j *jsiiProxy_TaskWorkflow) Concurrency() *ConcurrencyOptions {
	var returns *ConcurrencyOptions
	_jsii_.Get(
		j,
		"concurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskWorkflow) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
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

func (j *jsiiProxy_TaskWorkflow) Jobs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"jobs",
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

func (j *jsiiProxy_TaskWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

// Test whether the given construct is a component.
// Experimental.
func TaskWorkflow_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaskWorkflow_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.TaskWorkflow",
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
func TaskWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaskWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.github.TaskWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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

func (t *jsiiProxy_TaskWorkflow) AppendStep(jobId *string, step *workflows.JobStep) {
	if err := t.validateAppendStepParameters(jobId, step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"appendStep",
		[]interface{}{jobId, step},
	)
}

func (t *jsiiProxy_TaskWorkflow) GetJob(id *string) interface{} {
	if err := t.validateGetJobParameters(id); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"getJob",
		[]interface{}{id},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskWorkflow) GetStep(jobId *string, stepId *string) *workflows.JobStep {
	if err := t.validateGetStepParameters(jobId, stepId); err != nil {
		panic(err)
	}
	var returns *workflows.JobStep

	_jsii_.Invoke(
		t,
		"getStep",
		[]interface{}{jobId, stepId},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskWorkflow) InsertStepAfter(jobId *string, referenceStepId *string, step *workflows.JobStep) {
	if err := t.validateInsertStepAfterParameters(jobId, referenceStepId, step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"insertStepAfter",
		[]interface{}{jobId, referenceStepId, step},
	)
}

func (t *jsiiProxy_TaskWorkflow) InsertStepBefore(jobId *string, referenceStepId *string, step *workflows.JobStep) {
	if err := t.validateInsertStepBeforeParameters(jobId, referenceStepId, step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"insertStepBefore",
		[]interface{}{jobId, referenceStepId, step},
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

func (t *jsiiProxy_TaskWorkflow) PatchStep(jobId *string, stepId *string, patch *workflows.JobStep) {
	if err := t.validatePatchStepParameters(jobId, stepId, patch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"patchStep",
		[]interface{}{jobId, stepId, patch},
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

func (t *jsiiProxy_TaskWorkflow) RemoveJob(id *string) {
	if err := t.validateRemoveJobParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"removeJob",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TaskWorkflow) RemoveStep(jobId *string, stepId *string) {
	if err := t.validateRemoveStepParameters(jobId, stepId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"removeStep",
		[]interface{}{jobId, stepId},
	)
}

func (t *jsiiProxy_TaskWorkflow) ReplaceStep(jobId *string, stepId *string, replacementStep *workflows.JobStep) {
	if err := t.validateReplaceStepParameters(jobId, stepId, replacementStep); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"replaceStep",
		[]interface{}{jobId, stepId, replacementStep},
	)
}

func (t *jsiiProxy_TaskWorkflow) Synthesize() {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TaskWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskWorkflow) UpdateJob(id *string, job interface{}) {
	if err := t.validateUpdateJobParameters(id, job); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"updateJob",
		[]interface{}{id, job},
	)
}

func (t *jsiiProxy_TaskWorkflow) UpdateJobs(jobs *map[string]interface{}) {
	if err := t.validateUpdateJobsParameters(jobs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"updateJobs",
		[]interface{}{jobs},
	)
}

func (t *jsiiProxy_TaskWorkflow) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

