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
	Action_PREPARE Action = "PREPARE"
	Action_START Action = "START"
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
	ExitCodes interface{} `json:"exitCodes"`
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
	Exclude *[]*string `json:"exclude"`
	// How long artifacts should be kept.
	//
	// They are saved 30 days by default. Artifacts that have expired are removed periodically via cron job. Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'.
	// Experimental.
	ExpireIn *string `json:"expireIn"`
	// Can be used to expose job artifacts in the merge request UI.
	//
	// GitLab will add a link <expose_as> to the relevant merge request that points to the artifact.
	// Experimental.
	ExposeAs *string `json:"exposeAs"`
	// Name for the archive created on job success.
	//
	// Can use variables in the name, e.g. '$CI_JOB_NAME'
	// Experimental.
	Name *string `json:"name"`
	// A list of paths to files/folders that should be included in the artifact.
	// Experimental.
	Paths *[]*string `json:"paths"`
	// Reports will be uploaded as artifacts, and often displayed in the Gitlab UI, such as in Merge Requests.
	// Experimental.
	Reports *Reports `json:"reports"`
	// Whether to add all untracked files (along with 'artifacts.paths') to the artifact.
	// Experimental.
	Untracked *bool `json:"untracked"`
	// Configure when artifacts are uploaded depended on job status.
	// Experimental.
	When CacheWhen `json:"when"`
}

// Asset configuration for a release.
// Experimental.
type Assets struct {
	// Include asset links in the release.
	// Experimental.
	Links *[]*Link `json:"links"`
}

// Cache Definition.
// See: https://docs.gitlab.com/ee/ci/yaml/#cache
//
// Experimental.
type Cache struct {
	// Defines when to save the cache, based on the status of the job (Default: Job Success).
	// Experimental.
	When CacheWhen `json:"when"`
}

// Configure when artifacts are uploaded depended on job status.
// See: https://docs.gitlab.com/ee/ci/yaml/#cachewhen
//
// Experimental.
type CacheWhen string

const (
	CacheWhen_ALWAYS CacheWhen = "ALWAYS"
	CacheWhen_ON_FAILURE CacheWhen = "ON_FAILURE"
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
	DefaultAfterScript() *[]*string
	DefaultArtifacts() *Artifacts
	DefaultBeforeScript() *[]*string
	DefaultCache() *Cache
	DefaultImage() *Image
	DefaultInterruptible() *bool
	DefaultRetry() *Retry
	DefaultTags() *[]*string
	DefaultTimeout() *string
	File() projen.YamlFile
	Jobs() *map[string]*Job
	Name() *string
	Pages() *Job
	Path() *string
	Project() projen.Project
	Stages() *[]*string
	Variables() *map[string]interface{}
	Workflow() *Workflow
	AddGlobalVariables(variables *map[string]interface{})
	AddIncludes(includes ...*Include)
	AddJobs(jobs *map[string]*Job)
	AddServices(services ...*Service)
	AddStages(stages ...*string)
	PostSynthesize()
	PreSynthesize()
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

// Add a globally defined variable to the CI configuration.
// Experimental.
func (c *jsiiProxy_CiConfiguration) AddGlobalVariables(variables *map[string]interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addGlobalVariables",
		[]interface{}{variables},
	)
}

// Add additional yml/yaml files to the CI includes.
// Experimental.
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

// Add jobs and their stages to the CI configuration.
// Experimental.
func (c *jsiiProxy_CiConfiguration) AddJobs(jobs *map[string]*Job) {
	_jsii_.InvokeVoid(
		c,
		"addJobs",
		[]interface{}{jobs},
	)
}

// Add additional services.
// Experimental.
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

// Add stages to the CI configuration if not already present.
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (c *jsiiProxy_CiConfiguration) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (c *jsiiProxy_CiConfiguration) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Default *Default `json:"default"`
	// An initial set of jobs to add to the configuration.
	// Experimental.
	Jobs *map[string]*Job `json:"jobs"`
	// A special job used to upload static sites to Gitlab pages.
	//
	// Requires a `public/` directory
	// with `artifacts.path` pointing to it.
	// Experimental.
	Pages *Job `json:"pages"`
	// Groups jobs into stages.
	//
	// All jobs in one stage must complete before next stage is
	// executed. If no stages are specified. Defaults to ['build', 'test', 'deploy'].
	// Experimental.
	Stages *[]*string `json:"stages"`
	// Global variables that are passed to jobs.
	//
	// If the job already has that variable defined, the job-level variable takes precedence.
	// Experimental.
	Variables *map[string]interface{} `json:"variables"`
	// Used to control pipeline behavior.
	// Experimental.
	Workflow *Workflow `json:"workflow"`
}

// Default settings for the CI Configuration.
//
// Jobs that do not define one or more of the listed keywords use the value defined in the default section.
// See: https://docs.gitlab.com/ee/ci/yaml/#default
//
// Experimental.
type Default struct {
	// Experimental.
	AfterScript *[]*string `json:"afterScript"`
	// Experimental.
	Artifacts *Artifacts `json:"artifacts"`
	// Experimental.
	BeforeScript *[]*string `json:"beforeScript"`
	// Experimental.
	Cache *Cache `json:"cache"`
	// Experimental.
	Image *Image `json:"image"`
	// Experimental.
	Interruptible *bool `json:"interruptible"`
	// Experimental.
	Retry *Retry `json:"retry"`
	// Experimental.
	Services *[]*Service `json:"services"`
	// Experimental.
	Tags *[]*string `json:"tags"`
	// Experimental.
	Timeout *string `json:"timeout"`
}

// Experimental.
type DefaultElement string

const (
	DefaultElement_AFTER_SCRIPT DefaultElement = "AFTER_SCRIPT"
	DefaultElement_ARTIFACTS DefaultElement = "ARTIFACTS"
	DefaultElement_BEFORE_SCRIPT DefaultElement = "BEFORE_SCRIPT"
	DefaultElement_CACHE DefaultElement = "CACHE"
	DefaultElement_IMAGE DefaultElement = "IMAGE"
	DefaultElement_INTERRUPTIBLE DefaultElement = "INTERRUPTIBLE"
	DefaultElement_RETRY DefaultElement = "RETRY"
	DefaultElement_SERVICES DefaultElement = "SERVICES"
	DefaultElement_TAGS DefaultElement = "TAGS"
	DefaultElement_TIMEOUT DefaultElement = "TIMEOUT"
)

// Explicitly specifies the tier of the deployment environment if non-standard environment name is used.
// Experimental.
type DeploymentTier string

const (
	DeploymentTier_DEVELOPMENT DeploymentTier = "DEVELOPMENT"
	DeploymentTier_OTHER DeploymentTier = "OTHER"
	DeploymentTier_PRODUCTION DeploymentTier = "PRODUCTION"
	DeploymentTier_STAGING DeploymentTier = "STAGING"
	DeploymentTier_TESTING DeploymentTier = "TESTING"
)

// The engine configuration for a secret.
// Experimental.
type Engine struct {
	// Name of the secrets engine.
	// Experimental.
	Name *string `json:"name"`
	// Path to the secrets engine.
	// Experimental.
	Path *string `json:"path"`
}

// The environment that a job deploys to.
// Experimental.
type Environment struct {
	// The name of the environment, e.g. 'qa', 'staging', 'production'.
	// Experimental.
	Name *string `json:"name"`
	// Specifies what this job will do.
	//
	// 'start' (default) indicates the job will start the deployment. 'prepare' indicates this will not affect the deployment. 'stop' indicates this will stop the deployment.
	// Experimental.
	Action Action `json:"action"`
	// The amount of time it should take before Gitlab will automatically stop the environment.
	//
	// Supports a wide variety of formats, e.g. '1 week', '3 mins 4 sec', '2 hrs 20 min', '2h20min', '6 mos 1 day', '47 yrs 6 mos and 4d', '3 weeks and 2 days'.
	// Experimental.
	AutoStopIn *string `json:"autoStopIn"`
	// Explicitly specifies the tier of the deployment environment if non-standard environment name is used.
	// Experimental.
	DeploymentTier DeploymentTier `json:"deploymentTier"`
	// Used to configure the kubernetes deployment for this environment.
	//
	// This is currently not supported for kubernetes clusters that are managed by Gitlab.
	// Experimental.
	Kubernetes *KubernetesConfig `json:"kubernetes"`
	// The name of a job to execute when the environment is about to be stopped.
	// Experimental.
	OnStop *string `json:"onStop"`
	// When set, this will expose buttons in various places for the current environment in Gitlab, that will take you to the defined URL.
	// Experimental.
	Url *string `json:"url"`
}

// Filtering options for when a job will run.
// Experimental.
type Filter struct {
	// Filter job creation based on files that were modified in a git push.
	// Experimental.
	Changes *[]*string `json:"changes"`
	// Filter job based on if Kubernetes integration is active.
	// Experimental.
	Kubernetes KubernetesEnum `json:"kubernetes"`
	// Control when to add jobs to a pipeline based on branch names or pipeline types.
	// Experimental.
	Refs *[]*string `json:"refs"`
	// Filter job by checking comparing values of environment variables.
	//
	// Read more about variable expressions: https://docs.gitlab.com/ee/ci/variables/README.html#variables-expressions
	// Experimental.
	Variables *[]*string `json:"variables"`
}

// A GitLab CI for the main `.gitlab-ci.yml` file.
// Experimental.
type GitlabConfiguration interface {
	CiConfiguration
	DefaultAfterScript() *[]*string
	DefaultArtifacts() *Artifacts
	DefaultBeforeScript() *[]*string
	DefaultCache() *Cache
	DefaultImage() *Image
	DefaultInterruptible() *bool
	DefaultRetry() *Retry
	DefaultTags() *[]*string
	DefaultTimeout() *string
	File() projen.YamlFile
	Jobs() *map[string]*Job
	Name() *string
	NestedTemplates() *map[string]NestedConfiguration
	Pages() *Job
	Path() *string
	Project() projen.Project
	Stages() *[]*string
	Variables() *map[string]interface{}
	Workflow() *Workflow
	AddGlobalVariables(variables *map[string]interface{})
	AddIncludes(includes ...*Include)
	AddJobs(jobs *map[string]*Job)
	AddServices(services ...*Service)
	AddStages(stages ...*string)
	CreateNestedTemplates(config *map[string]*CiConfigurationOptions)
	PostSynthesize()
	PreSynthesize()
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

// Add a globally defined variable to the CI configuration.
// Experimental.
func (g *jsiiProxy_GitlabConfiguration) AddGlobalVariables(variables *map[string]interface{}) {
	_jsii_.InvokeVoid(
		g,
		"addGlobalVariables",
		[]interface{}{variables},
	)
}

// Add additional yml/yaml files to the CI includes.
// Experimental.
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

// Add jobs and their stages to the CI configuration.
// Experimental.
func (g *jsiiProxy_GitlabConfiguration) AddJobs(jobs *map[string]*Job) {
	_jsii_.InvokeVoid(
		g,
		"addJobs",
		[]interface{}{jobs},
	)
}

// Add additional services.
// Experimental.
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

// Add stages to the CI configuration if not already present.
// Experimental.
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

// Creates and adds nested templates to the includes of the main CI.
//
// Additionally adds their stages to the main CI if they are not already present.
// You can futher customize nested templates through the `nestedTemplates` property.
// E.g. gitlabConfig.nestedTemplates['templateName']?.addStages('stageName')
// Experimental.
func (g *jsiiProxy_GitlabConfiguration) CreateNestedTemplates(config *map[string]*CiConfigurationOptions) {
	_jsii_.InvokeVoid(
		g,
		"createNestedTemplates",
		[]interface{}{config},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (g *jsiiProxy_GitlabConfiguration) PostSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (g *jsiiProxy_GitlabConfiguration) PreSynthesize() {
	_jsii_.InvokeVoid(
		g,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Name *string `json:"name"`
	// Command or script that should be executed as the container's entrypoint.
	//
	// It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array.
	// Experimental.
	Entrypoint *[]interface{} `json:"entrypoint"`
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
	File *[]*string `json:"file"`
	// Relative path from local repository root (`/`) to the `yaml`/`yml` file template.
	//
	// The file must be on the same branch, and does not work across git submodules.
	// Experimental.
	Local *string `json:"local"`
	// Path to the project, e.g. `group/project`, or `group/sub-group/project`.
	// Experimental.
	Project *string `json:"project"`
	// Branch/Tag/Commit-hash for the target project.
	// Experimental.
	Ref *string `json:"ref"`
	// URL to a `yaml`/`yml` template file using HTTP/HTTPS.
	// Experimental.
	Remote *string `json:"remote"`
	// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
	// Experimental.
	Rules *[]*IncludeRule `json:"rules"`
	// Use a `.gitlab-ci.yml` template as a base, e.g. `Nodejs.gitlab-ci.yml`.
	// Experimental.
	Template *string `json:"template"`
}

// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
// See: https://docs.gitlab.com/ee/ci/yaml/includes.html#use-rules-with-include
//
// Experimental.
type IncludeRule struct {
	// Experimental.
	AllowFailure interface{} `json:"allowFailure"`
	// Experimental.
	Changes *[]*string `json:"changes"`
	// Experimental.
	Exists *[]*string `json:"exists"`
	// Experimental.
	If *string `json:"if"`
	// Experimental.
	StartIn *string `json:"startIn"`
	// Experimental.
	Variables *map[string]interface{} `json:"variables"`
	// Experimental.
	When JobWhen `json:"when"`
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
	// Or subset of inherited defaults
	// Experimental.
	Default interface{} `json:"default"`
	// Whether to inherit all globally-defined variables or not.
	//
	// Or subset of inherited variables
	// Experimental.
	Variables interface{} `json:"variables"`
}

// Jobs are the most fundamental element of a .gitlab-ci.yml file.
// See: https://docs.gitlab.com/ee/ci/jobs/
//
// Experimental.
type Job struct {
	// Experimental.
	AfterScript *[]*string `json:"afterScript"`
	// Whether to allow the pipeline to continue running on job failure (Default: false).
	// Experimental.
	AllowFailure interface{} `json:"allowFailure"`
	// Experimental.
	Artifacts *Artifacts `json:"artifacts"`
	// Experimental.
	BeforeScript *[]*string `json:"beforeScript"`
	// Experimental.
	Cache *Cache `json:"cache"`
	// Must be a regular expression, optionally but recommended to be quoted, and must be surrounded with '/'.
	//
	// Example: '/Code coverage: \d+\.\d+/'
	// Experimental.
	Coverage *string `json:"coverage"`
	// Specify a list of job names from earlier stages from which artifacts should be loaded.
	//
	// By default, all previous artifacts are passed. Use an empty array to skip downloading artifacts.
	// Experimental.
	Dependencies *[]*string `json:"dependencies"`
	// Used to associate environment metadata with a deploy.
	//
	// Environment can have a name and URL attached to it, and will be displayed under /environments under the project.
	// Experimental.
	Environment interface{} `json:"environment"`
	// Job will run *except* for when these filtering options match.
	// Experimental.
	Except interface{} `json:"except"`
	// The name of one or more jobs to inherit configuration from.
	// Experimental.
	Extends *[]*string `json:"extends"`
	// Experimental.
	Image *Image `json:"image"`
	// Controls inheritance of globally-defined defaults and variables.
	//
	// Boolean values control inheritance of all default: or variables: keywords. To inherit only a subset of default: or variables: keywords, specify what you wish to inherit. Anything not listed is not inherited.
	// Experimental.
	Inherit *Inherit `json:"inherit"`
	// Experimental.
	Interruptible *bool `json:"interruptible"`
	// The list of jobs in previous stages whose sole completion is needed to start the current job.
	// Experimental.
	Needs *[]interface{} `json:"needs"`
	// Job will run *only* when these filtering options match.
	// Experimental.
	Only interface{} `json:"only"`
	// Parallel will split up a single job into several, and provide `CI_NODE_INDEX` and `CI_NODE_TOTAL` environment variables for the running jobs.
	// Experimental.
	Parallel interface{} `json:"parallel"`
	// Indicates that the job creates a Release.
	// Experimental.
	Release *Release `json:"release"`
	// Limit job concurrency.
	//
	// Can be used to ensure that the Runner will not run certain jobs simultaneously.
	// Experimental.
	ResourceGroup *string `json:"resourceGroup"`
	// Experimental.
	Retry *Retry `json:"retry"`
	// Rules allows for an array of individual rule objects to be evaluated in order, until one matches and dynamically provides attributes to the job.
	// Experimental.
	Rules *[]*IncludeRule `json:"rules"`
	// Shell scripts executed by the Runner.
	//
	// The only required property of jobs. Be careful with special characters (e.g. `:`, `{`, `}`, `&`) and use single or double quotes to avoid issues.
	// Experimental.
	Script *[]*string `json:"script"`
	// CI/CD secrets.
	// Experimental.
	Secrets *map[string]*map[string]*Secret `json:"secrets"`
	// Experimental.
	Services *[]*Service `json:"services"`
	// Define what stage the job will run in.
	// Experimental.
	Stage *string `json:"stage"`
	// Experimental.
	StartIn *string `json:"startIn"`
	// Experimental.
	Tags *[]*string `json:"tags"`
	// Experimental.
	Timeout *string `json:"timeout"`
	// Trigger allows you to define downstream pipeline trigger.
	//
	// When a job created from trigger definition is started by GitLab, a downstream pipeline gets created. Read more: https://docs.gitlab.com/ee/ci/yaml/README.html#trigger
	// Experimental.
	Trigger interface{} `json:"trigger"`
	// Configurable values that are passed to the Job.
	// Experimental.
	Variables *map[string]interface{} `json:"variables"`
	// Describes the conditions for when to run the job.
	//
	// Defaults to 'on_success'.
	// Experimental.
	When JobWhen `json:"when"`
}

// Describes the conditions for when to run the job.
//
// Defaults to 'on_success'.
// See: https://docs.gitlab.com/ee/ci/yaml/#when
//
// Experimental.
type JobWhen string

const (
	JobWhen_ALWAYS JobWhen = "ALWAYS"
	JobWhen_DELAYED JobWhen = "DELAYED"
	JobWhen_MANUAL JobWhen = "MANUAL"
	JobWhen_NEVER JobWhen = "NEVER"
	JobWhen_ON_FAILURE JobWhen = "ON_FAILURE"
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
	Namespace *string `json:"namespace"`
}

// Filter job based on if Kubernetes integration is active.
// Experimental.
type KubernetesEnum string

const (
	KubernetesEnum_ACTIVE KubernetesEnum = "ACTIVE"
)

// Link configuration for an asset.
// Experimental.
type Link struct {
	// The name of the link.
	// Experimental.
	Name *string `json:"name"`
	// The URL to download a file.
	// Experimental.
	Url *string `json:"url"`
	// The redirect link to the url.
	// Experimental.
	Filepath *string `json:"filepath"`
	// The content kind of what users can download via url.
	// Experimental.
	LinkType LinkType `json:"linkType"`
}

// The content kind of what users can download via url.
// Experimental.
type LinkType string

const (
	LinkType_IMAGE LinkType = "IMAGE"
	LinkType_OTHER LinkType = "OTHER"
	LinkType_PACKAGE LinkType = "PACKAGE"
	LinkType_RUNBOOK LinkType = "RUNBOOK"
)

// A jobs in a previous stage whose sole completion is needed to start the current job.
// Experimental.
type Need struct {
	// Experimental.
	Job *string `json:"job"`
	// Experimental.
	Artifacts *bool `json:"artifacts"`
	// Experimental.
	Optional *bool `json:"optional"`
	// Experimental.
	Pipeline *string `json:"pipeline"`
	// Experimental.
	Project *string `json:"project"`
	// Experimental.
	Ref *string `json:"ref"`
}

// A GitLab CI for templates that are created and included in the `.gitlab-ci.yml` file.
// Experimental.
type NestedConfiguration interface {
	CiConfiguration
	DefaultAfterScript() *[]*string
	DefaultArtifacts() *Artifacts
	DefaultBeforeScript() *[]*string
	DefaultCache() *Cache
	DefaultImage() *Image
	DefaultInterruptible() *bool
	DefaultRetry() *Retry
	DefaultTags() *[]*string
	DefaultTimeout() *string
	File() projen.YamlFile
	Jobs() *map[string]*Job
	Name() *string
	Pages() *Job
	Parent() GitlabConfiguration
	Path() *string
	Project() projen.Project
	Stages() *[]*string
	Variables() *map[string]interface{}
	Workflow() *Workflow
	AddGlobalVariables(variables *map[string]interface{})
	AddIncludes(includes ...*Include)
	AddJobs(jobs *map[string]*Job)
	AddServices(services ...*Service)
	AddStages(stages ...*string)
	PostSynthesize()
	PreSynthesize()
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

// Add a globally defined variable to the CI configuration.
// Experimental.
func (n *jsiiProxy_NestedConfiguration) AddGlobalVariables(variables *map[string]interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addGlobalVariables",
		[]interface{}{variables},
	)
}

// Add additional yml/yaml files to the CI includes.
// Experimental.
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

// Add jobs and their stages to the CI configuration.
// Experimental.
func (n *jsiiProxy_NestedConfiguration) AddJobs(jobs *map[string]*Job) {
	_jsii_.InvokeVoid(
		n,
		"addJobs",
		[]interface{}{jobs},
	)
}

// Add additional services.
// Experimental.
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

// Add stages to the Nested configuration and the main CI file if not already present.
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (n *jsiiProxy_NestedConfiguration) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (n *jsiiProxy_NestedConfiguration) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Matrix *[]*map[string]*[]interface{} `json:"matrix"`
}

// Indicates that the job creates a Release.
// Experimental.
type Release struct {
	// Specifies the longer description of the Release.
	// Experimental.
	Description *string `json:"description"`
	// The tag_name must be specified.
	//
	// It can refer to an existing Git tag or can be specified by the user.
	// Experimental.
	TagName *string `json:"tagName"`
	// Experimental.
	Assets *Assets `json:"assets"`
	// The title of each milestone the release is associated with.
	// Experimental.
	Milestones *[]*string `json:"milestones"`
	// The Release name.
	//
	// If omitted, it is populated with the value of release: tag_name.
	// Experimental.
	Name *string `json:"name"`
	// If the release: tag_name doesn’t exist yet, the release is created from ref.
	//
	// ref can be a commit SHA, another tag name, or a branch name.
	// Experimental.
	Ref *string `json:"ref"`
	// The date and time when the release is ready.
	//
	// Defaults to the current date and time if not defined. Should be enclosed in quotes and expressed in ISO 8601 format.
	// Experimental.
	ReleasedAt *string `json:"releasedAt"`
}

// Reports will be uploaded as artifacts, and often displayed in the Gitlab UI, such as in Merge Requests.
// See: https://docs.gitlab.com/ee/ci/yaml/#artifactsreports
//
// Experimental.
type Reports struct {
	// Path for file(s) that should be parsed as Cobertura XML coverage report.
	// Experimental.
	Cobertura *[]*string `json:"cobertura"`
	// Path to file or list of files with code quality report(s) (such as Code Climate).
	// Experimental.
	Codequality *[]*string `json:"codequality"`
	// Path to file or list of files with Container scanning vulnerabilities report(s).
	// Experimental.
	ContainerScanning *[]*string `json:"containerScanning"`
	// Path to file or list of files with DAST vulnerabilities report(s).
	// Experimental.
	Dast *[]*string `json:"dast"`
	// Path to file or list of files with Dependency scanning vulnerabilities report(s).
	// Experimental.
	DependencyScanning *[]*string `json:"dependencyScanning"`
	// Path to file or list of files containing runtime-created variables for this job.
	// Experimental.
	Dotenv *[]*string `json:"dotenv"`
	// Path for file(s) that should be parsed as JUnit XML result.
	// Experimental.
	Junit *[]*string `json:"junit"`
	// Deprecated in 12.8: Path to file or list of files with license report(s).
	// Experimental.
	LicenseManagement *[]*string `json:"licenseManagement"`
	// Path to file or list of files with license report(s).
	// Experimental.
	LicenseScanning *[]*string `json:"licenseScanning"`
	// Path to file or list of files containing code intelligence (Language Server Index Format).
	// Experimental.
	Lsif *[]*string `json:"lsif"`
	// Path to file or list of files with custom metrics report(s).
	// Experimental.
	Metrics *[]*string `json:"metrics"`
	// Path to file or list of files with performance metrics report(s).
	// Experimental.
	Performance *[]*string `json:"performance"`
	// Path to file or list of files with requirements report(s).
	// Experimental.
	Requirements *[]*string `json:"requirements"`
	// Path to file or list of files with SAST vulnerabilities report(s).
	// Experimental.
	Sast *[]*string `json:"sast"`
	// Path to file or list of files with secret detection report(s).
	// Experimental.
	SecretDetection *[]*string `json:"secretDetection"`
	// Path to file or list of files with terraform plan(s).
	// Experimental.
	Terraform *[]*string `json:"terraform"`
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
	Max *float64 `json:"max"`
	// Either a single or array of error types to trigger job retry.
	// Experimental.
	When interface{} `json:"when"`
}

// A CI/CD secret.
// Experimental.
type Secret struct {
	// Experimental.
	Vault *VaultConfig `json:"vault"`
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
	Name *string `json:"name"`
	// Additional alias that can be used to access the service from the job's container.
	//
	// Read Accessing the services for more information.
	// Experimental.
	Alias *string `json:"alias"`
	// Command or script that should be used as the container's command.
	//
	// It will be translated to arguments passed to Docker after the image's name. The syntax is similar to Dockerfile's CMD directive, where each shell token is a separate string in the array.
	// Experimental.
	Command *[]*string `json:"command"`
	// Command or script that should be executed as the container's entrypoint.
	//
	// It will be translated to Docker's --entrypoint option while creating the container. The syntax is similar to Dockerfile's ENTRYPOINT directive, where each shell token is a separate string in the array.
	// Experimental.
	Entrypoint *[]*string `json:"entrypoint"`
}

// You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerstrategy
//
// Experimental.
type Strategy string

const (
	Strategy_DEPEND Strategy = "DEPEND"
)

// Trigger a multi-project or a child pipeline.
//
// Read more:
// See: https://docs.gitlab.com/ee/ci/yaml/README.html#trigger-syntax-for-child-pipeline
//
// Experimental.
type Trigger struct {
	// The branch name that a downstream pipeline will use.
	// Experimental.
	Branch *string `json:"branch"`
	// A list of local files or artifacts from other jobs to define the pipeline.
	// Experimental.
	Include *[]*TriggerInclude `json:"include"`
	// Path to the project, e.g. `group/project`, or `group/sub-group/project`.
	// Experimental.
	Project *string `json:"project"`
	// You can mirror the pipeline status from the triggered pipeline to the source bridge job by using strategy: depend.
	// Experimental.
	Strategy Strategy `json:"strategy"`
}

// References a local file or an artifact from another job to define the pipeline configuration.
// See: https://docs.gitlab.com/ee/ci/yaml/#triggerinclude
//
// Experimental.
type TriggerInclude struct {
	// Relative path to the generated YAML file which is extracted from the artifacts and used as the configuration for triggering the child pipeline.
	// Experimental.
	Artifact *string `json:"artifact"`
	// Relative path from repository root (`/`) to the pipeline configuration YAML file.
	// Experimental.
	File *string `json:"file"`
	// Job name which generates the artifact.
	// Experimental.
	Job *string `json:"job"`
	// Relative path from local repository root (`/`) to the local YAML file to define the pipeline configuration.
	// Experimental.
	Local *string `json:"local"`
	// Path to another private project under the same GitLab instance, like `group/project` or `group/sub-group/project`.
	// Experimental.
	Project *string `json:"project"`
	// Branch/Tag/Commit hash for the target project.
	// Experimental.
	Ref *string `json:"ref"`
	// Name of the template YAML file to use in the pipeline configuration.
	// Experimental.
	Template *string `json:"template"`
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
	Description *string `json:"description"`
	// The variable value.
	// Experimental.
	Value *string `json:"value"`
}

// Specification for a secret provided by a HashiCorp Vault.
// See: https://www.vaultproject.io/
//
// Experimental.
type VaultConfig struct {
	// Experimental.
	Engine *Engine `json:"engine"`
	// Experimental.
	Field *string `json:"field"`
	// Path to the secret.
	// Experimental.
	Path *string `json:"path"`
}

// Used to control pipeline behavior.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflow
//
// Experimental.
type Workflow struct {
	// Used to control whether or not a whole pipeline is created.
	// Experimental.
	Rules *[]*WorkflowRule `json:"rules"`
}

// Used to control whether or not a whole pipeline is created.
// See: https://docs.gitlab.com/ee/ci/yaml/#workflowrules
//
// Experimental.
type WorkflowRule struct {
	// Experimental.
	Changes *[]*string `json:"changes"`
	// Experimental.
	Exists *[]*string `json:"exists"`
	// Experimental.
	If *string `json:"if"`
	// Experimental.
	Variables *map[string]interface{} `json:"variables"`
	// Experimental.
	When JobWhen `json:"when"`
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
	WorkflowWhen_ALWAYS WorkflowWhen = "ALWAYS"
	WorkflowWhen_NEVER WorkflowWhen = "NEVER"
)

