package circleci

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/circleci/internal"
)

// Options for class {@link Circleci}.
// See: https://circleci.com/docs/2.0/configuration-reference/
//
// Experimental.
type CircleCiProps struct {
	// List of Jobs to create unique steps per pipeline, e.g. ```json jobs: [{   identifier: "compile",   docker: { image: "golang:alpine" }   steps: ["checkout", run: {command: "go build ."}] }] ```.
	// See: https://circleci.com/docs/2.0/configuration-reference/#jobs
	//
	// Experimental.
	Jobs *[]*Job `field:"optional" json:"jobs" yaml:"jobs"`
	// Contains a map of CirclCi Orbs ```json orbs: {   node: "circleci/node@5.0.1"   slack: "circleci/slack@4.8.3" } ```.
	// Experimental.
	Orbs *map[string]*string `field:"optional" json:"orbs" yaml:"orbs"`
	// The setup field enables you to conditionally trigger configurations from outside the primary .circleci parent directory, update pipeline parameters, or generate customized configurations.
	// See: https://circleci.com/docs/2.0/configuration-reference/#setup
	//
	// Experimental.
	Setup *bool `field:"optional" json:"setup" yaml:"setup"`
	// pipeline version.
	// See: https://circleci.com/docs/2.0/configuration-reference/#version
	//
	// Experimental.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
	// List of Workflows of pipeline, e.g. ```json workflows: {    {      identifier: "build",        jobs: [{           identifier: "node/install",           context: ["npm"],        }]    } } ```.
	// See: https://circleci.com/docs/2.0/configuration-reference/#workflows
	//
	// Experimental.
	Workflows *[]*Workflow `field:"optional" json:"workflows" yaml:"workflows"`
}

// Circleci Class to manage `.circleci/config.yml`. Check projen's docs for more information.
// See: https://circleci.com/docs/2.0/configuration-reference/
//
// Experimental.
type Circleci interface {
	projen.Component
	// The yaml file for the Circleci pipeline.
	// Experimental.
	File() projen.YamlFile
	// Experimental.
	Project() projen.Project
	// Add a Circleci Orb to pipeline.
	//
	// Will throw error if the orb already exists.
	// Experimental.
	AddOrb(name *string, orb *string)
	// add new workflow to existing pipeline.
	// Experimental.
	AddWorkflow(workflow *Workflow)
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

// The jsii proxy struct for Circleci
type jsiiProxy_Circleci struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Circleci) File() projen.YamlFile {
	var returns projen.YamlFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Circleci) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCircleci(project projen.Project, options *CircleCiProps) Circleci {
	_init_.Initialize()

	j := jsiiProxy_Circleci{}

	_jsii_.Create(
		"projen.circleci.Circleci",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewCircleci_Override(c Circleci, project projen.Project, options *CircleCiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.circleci.Circleci",
		[]interface{}{project, options},
		c,
	)
}

func (c *jsiiProxy_Circleci) AddOrb(name *string, orb *string) {
	_jsii_.InvokeVoid(
		c,
		"addOrb",
		[]interface{}{name, orb},
	)
}

func (c *jsiiProxy_Circleci) AddWorkflow(workflow *Workflow) {
	_jsii_.InvokeVoid(
		c,
		"addWorkflow",
		[]interface{}{workflow},
	)
}

func (c *jsiiProxy_Circleci) PostSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"postSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Circleci) PreSynthesize() {
	_jsii_.InvokeVoid(
		c,
		"preSynthesize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Circleci) Synthesize() {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		nil, // no parameters
	)
}

// Options for docker executor.
// See: https://circleci.com/docs/2.0/configuration-reference/#docker
//
// Experimental.
type Docker struct {
	// The name of a custom docker image to use.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Authentication for registries using standard docker login credentials.
	// Experimental.
	Auth *map[string]*string `field:"optional" json:"auth" yaml:"auth"`
	// Authentication for AWS Elastic Container Registry (ECR).
	// Experimental.
	AwsAuth *map[string]*string `field:"optional" json:"awsAuth" yaml:"awsAuth"`
	// The command used as pid 1 (or args for entrypoint) when launching the container.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The command used as executable when launching the container.
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// A map of environment variable names and values.
	// Experimental.
	Environment *map[string]interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The name the container is reachable by.
	//
	// By default, container services are accessible through localhost.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Which user to run commands as within the Docker container.
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
}

// The branches key controls whether the current branch should have a schedule trigger created for it, where current branch is the branch containing the config.yml file with the trigger stanza. That is, a push on the main branch will only schedule a workflow for the main branch.
//
// Branches can have the keys only and ignore which either map to a single string naming a branch.
// You may also use regular expressions to match against branches by enclosing them with /’s, or map to a list of such strings.
// Regular expressions must match the entire string.
//
// Any branches that match only will run the job.
// Any branches that match ignore will not run the job.
// If neither only nor ignore are specified then all branches will run the job.
// If both only and ignore are specified the only is considered before ignore.
// See: https://circleci.com/docs/2.0/configuration-reference/#filters
//
// Experimental.
type Filter struct {
	// Experimental.
	Branches *FilterConfig `field:"optional" json:"branches" yaml:"branches"`
	// Experimental.
	Tags *FilterConfig `field:"optional" json:"tags" yaml:"tags"`
}

// set an inclusive or exclusive filter.
// See: https://circleci.com/docs/2.0/configuration-reference/#filters
//
// Experimental.
type FilterConfig struct {
	// Either a single branch specifier, or a list of branch specifiers.
	// Experimental.
	Ignore *[]*string `field:"optional" json:"ignore" yaml:"ignore"`
	// Either a single branch specifier, or a list of branch specifiers.
	// Experimental.
	Only *[]*string `field:"optional" json:"only" yaml:"only"`
}

// A Workflow is comprised of one or more uniquely named jobs.
//
// Jobs are specified in the jobs map,
// see Sample 2.0 config.yml for two examples of a job map.
// The name of the job is the key in the map, and the value is a map describing the job.
// Each job consists of the job’s name as a key and a map as a value. A name should be case insensitive unique within a current jobs list.
// See: https://circleci.com/docs/2.0/configuration-reference/#job_name
//
// Experimental.
type Job struct {
	// name of dynamic key *.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Experimental.
	Docker *[]*Docker `field:"optional" json:"docker" yaml:"docker"`
	// A map of environment variable names and values.
	// Experimental.
	Environment *map[string]interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Experimental.
	Machine *Machine `field:"optional" json:"machine" yaml:"machine"`
	// Experimental.
	Macos *Macos `field:"optional" json:"macos" yaml:"macos"`
	// Number of parallel instances of this job to run (default: 1).
	// Experimental.
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Parameters for making a job explicitly configurable in a workflow.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// {@link ResourceClass}.
	// Experimental.
	ResourceClass *string `field:"optional" json:"resourceClass" yaml:"resourceClass"`
	// Shell to use for execution command in all steps.
	//
	// Can be overridden by shell in each step.
	// Experimental.
	Shell *string `field:"optional" json:"shell" yaml:"shell"`
	// no type support here, for syntax {@see https://circleci.com/docs/2.0/configuration-reference/#steps}.
	// Experimental.
	Steps *[]interface{} `field:"optional" json:"steps" yaml:"steps"`
	// In which directory to run the steps.
	//
	// Will be interpreted as an absolute path. Default: `~/project`
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// A job may have a type of approval indicating it must be manually approved before downstream jobs may proceed.
// See: https://circleci.com/docs/2.0/configuration-reference/#type
//
// Experimental.
type JobType string

const (
	// Experimental.
	JobType_APPROVAL JobType = "APPROVAL"
)

// Specify when to enable or disable the step.
// See: https://circleci.com/docs/2.0/configuration-reference/#steps
//
// Experimental.
type JobWhen string

const (
	// Experimental.
	JobWhen_ALWAYS JobWhen = "ALWAYS"
	// Experimental.
	JobWhen_ON_SUCCESS JobWhen = "ON_SUCCESS"
	// Experimental.
	JobWhen_ON_FAIL JobWhen = "ON_FAIL"
)

// Experimental.
type Machine struct {
	// The VM image to use.
	// See: https://circleci.com/docs/2.0/configuration-reference/#available-machine-images
	//
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// enable docker layer caching.
	// See: https://circleci.com/docs/2.0/configuration-reference/#available-machine-images
	//
	// Experimental.
	DockerLayerCaching *string `field:"optional" json:"dockerLayerCaching" yaml:"dockerLayerCaching"`
}

// CircleCI supports running jobs on macOS, to allow you to build, test, and deploy apps for macOS, iOS, tvOS and watchOS.
//
// To run a job in a macOS virtual machine,
// you must add the macos key to the top-level configuration for the job and specify
// the version of Xcode you would like to use.
// See: https://circleci.com/docs/2.0/configuration-reference/#macos
//
// Experimental.
type Macos struct {
	// The version of Xcode that is installed on the virtual machine.
	// Experimental.
	Xcode *string `field:"required" json:"xcode" yaml:"xcode"`
}

// The matrix stanza allows you to run a parameterized job multiple times with different arguments.
// See: https://circleci.com/docs/2.0/configuration-reference/#matrix-requires-version-21
//
// Experimental.
type Matrix struct {
	// An alias for the matrix, usable from another job’s requires stanza.
	//
	// Defaults to the name of the job being executed.
	// Experimental.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A map of parameter names to every value the job should be called with.
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// The resource_class feature allows configuring CPU and RAM resources for each job.
//
// Different resource classes are available for different executors, as described in the tables below.
// See: https://circleci.com/docs/2.0/configuration-reference/#resourceclass
//
// Experimental.
type ResourceClass string

const (
	// Experimental.
	ResourceClass_SMALL ResourceClass = "SMALL"
	// Experimental.
	ResourceClass_MEDIUM ResourceClass = "MEDIUM"
	// Experimental.
	ResourceClass_MEDIUM_PLUS ResourceClass = "MEDIUM_PLUS"
	// Experimental.
	ResourceClass_LARGE_X ResourceClass = "LARGE_X"
	// Experimental.
	ResourceClass_LARGE_2X ResourceClass = "LARGE_2X"
	// Experimental.
	ResourceClass_LARGE_2X_PLUS ResourceClass = "LARGE_2X_PLUS"
)

// Used for invoking all command-line programs, taking either a map of configuration values, or, when called in its short-form, a string that will be used as both the command and name.
//
// Run commands are executed using non-login shells by default,
// so you must explicitly source any dotfiles as part of the command.
//
// Not used because type incompatible types in steps array.
// See: https://circleci.com/docs/2.0/configuration-reference/#run
//
// Experimental.
type Run struct {
	// Command to run via the shell.
	// Experimental.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Whether this step should run in the background (default: false).
	// Experimental.
	Background *string `field:"optional" json:"background" yaml:"background"`
	// Additional environmental variables, locally scoped to command.
	// Experimental.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Title of the step to be shown in the CircleCI UI (default: full command).
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Elapsed time the command can run without output such as “20m”, “1.25h”, “5s”. The default is 10 minutes.
	// Experimental.
	NoOutputTimeout *string `field:"optional" json:"noOutputTimeout" yaml:"noOutputTimeout"`
	// Shell to use for execution command.
	// Experimental.
	Shell *string `field:"optional" json:"shell" yaml:"shell"`
	// Specify when to enable or disable the step.
	// Experimental.
	When *string `field:"optional" json:"when" yaml:"when"`
	// In which directory to run this step.
	//
	// Will be interpreted relative to the working_directory of the job). (default: .)
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// A workflow may have a schedule indicating it runs at a certain time.
// See: https://circleci.com/docs/2.0/configuration-reference/#schedule
//
// Experimental.
type Schedule struct {
	// Experimental.
	Filters *Filter `field:"required" json:"filters" yaml:"filters"`
	// The cron key is defined using POSIX crontab syntax.
	// Experimental.
	Cron *string `field:"optional" json:"cron" yaml:"cron"`
}

// Execution steps for Job.
// See: https://circleci.com/docs/2.0/configuration-reference/#steps
//
// Experimental.
type StepRun struct {
	// Experimental.
	Run *Run `field:"optional" json:"run" yaml:"run"`
}

// Specifies which triggers will cause this workflow to be executed.
//
// Default behavior is to trigger the workflow when pushing to a branch.
// See: https://circleci.com/docs/2.0/configuration-reference/#triggers
//
// Experimental.
type Triggers struct {
	// Experimental.
	Schedule *Schedule `field:"optional" json:"schedule" yaml:"schedule"`
}

// Used for orchestrating all jobs.
//
// Each workflow consists of the workflow name as a key and a map as a value.
// A name should be unique within the current config.yml.
// The top-level keys for the Workflows configuration are version and jobs.
// See: https://circleci.com/docs/2.0/configuration-reference/#workflows
//
// Experimental.
type Workflow struct {
	// name of dynamic key *.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Experimental.
	Jobs *[]*WorkflowJob `field:"optional" json:"jobs" yaml:"jobs"`
	// Experimental.
	Triggers *[]*Triggers `field:"optional" json:"triggers" yaml:"triggers"`
	// when is too dynamic to be casted to interfaces.
	//
	// Check Docu as reference.
	// See: https://circleci.com/docs/2.0/configuration-reference/#logic-statement-examples
	//
	// Experimental.
	When interface{} `field:"optional" json:"when" yaml:"when"`
}

// A Job is part of Workflow.
//
// A Job can be created with {@link Job} or it can be provided by the orb.
// See: https://circleci.com/docs/2.0/configuration-reference/#jobs-in-workflow
//
// Experimental.
type WorkflowJob struct {
	// name of dynamic key *.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The name of the context(s).
	//
	// The initial default name is org-global. Each context name must be unique.
	// Experimental.
	Context *[]*string `field:"optional" json:"context" yaml:"context"`
	// Job Filters can have the key branches or tags.
	// Experimental.
	Filters *Filter `field:"optional" json:"filters" yaml:"filters"`
	// Experimental.
	Matrix *Matrix `field:"optional" json:"matrix" yaml:"matrix"`
	// A replacement for the job name.
	//
	// Useful when calling a job multiple times.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Parameters passed to job when referencing a job from orb.
	// Experimental.
	OrbParameters *map[string]interface{} `field:"optional" json:"orbParameters" yaml:"orbParameters"`
	// Parameters for making a job explicitly configurable in a workflow.
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A list of jobs that must succeed for the job to start.
	// Experimental.
	Requires *[]*string `field:"optional" json:"requires" yaml:"requires"`
	// A job may have a type of approval indicating it must be manually approved before downstream jobs may proceed.
	// Experimental.
	Type JobType `field:"optional" json:"type" yaml:"type"`
}

