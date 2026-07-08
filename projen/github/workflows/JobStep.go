package workflows


// JobSteps run as part of a GitHub Workflow Job.
// See: https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idsteps
//
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
	// Overrides the default shell settings in the runner's operating system and the job's default.
	//
	// Refer to GitHub documentation for allowed values.
	// See: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idstepsshell
	//
	// Experimental.
	Shell *string `field:"optional" json:"shell" yaml:"shell"`
	// Specifies a working directory for a step.
	//
	// Overrides a job's working directory.
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
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
	// Runs a step asynchronously so the job continues to the next step without waiting for it to finish.
	//
	// Use for long-running processes, such as
	// databases, servers, or monitoring tasks, that need to run alongside other
	// steps.
	//
	// Synchronize with background steps later using `wait` or `waitAll`, or
	// stop them with `cancel`. Give the step an `id` so it can be referenced.
	//
	// A maximum of 10 background steps can run concurrently in a single job;
	// additional background steps are queued until a slot is free.
	// See: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idstepsbackground
	//
	// Experimental.
	Background *bool `field:"optional" json:"background" yaml:"background"`
	// Gracefully terminates a running background step, referenced by its `id`.
	//
	// The runner sends the step's process a termination signal (SIGTERM) so it
	// can clean up, and forcibly stops it (SIGKILL) if it does not exit within
	// a short grace period.
	// See: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idstepscancel
	//
	// Experimental.
	Cancel *string `field:"optional" json:"cancel" yaml:"cancel"`
	// Prevents a job from failing when a step fails.
	//
	// Set to true to allow a job
	// to pass when this step fails.
	// Experimental.
	ContinueOnError *bool `field:"optional" json:"continueOnError" yaml:"continueOnError"`
	// Runs a group of steps concurrently, then waits for all of them to finish before continuing.
	//
	// This is shorthand for declaring each step with
	// `background: true` followed by a `wait` step.
	//
	// Use this when you have a self-contained group of independent steps that
	// can all run at the same time and don't need to be referenced
	// individually. Use `background` instead when you need finer control, such
	// as starting a long-running process that stays up while later steps run.
	//
	// Each step in the group is subject to the same 10-step concurrency limit
	// as other background steps. Cannot be used inside a composite action.
	// See: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idstepsparallel
	//
	// Experimental.
	Parallel *[]*JobStep `field:"optional" json:"parallel" yaml:"parallel"`
	// The maximum number of minutes to run the step before killing the process.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// Pauses the job until one or more background steps complete. Provide the `id`s of the background steps to wait for.
	//
	// This step performs no work itself; it only blocks until the referenced
	// background steps finish. If a referenced background step failed, the
	// `wait` step fails too.
	// See: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idstepswait
	//
	// Experimental.
	Wait *[]*string `field:"optional" json:"wait" yaml:"wait"`
	// Pauses the job until all active background steps complete.
	//
	// Fails if any
	// of the background steps it waits on failed, unless `continueOnError` is
	// set on this step.
	// See: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions#jobsjob_idstepswait-all
	//
	// Experimental.
	WaitAll *bool `field:"optional" json:"waitAll" yaml:"waitAll"`
}

