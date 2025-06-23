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

