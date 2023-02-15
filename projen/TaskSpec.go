// CDK for software projects
package projen


// Specification of a single task.
// Experimental.
type TaskSpec struct {
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
	// Task name.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Task steps.
	// Experimental.
	Steps *[]*TaskStep `field:"optional" json:"steps" yaml:"steps"`
}

