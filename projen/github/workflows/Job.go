package workflows


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
	// Concurrency ensures that only a single job or workflow using the same concurrency group will run at a time.
	//
	// A concurrency group can be any
	// string or expression. The expression can use any context except for the
	// secrets context.
	// Experimental.
	Concurrency interface{} `field:"optional" json:"concurrency" yaml:"concurrency"`
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
	// A strategy creates a build matrix for your jobs.
	//
	// You can define different
	// variations to run each job in.
	// Experimental.
	Strategy *JobStrategy `field:"optional" json:"strategy" yaml:"strategy"`
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
	// The maximum number of minutes to let a job run before GitHub automatically cancels it.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// Tools required for this job.
	//
	// Translates into `actions/setup-xxx` steps at
	// the beginning of the job.
	// Experimental.
	Tools *Tools `field:"optional" json:"tools" yaml:"tools"`
}

