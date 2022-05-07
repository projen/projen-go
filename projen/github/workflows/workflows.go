package workflows


// Branch Protection Rule options.
// Experimental.
type BranchProtectionRuleOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Check run options.
// Experimental.
type CheckRunOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Check suite options.
// Experimental.
type CheckSuiteOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Credentials to use to authenticate to Docker registries.
// Experimental.
type ContainerCredentials struct {
	// The password.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

// Options petaining to container environments.
// Experimental.
type ContainerOptions struct {
	// The Docker image to use as the container to run the action.
	//
	// The value can
	// be the Docker Hub image name or a registry name.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// f the image's container registry requires authentication to pull the image, you can use credentials to set a map of the username and password.
	//
	// The credentials are the same values that you would provide to the docker
	// login command.
	// Experimental.
	Credentials *ContainerCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Sets a map of environment variables in the container.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Additional Docker container resource options.
	// See: https://docs.docker.com/engine/reference/commandline/create/#options
	//
	// Experimental.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Sets an array of ports to expose on the container.
	// Experimental.
	Ports *[]*float64 `field:"optional" json:"ports" yaml:"ports"`
	// Sets an array of volumes for the container to use.
	//
	// You can use volumes to
	// share data between services or other steps in a job. You can specify
	// named Docker volumes, anonymous Docker volumes, or bind mounts on the
	// host.
	//
	// To specify a volume, you specify the source and destination path:
	// `<source>:<destinationPath>`.
	// Experimental.
	Volumes *[]*string `field:"optional" json:"volumes" yaml:"volumes"`
}

// The Create event accepts no options.
// Experimental.
type CreateOptions struct {
}

// CRON schedule options.
// Experimental.
type CronScheduleOptions struct {
	// See: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/crontab.html#tag_20_25_07
	//
	// Experimental.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
}

// The Delete event accepts no options.
// Experimental.
type DeleteOptions struct {
}

// The Deployment event accepts no options.
// Experimental.
type DeploymentOptions struct {
}

// The Deployment status event accepts no options.
// Experimental.
type DeploymentStatusOptions struct {
}

// Discussion comment options.
// Experimental.
type DiscussionCommentOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Discussion options.
// Experimental.
type DiscussionOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// The Fork event accepts no options.
// Experimental.
type ForkOptions struct {
}

// The Gollum event accepts no options.
// Experimental.
type GollumOptions struct {
}

// Issue comment options.
// Experimental.
type IssueCommentOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Issues options.
// Experimental.
type IssuesOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// A GitHub Workflow job definition.
// Experimental.
type Job struct {
	// You can modify the default permissions granted to the GITHUB_TOKEN, adding or removing access as required, so that you only allow the minimum required access.
	//
	// Use `{ contents: READ }` if your job only needs to clone code.
	//
	// This is intentionally a required field since it is required in order to
	// allow workflows to run in GitHub repositories with restricted default
	// access.
	// See: https://docs.github.com/en/actions/reference/authentication-in-a-workflow#permissions-for-the-github_token
	//
	// Experimental.
	Permissions *JobPermissions `field:"required" json:"permissions" yaml:"permissions"`
	// The type of machine to run the job on.
	//
	// The machine can be either a
	// GitHub-hosted runner or a self-hosted runner.
	//
	// Example:
	//   ["ubuntu-latest"]
	//
	// Experimental.
	RunsOn *[]*string `field:"required" json:"runsOn" yaml:"runsOn"`
	// A job contains a sequence of tasks called steps.
	//
	// Steps can run commands,
	// run setup tasks, or run an action in your repository, a public repository,
	// or an action published in a Docker registry. Not all steps run actions,
	// but all actions run as a step. Each step runs in its own process in the
	// runner environment and has access to the workspace and filesystem.
	// Because steps run in their own process, changes to environment variables
	// are not preserved between steps. GitHub provides built-in steps to set up
	// and complete a job.
	// Experimental.
	Steps *[]*JobStep `field:"required" json:"steps" yaml:"steps"`
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	//
	// A concurrency group can be any
	// string or expression. The expression can use any context except for the
	// secrets context.
	// Experimental.
	Concurrency interface{} `field:"optional" json:"concurrency" yaml:"concurrency"`
	// A container to run any steps in a job that don't already specify a container.
	//
	// If you have steps that use both script and container actions,
	// the container actions will run as sibling containers on the same network
	// with the same volume mounts.
	// Experimental.
	Container *ContainerOptions `field:"optional" json:"container" yaml:"container"`
	// Prevents a workflow run from failing when a job fails.
	//
	// Set to true to
	// allow a workflow run to pass when this job fails.
	// Experimental.
	ContinueOnError *bool `field:"optional" json:"continueOnError" yaml:"continueOnError"`
	// A map of default settings that will apply to all steps in the job.
	//
	// You
	// can also set default settings for the entire workflow.
	// Experimental.
	Defaults *JobDefaults `field:"optional" json:"defaults" yaml:"defaults"`
	// A map of environment variables that are available to all steps in the job.
	//
	// You can also set environment variables for the entire workflow or an
	// individual step.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// The environment that the job references.
	//
	// All environment protection rules
	// must pass before a job referencing the environment is sent to a runner.
	// See: https://docs.github.com/en/actions/reference/environments
	//
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// You can use the if conditional to prevent a job from running unless a condition is met.
	//
	// You can use any supported context and expression to
	// create a conditional.
	// Experimental.
	If *string `field:"optional" json:"if" yaml:"if"`
	// The name of the job displayed on GitHub.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Identifies any jobs that must complete successfully before this job will run.
	//
	// It can be a string or array of strings. If a job fails, all jobs
	// that need it are skipped unless the jobs use a conditional expression
	// that causes the job to continue.
	// Experimental.
	Needs *[]*string `field:"optional" json:"needs" yaml:"needs"`
	// A map of outputs for a job.
	//
	// Job outputs are available to all downstream
	// jobs that depend on this job.
	// Experimental.
	Outputs *map[string]*JobStepOutput `field:"optional" json:"outputs" yaml:"outputs"`
	// Used to host service containers for a job in a workflow.
	//
	// Service
	// containers are useful for creating databases or cache services like Redis.
	// The runner automatically creates a Docker network and manages the life
	// cycle of the service containers.
	// Experimental.
	Services *map[string]*ContainerOptions `field:"optional" json:"services" yaml:"services"`
	// A strategy creates a build matrix for your jobs.
	//
	// You can define different
	// variations to run each job in.
	// Experimental.
	Strategy *JobStrategy `field:"optional" json:"strategy" yaml:"strategy"`
	// The maximum number of minutes to let a job run before GitHub automatically cancels it.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// Tools required for this job.
	//
	// Traslates into `actions/setup-xxx` steps at
	// the beginning of the job.
	// Experimental.
	Tools *Tools `field:"optional" json:"tools" yaml:"tools"`
}

// Default settings for all steps in the job.
// Experimental.
type JobDefaults struct {
	// Default run settings.
	// Experimental.
	Run *RunSettings `field:"optional" json:"run" yaml:"run"`
}

// A job matrix.
// Experimental.
type JobMatrix struct {
	// Each option you define in the matrix has a key and value.
	//
	// The keys you
	// define become properties in the matrix context and you can reference the
	// property in other areas of your workflow file. For example, if you define
	// the key os that contains an array of operating systems, you can use the
	// matrix.os property as the value of the runs-on keyword to create a job
	// for each operating system.
	// Experimental.
	Domain *map[string]*[]interface{} `field:"optional" json:"domain" yaml:"domain"`
	// You can remove a specific configurations defined in the build matrix using the exclude option.
	//
	// Using exclude removes a job defined by the
	// build matrix.
	// Experimental.
	Exclude *[]*map[string]interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// You can add additional configuration options to a build matrix job that already exists.
	//
	// For example, if you want to use a specific version of npm
	// when the job that uses windows-latest and version 8 of node runs, you can
	// use include to specify that additional option.
	// Experimental.
	Include *[]*map[string]interface{} `field:"optional" json:"include" yaml:"include"`
}

// Access level for workflow permission scopes.
// Experimental.
type JobPermission string

const (
	// Read-only access.
	// Experimental.
	JobPermission_READ JobPermission = "READ"
	// Read-write access.
	// Experimental.
	JobPermission_WRITE JobPermission = "WRITE"
	// No access at all.
	// Experimental.
	JobPermission_NONE JobPermission = "NONE"
)

// The available scopes and access values for workflow permissions.
//
// If you
// specify the access for any of these scopes, all those that are not
// specified are set to `JobPermission.NONE`, instead of the default behavior
// when none is specified.
// Experimental.
type JobPermissions struct {
	// Experimental.
	Actions JobPermission `field:"optional" json:"actions" yaml:"actions"`
	// Experimental.
	Checks JobPermission `field:"optional" json:"checks" yaml:"checks"`
	// Experimental.
	Contents JobPermission `field:"optional" json:"contents" yaml:"contents"`
	// Experimental.
	Deployments JobPermission `field:"optional" json:"deployments" yaml:"deployments"`
	// Experimental.
	Discussions JobPermission `field:"optional" json:"discussions" yaml:"discussions"`
	// Experimental.
	IdToken JobPermission `field:"optional" json:"idToken" yaml:"idToken"`
	// Experimental.
	Issues JobPermission `field:"optional" json:"issues" yaml:"issues"`
	// Experimental.
	Packages JobPermission `field:"optional" json:"packages" yaml:"packages"`
	// Experimental.
	Pages JobPermission `field:"optional" json:"pages" yaml:"pages"`
	// Experimental.
	PullRequests JobPermission `field:"optional" json:"pullRequests" yaml:"pullRequests"`
	// Experimental.
	RepositoryProjects JobPermission `field:"optional" json:"repositoryProjects" yaml:"repositoryProjects"`
	// Experimental.
	SecurityEvents JobPermission `field:"optional" json:"securityEvents" yaml:"securityEvents"`
	// Experimental.
	Statuses JobPermission `field:"optional" json:"statuses" yaml:"statuses"`
}

// A job step.
// Experimental.
type JobStep struct {
	// Sets environment variables for steps to use in the runner environment.
	//
	// You can also set environment variables for the entire workflow or a job.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// A unique identifier for the step.
	//
	// You can use the id to reference the
	// step in contexts.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// You can use the if conditional to prevent a job from running unless a condition is met.
	//
	// You can use any supported context and expression to
	// create a conditional.
	// Experimental.
	If *string `field:"optional" json:"if" yaml:"if"`
	// A name for your step to display on GitHub.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Runs command-line programs using the operating system's shell.
	//
	// If you do
	// not provide a name, the step name will default to the text specified in
	// the run command.
	// Experimental.
	Run *string `field:"optional" json:"run" yaml:"run"`
	// Selects an action to run as part of a step in your job.
	//
	// An action is a
	// reusable unit of code. You can use an action defined in the same
	// repository as the workflow, a public repository, or in a published Docker
	// container image.
	// Experimental.
	Uses *string `field:"optional" json:"uses" yaml:"uses"`
	// A map of the input parameters defined by the action.
	//
	// Each input parameter
	// is a key/value pair. Input parameters are set as environment variables.
	// The variable is prefixed with INPUT_ and converted to upper case.
	// Experimental.
	With *map[string]interface{} `field:"optional" json:"with" yaml:"with"`
	// Prevents a job from failing when a step fails.
	//
	// Set to true to allow a job
	// to pass when this step fails.
	// Experimental.
	ContinueOnError *bool `field:"optional" json:"continueOnError" yaml:"continueOnError"`
	// The maximum number of minutes to run the step before killing the process.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

// An output binding for a job.
// Experimental.
type JobStepOutput struct {
	// The name of the job output that is being bound.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The ID of the step that exposes the output.
	// Experimental.
	StepId *string `field:"required" json:"stepId" yaml:"stepId"`
}

// A strategy creates a build matrix for your jobs.
//
// You can define different
// variations to run each job in.
// Experimental.
type JobStrategy struct {
	// When set to true, GitHub cancels all in-progress jobs if any matrix job fails.
	//
	// Default: true.
	// Experimental.
	FailFast *bool `field:"optional" json:"failFast" yaml:"failFast"`
	// You can define a matrix of different job configurations.
	//
	// A matrix allows
	// you to create multiple jobs by performing variable substitution in a
	// single job definition. For example, you can use a matrix to create jobs
	// for more than one supported version of a programming language, operating
	// system, or tool. A matrix reuses the job's configuration and creates a
	// job for each matrix you configure.
	//
	// A job matrix can generate a maximum of 256 jobs per workflow run. This
	// limit also applies to self-hosted runners.
	// Experimental.
	Matrix *JobMatrix `field:"optional" json:"matrix" yaml:"matrix"`
	// The maximum number of jobs that can run simultaneously when using a matrix job strategy.
	//
	// By default, GitHub will maximize the number of jobs
	// run in parallel depending on the available runners on GitHub-hosted
	// virtual machines.
	// Experimental.
	MaxParallel *float64 `field:"optional" json:"maxParallel" yaml:"maxParallel"`
}

// label options.
// Experimental.
type LabelOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Milestone options.
// Experimental.
type MilestoneOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// The Page build event accepts no options.
// Experimental.
type PageBuildOptions struct {
}

// Project card options.
// Experimental.
type ProjectCardOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Probject column options.
// Experimental.
type ProjectColumnOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Project options.
// Experimental.
type ProjectOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// The Public event accepts no options.
// Experimental.
type PublicOptions struct {
}

// Pull request options.
// Experimental.
type PullRequestOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Pull request review comment options.
// Experimental.
type PullRequestReviewCommentOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Pull request review options.
// Experimental.
type PullRequestReviewOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Pull request target options.
// Experimental.
type PullRequestTargetOptions struct {
	// When using the push and pull_request events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// When using the push and pull_request events, you can configure a workflow to run when at least one file does not match paths-ignore or at least one modified file matches the configured paths.
	//
	// Path filters are not
	// evaluated for pushes to tags.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// When using the push and pull_request events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Options for push-like events.
// Experimental.
type PushOptions struct {
	// When using the push and pull_request events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// When using the push and pull_request events, you can configure a workflow to run when at least one file does not match paths-ignore or at least one modified file matches the configured paths.
	//
	// Path filters are not
	// evaluated for pushes to tags.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
	// When using the push and pull_request events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

// Registry package options.
// Experimental.
type RegistryPackageOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Release options.
// Experimental.
type ReleaseOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Repository dispatch options.
// Experimental.
type RepositoryDispatchOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// Run settings for a job.
// Experimental.
type RunSettings struct {
	// Which shell to use for running the step.
	//
	// Example:
	//   "bash"
	//
	// Experimental.
	Shell *string `field:"optional" json:"shell" yaml:"shell"`
	// Working directory to use when running the step.
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

// The Status event accepts no options.
// Experimental.
type StatusOptions struct {
}

// A generic step.
// Experimental.
type Step struct {
	// Sets environment variables for steps to use in the runner environment.
	//
	// You can also set environment variables for the entire workflow or a job.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// A unique identifier for the step.
	//
	// You can use the id to reference the
	// step in contexts.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// You can use the if conditional to prevent a job from running unless a condition is met.
	//
	// You can use any supported context and expression to
	// create a conditional.
	// Experimental.
	If *string `field:"optional" json:"if" yaml:"if"`
	// A name for your step to display on GitHub.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Runs command-line programs using the operating system's shell.
	//
	// If you do
	// not provide a name, the step name will default to the text specified in
	// the run command.
	// Experimental.
	Run *string `field:"optional" json:"run" yaml:"run"`
	// Selects an action to run as part of a step in your job.
	//
	// An action is a
	// reusable unit of code. You can use an action defined in the same
	// repository as the workflow, a public repository, or in a published Docker
	// container image.
	// Experimental.
	Uses *string `field:"optional" json:"uses" yaml:"uses"`
	// A map of the input parameters defined by the action.
	//
	// Each input parameter
	// is a key/value pair. Input parameters are set as environment variables.
	// The variable is prefixed with INPUT_ and converted to upper case.
	// Experimental.
	With *map[string]interface{} `field:"optional" json:"with" yaml:"with"`
}

// Version requirement for tools.
// Experimental.
type ToolRequirement struct {
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
}

// Supported tools.
// Experimental.
type Tools struct {
	// Setup .NET Core.
	// Experimental.
	Dotnet *ToolRequirement `field:"optional" json:"dotnet" yaml:"dotnet"`
	// Setup golang.
	// Experimental.
	Go *ToolRequirement `field:"optional" json:"go" yaml:"go"`
	// Setup java (temurin distribution).
	// Experimental.
	Java *ToolRequirement `field:"optional" json:"java" yaml:"java"`
	// Setup node.js.
	// Experimental.
	Node *ToolRequirement `field:"optional" json:"node" yaml:"node"`
	// Setup python.
	// Experimental.
	Python *ToolRequirement `field:"optional" json:"python" yaml:"python"`
}

// The set of available triggers for GitHub Workflows.
// See: https://docs.github.com/en/actions/reference/events-that-trigger-workflows
//
// Experimental.
type Triggers struct {
	// Runs your workflow anytime the branch_protection_rule event occurs.
	// Experimental.
	BranchProtectionRule *BranchProtectionRuleOptions `field:"optional" json:"branchProtectionRule" yaml:"branchProtectionRule"`
	// Runs your workflow anytime the check_run event occurs.
	// Experimental.
	CheckRun *CheckRunOptions `field:"optional" json:"checkRun" yaml:"checkRun"`
	// Runs your workflow anytime the check_suite event occurs.
	// Experimental.
	CheckSuite *CheckSuiteOptions `field:"optional" json:"checkSuite" yaml:"checkSuite"`
	// Runs your workflow anytime someone creates a branch or tag, which triggers the create event.
	// Experimental.
	Create *CreateOptions `field:"optional" json:"create" yaml:"create"`
	// Runs your workflow anytime someone deletes a branch or tag, which triggers the delete event.
	// Experimental.
	Delete *DeleteOptions `field:"optional" json:"delete" yaml:"delete"`
	// Runs your workflow anytime someone creates a deployment, which triggers the deployment event.
	//
	// Deployments created with a commit SHA may not have
	// a Git ref.
	// Experimental.
	Deployment *DeploymentOptions `field:"optional" json:"deployment" yaml:"deployment"`
	// Runs your workflow anytime a third party provides a deployment status, which triggers the deployment_status event.
	//
	// Deployments created with a
	// commit SHA may not have a Git ref.
	// Experimental.
	DeploymentStatus *DeploymentStatusOptions `field:"optional" json:"deploymentStatus" yaml:"deploymentStatus"`
	// Runs your workflow anytime the discussion event occurs.
	//
	// More than one activity type triggers this event.
	// See: https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	//
	// Experimental.
	Discussion *DiscussionOptions `field:"optional" json:"discussion" yaml:"discussion"`
	// Runs your workflow anytime the discussion_comment event occurs.
	//
	// More than one activity type triggers this event.
	// See: https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	//
	// Experimental.
	DiscussionComment *DiscussionCommentOptions `field:"optional" json:"discussionComment" yaml:"discussionComment"`
	// Runs your workflow anytime when someone forks a repository, which triggers the fork event.
	// Experimental.
	Fork *ForkOptions `field:"optional" json:"fork" yaml:"fork"`
	// Runs your workflow when someone creates or updates a Wiki page, which triggers the gollum event.
	// Experimental.
	Gollum *GollumOptions `field:"optional" json:"gollum" yaml:"gollum"`
	// Runs your workflow anytime the issue_comment event occurs.
	// Experimental.
	IssueComment *IssueCommentOptions `field:"optional" json:"issueComment" yaml:"issueComment"`
	// Runs your workflow anytime the issues event occurs.
	// Experimental.
	Issues *IssuesOptions `field:"optional" json:"issues" yaml:"issues"`
	// Runs your workflow anytime the label event occurs.
	// Experimental.
	Label *LabelOptions `field:"optional" json:"label" yaml:"label"`
	// Runs your workflow anytime the milestone event occurs.
	// Experimental.
	Milestone *MilestoneOptions `field:"optional" json:"milestone" yaml:"milestone"`
	// Runs your workflow anytime someone pushes to a GitHub Pages-enabled branch, which triggers the page_build event.
	// Experimental.
	PageBuild *PageBuildOptions `field:"optional" json:"pageBuild" yaml:"pageBuild"`
	// Runs your workflow anytime the project event occurs.
	// Experimental.
	Project *ProjectOptions `field:"optional" json:"project" yaml:"project"`
	// Runs your workflow anytime the project_card event occurs.
	// Experimental.
	ProjectCard *ProjectCardOptions `field:"optional" json:"projectCard" yaml:"projectCard"`
	// Runs your workflow anytime the project_column event occurs.
	// Experimental.
	ProjectColumn *ProjectColumnOptions `field:"optional" json:"projectColumn" yaml:"projectColumn"`
	// Runs your workflow anytime someone makes a private repository public, which triggers the public event.
	// Experimental.
	Public *PublicOptions `field:"optional" json:"public" yaml:"public"`
	// Runs your workflow anytime the pull_request event occurs.
	// Experimental.
	PullRequest *PullRequestOptions `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// Runs your workflow anytime the pull_request_review event occurs.
	// Experimental.
	PullRequestReview *PullRequestReviewOptions `field:"optional" json:"pullRequestReview" yaml:"pullRequestReview"`
	// Runs your workflow anytime a comment on a pull request's unified diff is modified, which triggers the pull_request_review_comment event.
	// Experimental.
	PullRequestReviewComment *PullRequestReviewCommentOptions `field:"optional" json:"pullRequestReviewComment" yaml:"pullRequestReviewComment"`
	// This event runs in the context of the base of the pull request, rather than in the merge commit as the pull_request event does.
	//
	// This prevents
	// executing unsafe workflow code from the head of the pull request that
	// could alter your repository or steal any secrets you use in your workflow.
	// This event allows you to do things like create workflows that label and
	// comment on pull requests based on the contents of the event payload.
	//
	// WARNING: The `pull_request_target` event is granted read/write repository
	// token and can access secrets, even when it is triggered from a fork.
	// Although the workflow runs in the context of the base of the pull request,
	// you should make sure that you do not check out, build, or run untrusted
	// code from the pull request with this event. Additionally, any caches
	// share the same scope as the base branch, and to help prevent cache
	// poisoning, you should not save the cache if there is a possibility that
	// the cache contents were altered.
	// See: https://securitylab.github.com/research/github-actions-preventing-pwn-requests
	//
	// Experimental.
	PullRequestTarget *PullRequestTargetOptions `field:"optional" json:"pullRequestTarget" yaml:"pullRequestTarget"`
	// Runs your workflow when someone pushes to a repository branch, which triggers the push event.
	// Experimental.
	Push *PushOptions `field:"optional" json:"push" yaml:"push"`
	// Runs your workflow anytime a package is published or updated.
	// Experimental.
	RegistryPackage *RegistryPackageOptions `field:"optional" json:"registryPackage" yaml:"registryPackage"`
	// Runs your workflow anytime the release event occurs.
	// Experimental.
	Release *ReleaseOptions `field:"optional" json:"release" yaml:"release"`
	// You can use the GitHub API to trigger a webhook event called repository_dispatch when you want to trigger a workflow for activity that happens outside of GitHub.
	// Experimental.
	RepositoryDispatch *RepositoryDispatchOptions `field:"optional" json:"repositoryDispatch" yaml:"repositoryDispatch"`
	// You can schedule a workflow to run at specific UTC times using POSIX cron syntax.
	//
	// Scheduled workflows run on the latest commit on the default or
	// base branch. The shortest interval you can run scheduled workflows is
	// once every 5 minutes.
	// See: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/crontab.html#tag_20_25_07
	//
	// Experimental.
	Schedule *[]*CronScheduleOptions `field:"optional" json:"schedule" yaml:"schedule"`
	// Runs your workflow anytime the status of a Git commit changes, which triggers the status event.
	// Experimental.
	Status *StatusOptions `field:"optional" json:"status" yaml:"status"`
	// Runs your workflow anytime the watch event occurs.
	// Experimental.
	Watch *WatchOptions `field:"optional" json:"watch" yaml:"watch"`
	// Can be called from another workflow.
	// See: https://docs.github.com/en/actions/learn-github-actions/reusing-workflows
	//
	// Experimental.
	WorkflowCall *WorkflowCallOptions `field:"optional" json:"workflowCall" yaml:"workflowCall"`
	// You can configure custom-defined input properties, default input values, and required inputs for the event directly in your workflow.
	//
	// When the
	// workflow runs, you can access the input values in the github.event.inputs
	// context.
	// Experimental.
	WorkflowDispatch *WorkflowDispatchOptions `field:"optional" json:"workflowDispatch" yaml:"workflowDispatch"`
	// This event occurs when a workflow run is requested or completed, and allows you to execute a workflow based on the finished result of another workflow.
	//
	// A workflow run is triggered regardless of the result of the
	// previous workflow.
	// Experimental.
	WorkflowRun *WorkflowRunOptions `field:"optional" json:"workflowRun" yaml:"workflowRun"`
}

// Watch options.
// Experimental.
type WatchOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

// The Workflow Call event accepts no options.
// Experimental.
type WorkflowCallOptions struct {
}

// The Workflow dispatch event accepts no options.
// Experimental.
type WorkflowDispatchOptions struct {
}

// Workflow run options.
// Experimental.
type WorkflowRunOptions struct {
	// Which branches or branch-ignore to limit the trigger to.
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Which workflow to trigger on.
	// Experimental.
	Workflows *[]*string `field:"optional" json:"workflows" yaml:"workflows"`
}

