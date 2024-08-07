package gitlab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/projen/projen-go/projen"
)

// A GitLab CI for templates that are created and included in the `.gitlab-ci.yml` file.
// Experimental.
type NestedConfiguration interface {
	CiConfiguration
	// Defines default scripts that should run *after* all jobs.
	//
	// Can be overriden by the job level `afterScript`.
	// Experimental.
	DefaultAfterScript() *[]*string
	// Default list of files and directories that should be attached to the job if it succeeds.
	//
	// Artifacts are sent to Gitlab where they can be downloaded.
	// Experimental.
	DefaultArtifacts() *Artifacts
	// Defines default scripts that should run *before* all jobs.
	//
	// Can be overriden by the job level `afterScript`.
	// Experimental.
	DefaultBeforeScript() *[]*string
	// Experimental.
	DefaultCache() *[]*Cache
	// Default ID tokens (JSON Web Tokens) that are used for CI/CD authentication to use globally for all jobs.
	// Experimental.
	DefaultIdTokens() *map[string]IDToken
	// Specifies the default docker image to use globally for all jobs.
	// Experimental.
	DefaultImage() *Image
	// The default behavior for whether a job should be canceled when a newer pipeline starts before the job completes (Default: false).
	// Experimental.
	DefaultInterruptible() *bool
	// How many times a job is retried if it fails.
	//
	// If not defined, defaults to 0 and jobs do not retry.
	// Experimental.
	DefaultRetry() *Retry
	// Used to select a specific runner from the list of all runners that are available for the project.
	// Experimental.
	DefaultTags() *[]*string
	// A default timeout job written in natural language (Ex.
	//
	// one hour, 3600 seconds, 60 minutes).
	// Experimental.
	DefaultTimeout() *string
	// The workflow YAML file.
	// Experimental.
	File() projen.YamlFile
	// The jobs in the CI configuration.
	// Experimental.
	Jobs() *map[string]*Job
	// The name of the configuration.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// A special job used to upload static sites to Gitlab pages.
	//
	// Requires a `public/` directory
	// with `artifacts.path` pointing to it.
	// Experimental.
	Pages() *Job
	// Experimental.
	Parent() GitlabConfiguration
	// Path to CI file generated by the configuration.
	// Experimental.
	Path() *string
	// Experimental.
	Project() projen.Project
	// Groups jobs into stages.
	//
	// All jobs in one stage must complete before next stage is
	// executed. Defaults to ['build', 'test', 'deploy'].
	// Experimental.
	Stages() *[]*string
	// Global variables that are passed to jobs.
	//
	// If the job already has that variable defined, the job-level variable takes precedence.
	// Experimental.
	Variables() *map[string]interface{}
	// Used to control pipeline behavior.
	// Experimental.
	Workflow() *Workflow
	// Adds up to 4 default caches configuration to the CI configuration.
	// Experimental.
	AddDefaultCaches(caches *[]*Cache)
	// Add a globally defined variable to the CI configuration.
	// Experimental.
	AddGlobalVariables(variables *map[string]interface{})
	// Add additional yml/yaml files to the CI includes.
	// Experimental.
	AddIncludes(includes ...*Include)
	// Add jobs and their stages to the CI configuration.
	// Experimental.
	AddJobs(jobs *map[string]*Job)
	// Add additional services.
	// Experimental.
	AddServices(services ...*Service)
	// Add stages to the CI configuration if not already present.
	// Experimental.
	AddStages(stages ...*string)
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

// The jsii proxy struct for NestedConfiguration
type jsiiProxy_NestedConfiguration struct {
	jsiiProxy_CiConfiguration
}

func (j *jsiiProxy_NestedConfiguration) DefaultAfterScript() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAfterScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultArtifacts() *Artifacts {
	var returns *Artifacts
	_jsii_.Get(
		j,
		"defaultArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultBeforeScript() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultBeforeScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultCache() *[]*Cache {
	var returns *[]*Cache
	_jsii_.Get(
		j,
		"defaultCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultIdTokens() *map[string]IDToken {
	var returns *map[string]IDToken
	_jsii_.Get(
		j,
		"defaultIdTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultImage() *Image {
	var returns *Image
	_jsii_.Get(
		j,
		"defaultImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultInterruptible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultInterruptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultRetry() *Retry {
	var returns *Retry
	_jsii_.Get(
		j,
		"defaultRetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) DefaultTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Jobs() *map[string]*Job {
	var returns *map[string]*Job
	_jsii_.Get(
		j,
		"jobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Pages() *Job {
	var returns *Job
	_jsii_.Get(
		j,
		"pages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Parent() GitlabConfiguration {
	var returns GitlabConfiguration
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Variables() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedConfiguration) Workflow() *Workflow {
	var returns *Workflow
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}


// Experimental.
func NewNestedConfiguration(project projen.Project, parent GitlabConfiguration, name *string, options *CiConfigurationOptions) NestedConfiguration {
	_init_.Initialize()

	if err := validateNewNestedConfigurationParameters(project, parent, name, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_NestedConfiguration{}

	_jsii_.Create(
		"projen.gitlab.NestedConfiguration",
		[]interface{}{project, parent, name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNestedConfiguration_Override(n NestedConfiguration, project projen.Project, parent GitlabConfiguration, name *string, options *CiConfigurationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.gitlab.NestedConfiguration",
		[]interface{}{project, parent, name, options},
		n,
	)
}

// Test whether the given construct is a component.
// Experimental.
func NestedConfiguration_IsComponent(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNestedConfiguration_IsComponentParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.gitlab.NestedConfiguration",
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
func NestedConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNestedConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"projen.gitlab.NestedConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedConfiguration) AddDefaultCaches(caches *[]*Cache) {
	if err := n.validateAddDefaultCachesParameters(caches); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addDefaultCaches",
		[]interface{}{caches},
	)
}

func (n *jsiiProxy_NestedConfiguration) AddGlobalVariables(variables *map[string]interface{}) {
	if err := n.validateAddGlobalVariablesParameters(variables); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addGlobalVariables",
		[]interface{}{variables},
	)
}

func (n *jsiiProxy_NestedConfiguration) AddIncludes(includes ...*Include) {
	if err := n.validateAddIncludesParameters(&includes); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range includes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addIncludes",
		args,
	)
}

func (n *jsiiProxy_NestedConfiguration) AddJobs(jobs *map[string]*Job) {
	if err := n.validateAddJobsParameters(jobs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (n *jsiiProxy_NestedConfiguration) AddServices(services ...*Service) {
	if err := n.validateAddServicesParameters(&services); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range services {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addServices",
		args,
	)
}

func (n *jsiiProxy_NestedConfiguration) AddStages(stages ...*string) {
	args := []interface{}{}
	for _, a := range stages {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addStages",
		args,
	)
}

func (n *jsiiProxy_NestedConfiguration) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NestedConfiguration) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NestedConfiguration) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NestedConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

