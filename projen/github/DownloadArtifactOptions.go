package github


// Experimental.
type DownloadArtifactOptions struct {
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
	// Prevents a job from failing when a step fails.
	//
	// Set to true to allow a job
	// to pass when this step fails.
	// Experimental.
	ContinueOnError *bool `field:"optional" json:"continueOnError" yaml:"continueOnError"`
	// The maximum number of minutes to run the step before killing the process.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// Options for `download-artifact`.
	// Experimental.
	With *DownloadArtifactWith `field:"required" json:"with" yaml:"with"`
}

