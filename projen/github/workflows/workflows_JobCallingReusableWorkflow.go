package workflows


// A GitHub Workflow Job calling a reusable workflow.
// Experimental.
type JobCallingReusableWorkflow struct {
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
	// The location and version of a reusable workflow file to run as a job.
	// Experimental.
	Uses *string `field:"required" json:"uses" yaml:"uses"`
	// When a job is used to call a reusable workflow, you can use secrets to provide a map of secrets that are passed to the called workflow.
	//
	// Use the 'inherit' keyword to pass all the calling workflow's secrets to the called workflow.
	// Experimental.
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// When a job is used to call a reusable workflow, you can use with to provide a map of inputs that are passed to the called workflow.
	//
	// Allowed expression contexts: `github`, and `needs`.
	// Experimental.
	With *map[string]interface{} `field:"optional" json:"with" yaml:"with"`
}

