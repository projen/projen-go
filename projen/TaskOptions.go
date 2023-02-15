// CDK for software projects
package projen


// Experimental.
type TaskOptions struct {
	// A shell command which determines if the this task should be executed.
	//
	// If
	// the program exits with a zero exit code, steps will be executed. A non-zero
	// code means that task will be skipped.
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The working directory for all steps in this task (unless overridden by the step).
	// Experimental.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// The description of this build command.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Defines environment variables for the execution of this task.
	//
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// A set of environment variables that must be defined in order to execute this task.
	//
	// Task execution will fail if one of these is not defined.
	// Experimental.
	RequiredEnv *[]*string `field:"optional" json:"requiredEnv" yaml:"requiredEnv"`
	// Shell command to execute as the first command of the task.
	// Experimental.
	Exec *string `field:"optional" json:"exec" yaml:"exec"`
	// Should the provided `exec` shell command receive args passed to the task.
	// See: {@link TaskStepOptions.receiveArgs}
	//
	// Experimental.
	ReceiveArgs *bool `field:"optional" json:"receiveArgs" yaml:"receiveArgs"`
	// List of task steps to run.
	// Experimental.
	Steps *[]*TaskStep `field:"optional" json:"steps" yaml:"steps"`
}

