package workflows


// Branch Protection Rule options.
// Experimental.
type BranchProtectionRuleOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Check run options.
// Experimental.
type CheckRunOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Check suite options.
// Experimental.
type CheckSuiteOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Credentials to use to authenticate to Docker registries.
// Experimental.
type ContainerCredentials struct {
	// The password.
	// Experimental.
	Password *string `json:"password"`
	// The username.
	// Experimental.
	Username *string `json:"username"`
}

// Options petaining to container environments.
// Experimental.
type ContainerOptions struct {
	// The Docker image to use as the container to run the action.
	//
	// The value can
	// be the Docker Hub image name or a registry name.
	// Experimental.
	Image *string `json:"image"`
	// f the image's container registry requires authentication to pull the image, you can use credentials to set a map of the username and password.
	//
	// The credentials are the same values that you would provide to the docker
	// login command.
	// Experimental.
	Credentials *ContainerCredentials `json:"credentials"`
	// Sets a map of environment variables in the container.
	// Experimental.
	Env *map[string]*string `json:"env"`
	// Additional Docker container resource options.
	// See: https://docs.docker.com/engine/reference/commandline/create/#options
	//
	// Experimental.
	Options *[]*string `json:"options"`
	// Sets an array of ports to expose on the container.
	// Experimental.
	Ports *[]*float64 `json:"ports"`
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
	Volumes *[]*string `json:"volumes"`
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
	Cron *string `json:"cron"`
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
	Types *[]*string `json:"types"`
}

// Discussion options.
// Experimental.
type DiscussionOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
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
	Types *[]*string `json:"types"`
}

// Issues options.
// Experimental.
type IssuesOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
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
	Permissions *JobPermissions `json:"permissions"`
	// The type of machine to run the job on.
	//
	// The machine can be either a
	// GitHub-hosted runner or a self-hosted runner.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	RunsOn *[]*string `json:"runsOn"`
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
	Steps *[]*JobStep `json:"steps"`
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	//
	// A concurrency group can be any
	// string or expression. The expression can use any context except for the
	// secrets context.
	// Experimental.
	Concurrency interface{} `json:"concurrency"`
	// A container to run any steps in a job that don't already specify a container.
	//
	// If you have steps that use both script and container actions,
	// the container actions will run as sibling containers on the same network
	// with the same volume mounts.
	// Experimental.
	Container *ContainerOptions `json:"container"`
	// Prevents a workflow run from failing when a job fails.
	//
	// Set to true to
	// allow a workflow run to pass when this job fails.
	// Experimental.
	ContinueOnError *bool `json:"continueOnError"`
	// A map of default settings that will apply to all steps in the job.
	//
	// You
	// can also set default settings for the entire workflow.
	// Experimental.
	Defaults *JobDefaults `json:"defaults"`
	// A map of environment variables that are available to all steps in the job.
	//
	// You can also set environment variables for the entire workflow or an
	// individual step.
	// Experimental.
	Env *map[string]*string `json:"env"`
	// The environment that the job references.
	//
	// All environment protection rules
	// must pass before a job referencing the environment is sent to a runner.
	// See: https://docs.github.com/en/actions/reference/environments
	//
	// Experimental.
	Environment interface{} `json:"environment"`
	// You can use the if conditional to prevent a job from running unless a condition is met.
	//
	// You can use any supported context and expression to
	// create a conditional.
	// Experimental.
	If *string `json:"if"`
	// The name of the job displayed on GitHub.
	// Experimental.
	Name *string `json:"name"`
	// Identifies any jobs that must complete successfully before this job will run.
	//
	// It can be a string or array of strings. If a job fails, all jobs
	// that need it are skipped unless the jobs use a conditional expression
	// that causes the job to continue.
	// Experimental.
	Needs *[]*string `json:"needs"`
	// A map of outputs for a job.
	//
	// Job outputs are available to all downstream
	// jobs that depend on this job.
	// Experimental.
	Outputs *map[string]*JobStepOutput `json:"outputs"`
	// Used to host service containers for a job in a workflow.
	//
	// Service
	// containers are useful for creating databases or cache services like Redis.
	// The runner automatically creates a Docker network and manages the life
	// cycle of the service containers.
	// Experimental.
	Services *map[string]*ContainerOptions `json:"services"`
	// A strategy creates a build matrix for your jobs.
	//
	// You can define different
	// variations to run each job in.
	// Experimental.
	Strategy *JobStrategy `json:"strategy"`
	// The maximum number of minutes to let a job run before GitHub automatically cancels it.
	// Experimental.
	TimeoutMinutes *float64 `json:"timeoutMinutes"`
}

// Default settings for all steps in the job.
// Experimental.
type JobDefaults struct {
	// Default run settings.
	// Experimental.
	Run *RunSettings `json:"run"`
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
	Domain *map[string]*[]*string `json:"domain"`
	// You can remove a specific configurations defined in the build matrix using the exclude option.
	//
	// Using exclude removes a job defined by the
	// build matrix.
	// Experimental.
	Exclude *[]*map[string]*string `json:"exclude"`
	// You can add additional configuration options to a build matrix job that already exists.
	//
	// For example, if you want to use a specific version of npm
	// when the job that uses windows-latest and version 8 of node runs, you can
	// use include to specify that additional option.
	// Experimental.
	Include *[]*map[string]*string `json:"include"`
}

// Access level for workflow permission scopes.
// Experimental.
type JobPermission string

const (
	JobPermission_READ JobPermission = "READ"
	JobPermission_WRITE JobPermission = "WRITE"
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
	Actions JobPermission `json:"actions"`
	// Experimental.
	Checks JobPermission `json:"checks"`
	// Experimental.
	Contents JobPermission `json:"contents"`
	// Experimental.
	Deployments JobPermission `json:"deployments"`
	// Experimental.
	Issues JobPermission `json:"issues"`
	// Experimental.
	Packages JobPermission `json:"packages"`
	// Experimental.
	PullRequests JobPermission `json:"pullRequests"`
	// Experimental.
	RepositoryProjects JobPermission `json:"repositoryProjects"`
	// Experimental.
	SecurityEvents JobPermission `json:"securityEvents"`
	// Experimental.
	Statuses JobPermission `json:"statuses"`
}

// A job step.
// Experimental.
type JobStep struct {
	// Prevents a job from failing when a step fails.
	//
	// Set to true to allow a job
	// to pass when this step fails.
	// Experimental.
	ContinueOnError *bool `json:"continueOnError"`
	// Sets environment variables for steps to use in the runner environment.
	//
	// You can also set environment variables for the entire workflow or a job.
	// Experimental.
	Env *map[string]*string `json:"env"`
	// A unique identifier for the step.
	//
	// You can use the id to reference the
	// step in contexts.
	// Experimental.
	Id *string `json:"id"`
	// You can use the if conditional to prevent a job from running unless a condition is met.
	//
	// You can use any supported context and expression to
	// create a conditional.
	// Experimental.
	If *string `json:"if"`
	// A name for your step to display on GitHub.
	// Experimental.
	Name *string `json:"name"`
	// Runs command-line programs using the operating system's shell.
	//
	// If you do
	// not provide a name, the step name will default to the text specified in
	// the run command.
	// Experimental.
	Run *string `json:"run"`
	// The maximum number of minutes to run the step before killing the process.
	// Experimental.
	TimeoutMinutes *float64 `json:"timeoutMinutes"`
	// Selects an action to run as part of a step in your job.
	//
	// An action is a
	// reusable unit of code. You can use an action defined in the same
	// repository as the workflow, a public repository, or in a published Docker
	// container image.
	// Experimental.
	Uses *string `json:"uses"`
	// A map of the input parameters defined by the action.
	//
	// Each input parameter
	// is a key/value pair. Input parameters are set as environment variables.
	// The variable is prefixed with INPUT_ and converted to upper case.
	// Experimental.
	With *map[string]interface{} `json:"with"`
}

// An output binding for a job.
// Experimental.
type JobStepOutput struct {
	// The name of the job output that is being bound.
	// Experimental.
	OutputName *string `json:"outputName"`
	// The ID of the step that exposes the output.
	// Experimental.
	StepId *string `json:"stepId"`
}

// A strategy creates a build matrix for your jobs.
//
// You can define different
// variations to run each job in.
// Experimental.
type JobStrategy struct {
	// When set to true, GitHub cancels all in-progress jobs if any matrix job fails.
	//
	// Default: true
	// Experimental.
	FailFast *bool `json:"failFast"`
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
	Matrix *JobMatrix `json:"matrix"`
	// The maximum number of jobs that can run simultaneously when using a matrix job strategy.
	//
	// By default, GitHub will maximize the number of jobs
	// run in parallel depending on the available runners on GitHub-hosted
	// virtual machines.
	// Experimental.
	MaxParallel *float64 `json:"maxParallel"`
}

// label options.
// Experimental.
type LabelOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Milestone options.
// Experimental.
type MilestoneOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
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
	Types *[]*string `json:"types"`
}

// Probject column options.
// Experimental.
type ProjectColumnOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Project options.
// Experimental.
type ProjectOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
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
	Types *[]*string `json:"types"`
}

// Pull request review comment options.
// Experimental.
type PullRequestReviewCommentOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Pull request review options.
// Experimental.
type PullRequestReviewOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
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
	Branches *[]*string `json:"branches"`
	// When using the push and pull_request events, you can configure a workflow to run when at least one file does not match paths-ignore or at least one modified file matches the configured paths.
	//
	// Path filters are not
	// evaluated for pushes to tags.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Paths *[]*string `json:"paths"`
	// When using the push and pull_request events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Tags *[]*string `json:"tags"`
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
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
	Branches *[]*string `json:"branches"`
	// When using the push and pull_request events, you can configure a workflow to run when at least one file does not match paths-ignore or at least one modified file matches the configured paths.
	//
	// Path filters are not
	// evaluated for pushes to tags.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Paths *[]*string `json:"paths"`
	// When using the push and pull_request events, you can configure a workflow to run on specific branches or tags.
	//
	// For a pull_request event, only
	// branches and tags on the base are evaluated. If you define only tags or
	// only branches, the workflow won't run for events affecting the undefined
	// Git ref.
	// See: https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#filter-pattern-cheat-sheet
	//
	// Experimental.
	Tags *[]*string `json:"tags"`
}

// Registry package options.
// Experimental.
type RegistryPackageOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Release options.
// Experimental.
type ReleaseOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Repository dispatch options.
// Experimental.
type RepositoryDispatchOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

// Run settings for a job.
// Experimental.
type RunSettings struct {
	// Which shell to use for running the step.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Shell *string `json:"shell"`
	// Working directory to use when running the step.
	// Experimental.
	WorkingDirectory *string `json:"workingDirectory"`
}

// The Status event accepts no options.
// Experimental.
type StatusOptions struct {
}

// The set of available triggers for GitHub Workflows.
// See: https://docs.github.com/en/actions/reference/events-that-trigger-workflows
//
// Experimental.
type Triggers struct {
	// Runs your workflow anytime the branch_protection_rule event occurs.
	// Experimental.
	BranchProtectionRule *BranchProtectionRuleOptions `json:"branchProtectionRule"`
	// Runs your workflow anytime the check_run event occurs.
	// Experimental.
	CheckRun *CheckRunOptions `json:"checkRun"`
	// Runs your workflow anytime the check_suite event occurs.
	// Experimental.
	CheckSuite *CheckSuiteOptions `json:"checkSuite"`
	// Runs your workflow anytime someone creates a branch or tag, which triggers the create event.
	// Experimental.
	Create *CreateOptions `json:"create"`
	// Runs your workflow anytime someone deletes a branch or tag, which triggers the delete event.
	// Experimental.
	Delete *DeleteOptions `json:"delete"`
	// Runs your workflow anytime someone creates a deployment, which triggers the deployment event.
	//
	// Deployments created with a commit SHA may not have
	// a Git ref.
	// Experimental.
	Deployment *DeploymentOptions `json:"deployment"`
	// Runs your workflow anytime a third party provides a deployment status, which triggers the deployment_status event.
	//
	// Deployments created with a
	// commit SHA may not have a Git ref.
	// Experimental.
	DeploymentStatus *DeploymentStatusOptions `json:"deploymentStatus"`
	// Runs your workflow anytime the discussion event occurs.
	//
	// More than one activity type triggers this event.
	// See: https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	//
	// Experimental.
	Discussion *DiscussionOptions `json:"discussion"`
	// Runs your workflow anytime the discussion_comment event occurs.
	//
	// More than one activity type triggers this event.
	// See: https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions
	//
	// Experimental.
	DiscussionComment *DiscussionCommentOptions `json:"discussionComment"`
	// Runs your workflow anytime when someone forks a repository, which triggers the fork event.
	// Experimental.
	Fork *ForkOptions `json:"fork"`
	// Runs your workflow when someone creates or updates a Wiki page, which triggers the gollum event.
	// Experimental.
	Gollum *GollumOptions `json:"gollum"`
	// Runs your workflow anytime the issue_comment event occurs.
	// Experimental.
	IssueComment *IssueCommentOptions `json:"issueComment"`
	// Runs your workflow anytime the issues event occurs.
	// Experimental.
	Issues *IssuesOptions `json:"issues"`
	// Runs your workflow anytime the label event occurs.
	// Experimental.
	Label *LabelOptions `json:"label"`
	// Runs your workflow anytime the milestone event occurs.
	// Experimental.
	Milestone *MilestoneOptions `json:"milestone"`
	// Runs your workflow anytime someone pushes to a GitHub Pages-enabled branch, which triggers the page_build event.
	// Experimental.
	PageBuild *PageBuildOptions `json:"pageBuild"`
	// Runs your workflow anytime the project event occurs.
	// Experimental.
	Project *ProjectOptions `json:"project"`
	// Runs your workflow anytime the project_card event occurs.
	// Experimental.
	ProjectCard *ProjectCardOptions `json:"projectCard"`
	// Runs your workflow anytime the project_column event occurs.
	// Experimental.
	ProjectColumn *ProjectColumnOptions `json:"projectColumn"`
	// Runs your workflow anytime someone makes a private repository public, which triggers the public event.
	// Experimental.
	Public *PublicOptions `json:"public"`
	// Runs your workflow anytime the pull_request event occurs.
	// Experimental.
	PullRequest *PullRequestOptions `json:"pullRequest"`
	// Runs your workflow anytime the pull_request_review event occurs.
	// Experimental.
	PullRequestReview *PullRequestReviewOptions `json:"pullRequestReview"`
	// Runs your workflow anytime a comment on a pull request's unified diff is modified, which triggers the pull_request_review_comment event.
	// Experimental.
	PullRequestReviewComment *PullRequestReviewCommentOptions `json:"pullRequestReviewComment"`
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
	PullRequestTarget *PullRequestTargetOptions `json:"pullRequestTarget"`
	// Runs your workflow when someone pushes to a repository branch, which triggers the push event.
	// Experimental.
	Push *PushOptions `json:"push"`
	// Runs your workflow anytime a package is published or updated.
	// Experimental.
	RegistryPackage *RegistryPackageOptions `json:"registryPackage"`
	// Runs your workflow anytime the release event occurs.
	// Experimental.
	Release *ReleaseOptions `json:"release"`
	// You can use the GitHub API to trigger a webhook event called repository_dispatch when you want to trigger a workflow for activity that happens outside of GitHub.
	// Experimental.
	RepositoryDispatch *RepositoryDispatchOptions `json:"repositoryDispatch"`
	// You can schedule a workflow to run at specific UTC times using POSIX cron syntax.
	//
	// Scheduled workflows run on the latest commit on the default or
	// base branch. The shortest interval you can run scheduled workflows is
	// once every 5 minutes.
	// See: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/crontab.html#tag_20_25_07
	//
	// Experimental.
	Schedule *[]*CronScheduleOptions `json:"schedule"`
	// Runs your workflow anytime the status of a Git commit changes, which triggers the status event.
	// Experimental.
	Status *StatusOptions `json:"status"`
	// Runs your workflow anytime the watch event occurs.
	// Experimental.
	Watch *WatchOptions `json:"watch"`
	// Can be called from another workflow.
	// See: https://docs.github.com/en/actions/learn-github-actions/reusing-workflows
	//
	// Experimental.
	WorkflowCall *WorkflowCallOptions `json:"workflowCall"`
	// You can configure custom-defined input properties, default input values, and required inputs for the event directly in your workflow.
	//
	// When the
	// workflow runs, you can access the input values in the github.event.inputs
	// context.
	// Experimental.
	WorkflowDispatch *WorkflowDispatchOptions `json:"workflowDispatch"`
	// This event occurs when a workflow run is requested or completed, and allows you to execute a workflow based on the finished result of another workflow.
	//
	// A workflow run is triggered regardless of the result of the
	// previous workflow.
	// Experimental.
	WorkflowRun *WorkflowRunOptions `json:"workflowRun"`
}

// Watch options.
// Experimental.
type WatchOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
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
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `json:"types"`
}

