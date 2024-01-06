package workflows


// This contains the fields that are common amongst both: - JobStep, which is a step that is part of a Job in Github Actions.
//
// This is by far the most common use case.
// - The metadata file `action.yaml` that is used to define an Action when you are creating one. As in, if you were creating an Action to be used in a JobStep.
// There is some overlap between the two, and this captures that overlap.
// See: https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions
//
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
}

