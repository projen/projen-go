package gitlab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/gitlab/internal"
)

// Specifies what this job will do.
//
// 'start' (default) indicates the job will start the
// deployment. 'prepare' indicates this will not affect the deployment. 'stop' indicates
// this will stop the deployment.
// Experimental.
type Action string

const (
	// Experimental.
	Action_PREPARE Action = "PREPARE"
	// Experimental.
	Action_START Action = "START"
	// Experimental.
	Action_STOP Action = "STOP"
)

// Exit code that are not considered failure.
//
// The job fails for any other exit code.
// You can list which exit codes are not considered failures. The job fails for any other
// exit code.
// See: https://docs.gitlab.com/ee/ci/yaml/#allow_failure
//
// Experimental.
type AllowFailure struct {
	// Experimental.
	ExitCodes interface{} `json:"exitCodes" yaml:"exitCodes"`
}

// Used to specify a list of files and directories that should be attached to the job if it succeeds.
//
// Artifacts are sent to Gitlab where they can be downloaded.
// See: https://docs.gitlab.com/ee/ci/yaml/#artifacts
//
// Experimental.
type Artifacts struct {
	// A list of paths to files/folders that should be excluded in the artifact.
	// Experimental.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// How long artifacts should be kept.
	//
	// They are saved 30 days by default. Artifacts that have expired are removed periodically via cron job. Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'.
	// Experimental.
	ExpireIn *string `json:"expireIn" yaml:"expireIn"`
	// Can be used to expose job artifacts in the merge request UI.
	//
	// GitLab will add a link <expose_as> to the relevant merge request that points to the artifact.
	// Experimental.
	ExposeAs *string `json:"exposeAs" yaml:"exposeAs"`
	// Name for the archive created on job success.
	//
	// Can use variables in the name, e.g. '$CI_JOB_NAME'
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// A list of paths to files/folders that should be included in the artifact.
	// Experimental.
	Paths *[]*string `json:"paths" yaml:"paths"`
	// Reports will be uploaded as artifacts, and often displayed in the Gitlab UI, such as in Merge Requests.
	// Experimental.
	Reports *Reports `json:"reports" yaml:"reports"`
	// Whether to add all untracked files (along with 'artifacts.paths') to the artifact.
	// Experimental.
	Untracked *bool `json:"untracked" yaml:"untracked"`
	// Configure when artifacts are uploaded depended on job status.
	// Experimental.
	When CacheWhen `json:"when" yaml:"when"`
}

// Asset configuration for a release.
// Experimental.
type Assets struct {
	// Include asset links in the release.
	// Experimental.
	Links *[]*Link `json:"links" yaml:"links"`
}

// Cache Definition.
// See: https://docs.gitlab.com/ee/ci/yaml/#cache
//
// Experimental.
type Cache struct {
	// Defines when to save the cache, based on the status of the job (Default: Job Success).
	// Experimental.
	When CacheWhen `json:"when" yaml:"when"`
}

// Configure when artifacts are uploaded depended on job status.
// See: https://docs.gitlab.com/ee/ci/yaml/#cachewhen
//
// Experimental.
type CacheWhen string

const (
	// Upload artifacts regardless of job status.
	// Experimental.
	CacheWhen_ALWAYS CacheWhen = "ALWAYS"
	// Upload artifacts only when the job fails.
	// Experimental.
	CacheWhen_ON_FAILURE CacheWhen = "ON_FAILURE"
	// Upload artifacts only when the job succeeds (this is the default).
	// Experimental.
	CacheWhen_ON_SUCCESS CacheWhen = "ON_SUCCESS"
)

// CI for GitLab.
//
// A CI is a configurable automated process made up of one or more stages/jobs.
// See: https://docs.gitlab.com/ee/ci/yaml/
//
// Experimental.
type CiConfiguration interface {
	projen.Component
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
	// A default list of files and directories to cache between jobs.
	//
	// You can only use paths that are in the local working copy.
	// Experimental.
	DefaultCache() *Cache
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
	// A special job used to upload static sites to Gitlab pages.
	//
	// Requires a `public/` directory
	// with `artifacts.path` pointing to it.
	// Experimental.
	Pages() *Job
	// Path to CI file generated by the configuration.
	// Experimental.
	Path() *string
	// The project the configuration belongs to.
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
}

// The jsii proxy struct for CiConfiguration
type jsiiProxy_CiConfiguration struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_CiConfiguration) DefaultAfterScript() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAfterScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultArtifacts() *Artifacts {
	var returns *Artifacts
	_jsii_.Get(
		j,
		"defaultArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultBeforeScript() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultBeforeScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultCache() *Cache {
	var returns *Cache
	_jsii_.Get(
		j,
		"defaultCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultImage() *Image {
	var returns *Image
	_jsii_.Get(
		j,
		"defaultImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultInterruptible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultInterruptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultRetry() *Retry {
	var returns *Retry
	_jsii_.Get(
		j,
		"defaultRetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) DefaultTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Jobs() *map[string]*Job {
	var returns *map[string]*Job
	_jsii_.Get(
		j,
		"jobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Pages() *Job {
	var returns *Job
	_jsii_.Get(
		j,
		"pages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Variables() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CiConfiguration) Workflow() *Workflow {
	var returns *Workflow
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}


// Experimental.
func NewCiConfiguration(project projen.Project, name *string, options *CiConfigurationOptions) CiConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CiConfiguration{}

	_jsii_.Create(
		"projen.gitlab.CiConfiguration",
		[]interface{}{project, name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCiConfiguration_Override(c CiConfiguration, project projen.Project, name *string, options *CiConfigurationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.gitlab.CiConfiguration",
		[]interface{}{project, name, options},
		c,
	)
}

func (c *jsiiProxy_CiConfiguration) AddGlobalVariables(variables *map[string]interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addGlobalVariables",
		[]interface{}{variables},
	)
}

func (c *jsiiProxy_CiConfiguration) AddIncludes(includes ...*Include) {
	args := []interface{}{}
	for _, a := range includes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addIncludes",
		args,
	)
}

func (c *jsiiProxy_CiConfiguration) AddJobs(jobs *map[string]*Job) {
	_jsii_.InvokeVoid(
		c,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (c *jsiiProxy_CiConfiguration) AddServices(services ...*Service) {
	args := []interface{}{}
	for _, a := range services {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addServices",
		args,
	)
}

func (c *jsiiProxy_CiConfiguration) AddStages(stages ...*string) {
	args := []interface{}{}
	for _, a := range stages {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addStages",
		args,
	)
}

func (c *jsiiProxy_CiConfiguration) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CiConfiguration) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CiConfiguration) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `CiConfiguration`.
// Experimental.
type CiConfigurationOptions struct {
	// Default settings for the CI Configuration.
	//
	// Jobs that do not define one or more of the listed keywords use the value defined in the default section.
	// Experimental.
	Default *Default `json:"default" yaml:"default"`
	// An initial set of jobs to add to the configuration.
	// Experimental.
	Jobs *map[string]*Job `json:"jobs" yaml:"jobs"`
	// A special job used to upload static sites to Gitlab pages.
	//
	// Requires a `public/` directory
	// with `artifacts.path` pointing to it.
	// Experimental.
	Pages *Job `json:"pages" yaml:"pages"`
	// Groups jobs into stages.
	//
	// All jobs in one stage must complete before next stage is
	// executed. If no stages are specified. Defaults to ['build', 'test', 'deploy'].
	// Experimental.
	Stages *[]*string `json:"stages" yaml:"stages"`
	// Global variables that are passed to jobs.
	//
	// If the job already has that variable defined, the job-level variable takes precedence.
	// Experimental.
	Variables *map[string]interface{} `json:"variables" yaml:"variables"`
	// Used to control pipeline behavior.
	// Experimental.
	Workflow *Workflow `json:"workflow" yaml:"workflow"`
}

// Default settings for the CI Configuration.
//
// Jobs that do not define one or more of the listed keywords use the value defined in the default section.
// See: https://docs.gitlab.com/ee/ci/yaml/#default
//
// Experimental.
type Default struct {
	// Experimental.
	AfterScript *[]*string `json:"afterScript" yaml:"afterScript"`
	// Experimental.
	Artifacts *Artifacts `json:"artifacts" yaml:"artifacts"`
	// Experimental.
	BeforeScript *[]*string `json:"beforeScript" yaml:"beforeScript"`
	// Experimental.
	Cache *Cache `json:"cache" yaml:"cache"`
	// Experimental.
	Image *Image `json:"image" yaml:"image"`
	// Experimental.
	Interruptible *bool `json:"interruptible" yaml:"interruptible"`
	// Experimental.
	Retry *Retry `json:"retry" yaml:"retry"`
	// Experimental.
	Services *[]*Service `json:"services" yaml:"services"`
	// Experimental.
	Tags *[]*string `json:"tags" yaml:"tags"`
	// Experimental.
	Timeout *string `json:"timeout" yaml:"timeout"`
}

// Experimental.
type DefaultElement string

const (
	// Experimental.
	DefaultElement_AFTER_SCRIPT DefaultElement = "AFTER_SCRIPT"
	// Experimental.
	DefaultElement_ARTIFACTS DefaultElement = "ARTIFACTS"
	// Experimental.
	DefaultElement_BEFORE_SCRIPT DefaultElement = "BEFORE_SCRIPT"
	// Experimental.
	DefaultElement_CACHE DefaultElement = "CACHE"
	// Experimental.
	DefaultElement_IMAGE DefaultElement = "IMAGE"
	// Experimental.
	DefaultElement_INTERRUPTIBLE DefaultElement = "INTERRUPTIBLE"
	// Experimental.
	DefaultElement_RETRY DefaultElement = "RETRY"
	// Experimental.
	DefaultElement_SERVICES DefaultElement = "SERVICES"
	// Experimental.
	DefaultElement_TAGS DefaultElement = "TAGS"
	// Experimental.
	DefaultElement_TIMEOUT DefaultElement = "TIMEOUT"
)

// Explicitly specifies the tier of the deployment environment if non-standard environment name is used.
// Experimental.
type DeploymentTier string

const (
	// Experimental.
	DeploymentTier_DEVELOPMENT DeploymentTier = "DEVELOPMENT"
	// Experimental.
	DeploymentTier_OTHER DeploymentTier = "OTHER"
	// Experimental.
	DeploymentTier_PRODUCTION DeploymentTier = "PRODUCTION"
	// Experimental.
	DeploymentTier_STAGING DeploymentTier = "STAGING"
	// Experimental.
	DeploymentTier_TESTING DeploymentTier = "TESTING"
)

// The engine configuration for a secret.
// Experimental.
type Engine struct {
	// Name of the secrets engine.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Path to the secrets engine.
	// Experimental.
	Path *string `json:"path" yaml:"path"`
}

// The environment that a job deploys to.
// Experimental.
type Environment struct {
	// The name of the environment, e.g. 'qa', 'staging', 'production'.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Specifies what this job will do.
	//
	// 'start' (default) indicates the job will start the deployment. 'prepare' indicates this will not affect the deployment. 'stop' indicates this will stop the deployment.
	// Experimental.
	Action Action `json:"action" yaml:"action"`
	// The amount of time it should take before Gitlab will automatically stop the environment.
	//
	// Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'.
	// Experimental.
	AutoStopIn *string `json:"autoStopIn" yaml:"autoStopIn"`
	// Explicitly specifies the tier of the deployment environment if non-standard environment name is used.
	// Experimental.
	DeploymentTier DeploymentTier `json:"deploymentTier" yaml:"deploymentTier"`
	// Used to configure the kubernetes deployment for this environment.
	//
	// This is currently not supported for kubernetes clusters that are managed by Gitlab.
	// Experimental.
	Kubernetes *KubernetesConfig `json:"kubernetes" yaml:"kubernetes"`
	// The name of a job to execute when the environment is about to be stopped.
	// Experimental.
	OnStop *string `json:"onStop" yaml:"onStop"`
	// When set, this will expose buttons in various places for the current environment in Gitlab, that will take you to the defined URL.
	// Experimental.
	Url *string `json:"url" yaml:"url"`
}

// Filtering options for when a job will run.
// Experimental.
type Filter struct {
	// Filter job creation based on files that were modified in a git push.
	// Experimental.
	Changes *[]*string `json:"changes" yaml:"changes"`
	// Filter job based on if Kubernetes integration is active.
	// Experimental.
	Kubernetes KubernetesEnum `json:"kubernetes" yaml:"kubernetes"`
	// Control when to add jobs to a pipeline based on branch names or pipeline types.
	// Experimental.
	Refs *[]*string `json:"refs" yaml:"refs"`
	// Filter job by checking comparing values of environment variables.
	//
	// Read more about variable expressions: https://docs.gitlab.com/ee/ci/variables/README.html#variables-expressions
	// Experimental.
	Variables *[]*string `json:"variables" yaml:"variables"`
}

// A GitLab CI for the main `.gitlab-ci.yml` file.
// Experimental.
type GitlabConfiguration interface {
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
	// A default list of files and directories to cache between jobs.
	//
	// You can only use paths that are in the local working copy.
	// Experimental.
	DefaultCache() *Cache
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
	// Experimental.
	NestedTemplates() *map[string]NestedConfiguration
	// A special job used to upload static sites to Gitlab pages.
	//
	// Requires a `public/` directory
	// with `artifacts.path` pointing to it.
	// Experimental.
	Pages() *Job
	// Path to CI file generated by the configuration.
	// Experimental.
	Path() *string
	// The project the configuration belongs to.
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
	// Creates and adds nested templates to the includes of the main CI.
	//
	// Additionally adds their stages to the main CI if they are not already present.
	// You can futher customize nested templates through the `nestedTemplates` property.
	// E.g. gitlabConfig.nestedTemplates['templateName']?.addStages('stageName')
	// Experimental.
	CreateNestedTemplates(config *map[string]*CiConfigurationOptions)
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

// The jsii proxy struct for GitlabConfiguration
type jsiiProxy_GitlabConfiguration struct {
	jsiiProxy_CiConfiguration
}

func (j *jsiiProxy_GitlabConfiguration) DefaultAfterScript() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultAfterScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultArtifacts() *Artifacts {
	var returns *Artifacts
	_jsii_.Get(
		j,
		"defaultArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultBeforeScript() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultBeforeScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultCache() *Cache {
	var returns *Cache
	_jsii_.Get(
		j,
		"defaultCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultImage() *Image {
	var returns *Image
	_jsii_.Get(
		j,
		"defaultImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultInterruptible() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"defaultInterruptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultRetry() *Retry {
	var returns *Retry
	_jsii_.Get(
		j,
		"defaultRetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) DefaultTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Jobs() *map[string]*Job {
	var returns *map[string]*Job
	_jsii_.Get(
		j,
		"jobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) NestedTemplates() *map[string]NestedConfiguration {
	var returns *map[string]NestedConfiguration
	_jsii_.Get(
		j,
		"nestedTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Pages() *Job {
	var returns *Job
	_jsii_.Get(
		j,
		"pages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Stages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Variables() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitlabConfiguration) Workflow() *Workflow {
	var returns *Workflow
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitlabConfiguration(project projen.Project, options *CiConfigurationOptions) GitlabConfiguration {
	_init_.Initialize()

	j := jsiiProxy_GitlabConfiguration{}

	_jsii_.Create(
		"projen.gitlab.GitlabConfiguration",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGitlabConfiguration_Override(g GitlabConfiguration, project projen.Project, options *CiConfigurationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.gitlab.GitlabConfiguration",
		[]interface{}{project, options},
		g,
	)
}

func (g *jsiiProxy_GitlabConfiguration) AddGlobalVariables(variables *map[string]interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addGlobalVariables",
		[]interface{}{variables},
	)
}

func (g *jsiiProxy_GitlabConfiguration) AddIncludes(includes ...*Include) {
	args := []interface{}{}
	for _, a := range includes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addIncludes",
		args,
	)
}

func (g *jsiiProxy_GitlabConfiguration) AddJobs(jobs *map[string]*Job) {
	_jsii_.InvokeVoid(
		g,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (g *jsiiProxy_GitlabConfiguration) AddServices(services ...*Service) {
	args := []interface{}{}
	for _, a := range services {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addServices",
		args,
	)
}

func (g *jsiiProxy_GitlabConfiguration) AddStages(stages ...*string) {
	args := []interface{}{}
	for _, a := range stages {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addStages",
		args,
	)
}

func (g *jsiiProxy_GitlabConfiguration) CreateNestedTemplates(config *map[string]*CiConfigurationOptions) {
	_jsii_.InvokeVoid(
		g,
		"createNestedTemplates",
		[]interface{}{config},
	)
}

func (g *jsiiProxy_GitlabConfiguration) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabConfiguration) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GitlabConfiguration) Synthesize() {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		nil, // no parameters
	)
}

// Specifies the docker image to use for the job or globally for all jobs.
//
// Job configuration
// takes precedence over global setting. Requires a certain kind of Gitlab runner executor.
// See: https://docs.gitlab.com/ee/ci/yaml/#image
//
// Experimental.
type Image struct {
	// Full name of the image that should be used.
	//
	// It should contain the Registry part if needed.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Command or script that should be executed as the container's entrypoint.
	//
	// It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array.
	// Experimental.
	Entrypoint *[]interface{} `json:"entrypoint" yaml:"entrypoint"`
}

// An included YAML file.
// See: https://docs.gitlab.com/ee/ci/yaml/#include
//
// Experimental.
type Include struct {
	// Files from another private project on the same GitLab instance.
	//
	// You can use `file` in combination with `project` only.
	// Experimental.
	File *[]*string `json:"file" yaml:"file"`
	// Relative path from local repository root (`/`) to the `yaml`/`yml` file template.
	//
	// The file must be on the same branch, and does not work across git submodules.
	// Experimental.
	Local *string `json:"local" yaml:"local"`
	// Path to the project, e.g. `group/project`, or `group/sub-group/project`.
	// Experimental.
	Project *string `json:"project" yaml:"project"`
	// Branch/Tag/Commit-hash for the target project.
	// Experimental.
	Ref *string `json:"ref" yaml:"ref"`
	// URL to a `yaml`/`yml` template file using HTTP/HTTPS.
	// Experimental.
	Remote *string `json:"remote" yaml:"remote"`
	// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
	// Experimental.
	Rules *[]*IncludeRule `json:"rules" yaml:"rules"`
	// Use a `.gitlab-ci.yml` template as a base, e.g. `Nodejs.gitlab-ci.yml`.
	// Experimental.
	Template *string `json:"template" yaml:"template"`
}

// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
// See: https://docs.gitlab.com/ee/ci/yaml/includes.html#use-rules-with-include
//
// Experimental.
type IncludeRule struct {
	// Experimental.
	AllowFailure interface{} `json:"allowFailure" yaml:"allowFailure"`
	// Experimental.
	Changes *[]*string `json:"changes" yaml:"changes"`
	// Experimental.
	Exists *[]*string `json:"exists" yaml:"exists"`
	// Experimental.
	If *string `json:"if" yaml:"if"`
	// Experimental.
	StartIn *string `json:"startIn" yaml:"startIn"`
	// Experimental.
	Variables *map[string]interface{} `json:"variables" yaml:"variables"`
	// Experimental.
	When JobWhen `json:"when" yaml:"when"`
}

// Controls inheritance of globally-defined defaults and variables.
//
// Boolean values control
// inheritance of all default: or variables: keywords. To inherit only a subset of default:
// or variables: keywords, specify what you wish to inherit. Anything not listed is not
// inherited.
// Experimental.
type Inherit struct {
	// Whether to inherit all globally-defined defaults or not.
	//
	// Or subset of inherited defaults.
	// Experimental.
	Default interface{} `json:"default" yaml:"default"`
	// Whether to inherit all globally-defined variables or not.
	//
	// Or subset of inherited variables.
	// Experimental.
	Variables interface{} `json:"variables" yaml:"variables"`
}

// Jobs are the most fundamental element of a .gitlab-ci.yml file.
// See: https://docs.gitlab.com/ee/ci/jobs/
//
// Experimental.
type Job struct {
	// Experimental.
	AfterScript *[]*string `json:"afterScript" yaml:"afterScript"`
	// Whether to allow the pipeline to continue running on job failure (Default: false).
	// Experimental.
	AllowFailure interface{} `json:"allowFailure" yaml:"allowFailure"`
	// Experimental.
	Artifacts *Artifacts `json:"artifacts" yaml:"artifacts"`
	// Experimental.
	BeforeScript *[]*string `json:"beforeScript" yaml:"beforeScript"`
	// Experimental.
	Cache *Cache `json:"cache" yaml:"cache"`
	// Must be a regular expression, optionally but recommended to be quoted, and must be surrounded with '/'.
	//
	// Example: '/Code coverage: \d+\.\d+/'
	// Experimental.
	Coverage *string `json:"coverage" yaml:"coverage"`
	// Specify a list of job names from earlier stages from which artifacts should be loaded.
	//
	// By default, all previous artifacts are passed. Use an empty array to skip downloading artifacts.
	// Experimental.
	Dependencies *[]*string `json:"dependencies" yaml:"dependencies"`
	// Used to associate environment metadata with a deploy.
	//
	// Environment can have a name and URL attached to it, and will be displayed under /environments under the project.
	// Experimental.
	Environment interface{} `json:"environment" yaml:"environment"`
	// Job will run *except* for when these filtering options match.
	// Experimental.
	Except interface{} `json:"except" yaml:"except"`
	// The name of one or more jobs to inherit configuration from.
	// Experimental.
	Extends *[]*string `json:"extends" yaml:"extends"`
	// Experimental.
	Image *Image `json:"image" yaml:"image"`
	// Controls inheritance of globally-defined defaults and variables.
	//
	// Boolean values control inheritance of all default: or variables: keywords. To inherit only a subset of default: or variables: keywords, specify what you wish to inherit. Anything not listed is not inherited.
	// Experimental.
	Inherit *Inherit `json:"inherit" yaml:"inherit"`
	// Experimental.
	Interruptible *bool `json:"interruptible" yaml:"interruptible"`
	// The list of jobs in previous stages whose sole completion is needed to start the current job.
	// Experimental.
	Needs *[]interface{} `json:"needs" yaml:"needs"`
	// Job will run *only* when these filtering options match.
	// Experimental.
	Only interface{} `json:"only" yaml:"only"`
	// Parallel will split up a single job into several, and provide `CI_NODE_INDEX` and `CI_NODE_TOTAL` environment variables for the running jobs.
	// Experimental.
	Parallel interface{} `json:"parallel" yaml:"parallel"`
	// Indicates that the job creates a Release.
	// Experimental.
	Release *Release `json:"release" yaml:"release"`
	// Limit job concurrency.
	//
	// Can be used to ensure that the Runner will not run certain jobs simultaneously.
	// Experimental.
	ResourceGroup *string `json:"resourceGroup" yaml:"resourceGroup"`
	// Experimental.
	Retry *Retry `json:"retry" yaml:"retry"`
	// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
	// Experimental.
	Rules *[]*IncludeRule `json:"rules" yaml:"rules"`
	// Shell scripts executed by the Runner.
	//
	// The only required property of jobs. Be careful with special characters (e.g. `:`, `{`, `}`, `&`) and use single or double quotes to avoid issues.
	// Experimental.
	Script *[]*string `json:"script" yaml:"script"`
	// CI/CD secrets.
	// Experimental.
	Secrets *map[string]*map[string]*Secret `json:"secrets" yaml:"secrets"`
	// Experimental.
	Services *[]*Service `json:"services" yaml:"services"`
	// Define what stage the job will run in.
	// Experimental.
	Stage *string `json:"stage" yaml:"stage"`
	// Experimental.
	StartIn *string `json:"startIn" yaml:"startIn"`
	// Experimental.
	Tags *[]*string `json:"tags" yaml:"tags"`
	// Experimental.
	Timeout *string `json:"timeout" yaml:"timeout"`
	// Trigger allows you to define downstream pipeline trigger.
	//
	// When a job created from trigger definition is started by GitLab, a downstream pipeline gets created. Read more: https://docs.gitlab.com/ee/ci/yaml/README.html#trigger
	// Experimental.
	Trigger interface{} `json:"trigger" yaml:"trigger"`
	// Configurable values that are passed to the Job.
	// Experimental.
	Variables *map[string]interface{} `json:"variables" yaml:"variables"`
	// Describes the conditions for when to run the job.
	//
	// Defaults to 'on_success'.
	// Experimental.
	When JobWhen `json:"when" yaml:"when"`
}

// Describes the conditions for when to run the job.
//
// Defaults to 'on_success'.
// See: https://docs.gitlab.com/ee/ci/yaml/#when
//
// Experimental.
type JobWhen string

const (
	// Experimental.
	JobWhen_ALWAYS JobWhen = "ALWAYS"
	// Experimental.
	JobWhen_DELAYED JobWhen = "DELAYED"
	// Experimental.
	JobWhen_MANUAL JobWhen = "MANUAL"
	// Experimental.
	JobWhen_NEVER JobWhen = "NEVER"
	// Experimental.
	JobWhen_ON_FAILURE JobWhen = "ON_FAILURE"
	// Experimental.
	JobWhen_ON_SUCCESS JobWhen = "ON_SUCCESS"
)

// Used to configure the kubernetes deployment for this environment.
//
// This is currently not
// supported for kubernetes clusters that are managed by Gitlab.
// Experimental.
type KubernetesConfig struct {
	// The kubernetes namespace where this environment should be deployed to.
	// Experimental.
	Namespace *string `json:"namespace" yaml:"namespace"`
}

// Filter job based on if Kubernetes integration is active.
// Experimental.
type KubernetesEnum string

const (
	// Experimental.
	KubernetesEnum_ACTIVE KubernetesEnum = "ACTIVE"
)

// Link configuration for an asset.
// Experimental.
type Link struct {
	// The name of the link.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// The URL to download a file.
	// Experimental.
	Url *string `json:"url" yaml:"url"`
	// The redirect link to the url.
	// Experimental.
	Filepath *string `json:"filepath" yaml:"filepath"`
	// The content kind of what users can download via url.
	// Experimental.
	LinkType LinkType `json:"linkType" yaml:"linkType"`
}

// The content kind of what users can download via url.
// Experimental.
type LinkType string

const (
	// Experimental.
	LinkType_IMAGE LinkType = "IMAGE"
	// Experimental.
	LinkType_OTHER LinkType = "OTHER"
	// Experimental.
	LinkType_PACKAGE LinkType = "PACKAGE"
	// Experimental.
	LinkType_RUNBOOK LinkType = "RUNBOOK"
)

// A jobs in a previous stage whose sole completion is needed to start the current job.
// Experimental.
type Need struct {
	// Experimental.
	Job *string `json:"job" yaml:"job"`
	// Experimental.
	Artifacts *bool `json:"artifacts" yaml:"artifacts"`
	// Experimental.
	Optional *bool `json:"optional" yaml:"optional"`
	// Experimental.
	Pipeline *string `json:"pipeline" yaml:"pipeline"`
	// Experimental.
	Project *string `json:"project" yaml:"project"`
	// Experimental.
	Ref *string `json:"ref" yaml:"ref"`
}

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
	// A default list of files and directories to cache between jobs.
	//
	// You can only use paths that are in the local working copy.
	// Experimental.
	DefaultCache() *Cache
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
	// The project the configuration belongs to.
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

func (j *jsiiProxy_NestedConfiguration) DefaultCache() *Cache {
	var returns *Cache
	_jsii_.Get(
		j,
		"defaultCache",
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

func (n *jsiiProxy_NestedConfiguration) AddGlobalVariables(variables *map[string]interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addGlobalVariables",
		[]interface{}{variables},
	)
}

func (n *jsiiProxy_NestedConfiguration) AddIncludes(includes ...*Include) {
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
	_jsii_.InvokeVoid(
		n,
		"addJobs",
		[]interface{}{jobs},
	)
}

func (n *jsiiProxy_NestedConfiguration) AddServices(services ...*Service) {
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

// Used to run a job multiple times in parallel in a single pipeline.
// Experimental.
type Parallel struct {
	// Defines different variables for jobs that are running in parallel.
	// Experimental.
	Matrix *[]*map[string]*[]interface{} `json:"matrix" yaml:"matrix"`
}

// Indicates that the job creates a Release.
// Experimental.
type Release struct {
	// Specifies the longer description of the Release.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// The tag_name must be specified.
	//
	// It can refer to an existing Git tag or can be specified by the user.
	// Experimental.
	TagName *string `json:"tagName" yaml:"tagName"`
	// Experimental.
	Assets *Assets `json:"assets" yaml:"assets"`
	// The title of each milestone the release is associated with.
	// Experimental.
	Milestones *[]*string `json:"milestones" yaml:"milestones"`
	// The Release name.
	//
	// If omitted, it is populated with the value of release: tag_name.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// If the release: tag_name doesn’t exist yet, the release is created from ref.
	//
	// ref can be a commit SHA, another tag name, or a branch name.
	// Experimental.
	Ref *string `json:"ref" yaml:"ref"`
	// The date and time when the release is ready.
	//
	// Defaults to the current date and time if not defined. Should be enclosed in quotes and expressed in ISO 8601 format.
	// Experimental.
	ReleasedAt *string `json:"releasedAt" yaml:"releasedAt"`
}

// Reports will be uploaded as artifacts, and often displayed in the Gitlab UI, such as in Merge Requests.
// See: https://docs.gitlab.com/ee/ci/yaml/#artifactsreports
//
// Experimental.
type Reports struct {
	// Path for file(s) that should be parsed as Cobertura XML coverage report.
	// Experimental.
	Cobertura *[]*string `json:"cobertura" yaml:"cobertura"`
	// Path to file or list of files with code quality report(s) (such as Code Climate).
	// Experimental.
	Codequality *[]*string `json:"codequality" yaml:"codequality"`
	// Path to file or list of files with Container scanning vulnerabilities report(s).
	// Experimental.
	ContainerScanning *[]*string `json:"containerScanning" yaml:"containerScanning"`
	// Path to file or list of files with DAST vulnerabilities report(s).
	// Experimental.
	Dast *[]*string `json:"dast" yaml:"dast"`
	// Path to file or list of files with Dependency scanning vulnerabilities report(s).
	// Experimental.
	DependencyScanning *[]*string `json:"dependencyScanning" yaml:"dependencyScanning"`
	// Path to file or list of files containing runtime-created variables for this job.
	// Experimental.
	Dotenv *[]*string `json:"dotenv" yaml:"dotenv"`
	// Path for file(s) that should be parsed as JUnit XML result.
	// Experimental.
	Junit *[]*string `json:"junit" yaml:"junit"`
	// Deprecated in 12.8: Path to file or list of files with license report(s).
	// Experimental.
	LicenseManagement *[]*string `json:"licenseManagement" yaml:"licenseManagement"`
	// Path to file or list of files with license report(s).
	// Experimental.
	LicenseScanning *[]*string `json:"licenseScanning" yaml:"licenseScanning"`
	// Path to file or list of files containing code intelligence (Language Server Index Format).
	// Experimental.
	Lsif *[]*string `json:"lsif" yaml:"lsif"`
	// Path to file or list of files with custom metrics report(s).
	// Experimental.
	Metrics *[]*string `json:"metrics" yaml:"metrics"`
	// Path to file or list of files with performance metrics report(s).
	// Experimental.
	Performance *[]*string `json:"performance" yaml:"performance"`
	// Path to file or list of files with requirements report(s).
	// Experimental.
	Requirements *[]*string `json:"requirements" yaml:"requirements"`
	// Path to file or list of files with SAST vulnerabilities report(s).
	// Experimental.
	Sast *[]*string `json:"sast" yaml:"sast"`
	// Path to file or list of files with secret detection report(s).
	// Experimental.
	SecretDetection *[]*string `json:"secretDetection" yaml:"secretDetection"`
	// Path to file or list of files with terraform plan(s).
	// Experimental.
	Terraform *[]*string `json:"terraform" yaml:"terraform"`
}

// How many times a job is retried if it fails.
//
// If not defined, defaults to 0 and jobs do not retry.
// See: https://docs.gitlab.com/ee/ci/yaml/#retry
//
// Experimental.
type Retry struct {
	// 0 (default), 1, or 2.
	// Experimental.
	Max *float64 `json:"max" yaml:"max"`
	// Either a single or array of error types to trigger job retry.
	// Experimental.
	When interface{} `json:"when" yaml:"when"`
}

// A CI/CD secret.
// Experimental.
type Secret struct {
	// Experimental.
	Vault *VaultConfig `json:"vault" yaml:"vault"`
}

// Used to specify an additional Docker image to run scripts in.
//
// The service image is linked to the image specified in the @Default image keyword.
// See: https://docs.gitlab.com/ee/ci/yaml/#services
//
// Experimental.
type Service struct {
	// Full name of the image that should be used.
	//
	// It should contain the Registry part if needed.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Additional alias that can be used to access the service from the job's container.
	//
	// Read Accessing the services for more information.
	// Experimental.
	Alias *string `json:"alias" yaml:"alias"`
	// Command or script that should be used as the container's command.
	//
	// It will be translated to arguments passed to Docker after the image's name. The syntax is similar to Dockerfile's CMD directive, where each shell token is a separate string in the array.
	// Experimental.
	Command *[]*string `json:"command" yaml:"command"`
	// Command or script that should be executed as the container's entrypoint.
	//
	// It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array.
	// Experimental.
	Entrypoint *[]*string `json:"entrypoint" yaml:"entrypoint"`
}

// You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerstrategy
//
// Experimental.
type Strategy string

const (
	// Experimental.
	Strategy_DEPEND Strategy = "DEPEND"
)

// Trigger a multi-project or a child pipeline.
//
// Read more:.
// See: https://docs.gitlab.com/ee/ci/yaml/README.html#trigger-syntax-for-child-pipeline
//
// Experimental.
type Trigger struct {
	// The branch name that a downstream pipeline will use.
	// Experimental.
	Branch *string `json:"branch" yaml:"branch"`
	// A list of local files or artifacts from other jobs to define the pipeline.
	// Experimental.
	Include *[]*TriggerInclude `json:"include" yaml:"include"`
	// Path to the project, e.g. `group/project`, or `group/sub-group/project`.
	// Experimental.
	Project *string `json:"project" yaml:"project"`
	// You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend.
	// Experimental.
	Strategy Strategy `json:"strategy" yaml:"strategy"`
}

// References a local file or an artifact from another job to define the pipeline configuration.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerinclude
//
// Experimental.
type TriggerInclude struct {
	// Relative path to the generated YAML file which is extracted from the artifacts and used as the configuration for triggering the child pipeline.
	// Experimental.
	Artifact *string `json:"artifact" yaml:"artifact"`
	// Relative path from repository root (`/`) to the pipeline configuration YAML file.
	// Experimental.
	File *string `json:"file" yaml:"file"`
	// Job name which generates the artifact.
	// Experimental.
	Job *string `json:"job" yaml:"job"`
	// Relative path from local repository root (`/`) to the local YAML file to define the pipeline configuration.
	// Experimental.
	Local *string `json:"local" yaml:"local"`
	// Path to another private project under the same GitLab instance, like `group/project` or `group/sub-group/project`.
	// Experimental.
	Project *string `json:"project" yaml:"project"`
	// Branch/Tag/Commit hash for the target project.
	// Experimental.
	Ref *string `json:"ref" yaml:"ref"`
	// Name of the template YAML file to use in the pipeline configuration.
	// Experimental.
	Template *string `json:"template" yaml:"template"`
}

// Explains what the global variable is used for, what the acceptable values are.
// See: https://docs.gitlab.com/ee/ci/yaml/#variables
//
// Experimental.
type VariableConfig struct {
	// Define a global variable that is prefilled when running a pipeline manually.
	//
	// Must be used with value.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// The variable value.
	// Experimental.
	Value *string `json:"value" yaml:"value"`
}

// Specification for a secret provided by a HashiCorp Vault.
// See: https://www.vaultproject.io/
//
// Experimental.
type VaultConfig struct {
	// Experimental.
	Engine *Engine `json:"engine" yaml:"engine"`
	// Experimental.
	Field *string `json:"field" yaml:"field"`
	// Path to the secret.
	// Experimental.
	Path *string `json:"path" yaml:"path"`
}

// Used to control pipeline behavior.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflow
//
// Experimental.
type Workflow struct {
	// Used to control whether or not a whole pipeline is created.
	// Experimental.
	Rules *[]*WorkflowRule `json:"rules" yaml:"rules"`
}

// Used to control whether or not a whole pipeline is created.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflowrules
//
// Experimental.
type WorkflowRule struct {
	// Experimental.
	Changes *[]*string `json:"changes" yaml:"changes"`
	// Experimental.
	Exists *[]*string `json:"exists" yaml:"exists"`
	// Experimental.
	If *string `json:"if" yaml:"if"`
	// Experimental.
	Variables *map[string]interface{} `json:"variables" yaml:"variables"`
	// Experimental.
	When JobWhen `json:"when" yaml:"when"`
}

// Describes the conditions for when to run the job.
//
// Defaults to 'on_success'.
// The value can only be 'always' or 'never' when used with workflow.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflowrules
//
// Experimental.
type WorkflowWhen string

const (
	// Experimental.
	WorkflowWhen_ALWAYS WorkflowWhen = "ALWAYS"
	// Experimental.
	WorkflowWhen_NEVER WorkflowWhen = "NEVER"
)

