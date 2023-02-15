package circleci


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

