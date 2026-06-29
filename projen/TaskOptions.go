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
	// Default: - process.cwd()
	//
	// Experimental.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// The description of this build command.
	// Default: - the task name.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Defines environment variables for the execution of this task.
	//
	// Values in this map will be evaluated in a shell, so you can do stuff like `$(echo "foo")`.
	// Default: {}.
	//
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// A set of environment variables that must be defined in order to execute this task.
	//
	// Task execution will fail if one of these is not defined.
	// Experimental.
	RequiredEnv *[]*string `field:"optional" json:"requiredEnv" yaml:"requiredEnv"`
	// The shell used to run this task's commands, including its `condition` and `$(...)` environment evaluation. Use {@link TaskShell} to pick a built-in or an explicit invocation. Set at project, task or step level; the nearest declared level wins.
	// Default: - inherited from the task/project, otherwise the built-in projen shell.
	//
	// Experimental.
	Shell TaskShell `field:"optional" json:"shell" yaml:"shell"`
	// Should the provided `exec` shell command receive fixed args.
	// See: {@link TaskStepOptions.args }
	//
	// Default: - no arguments are passed to the step.
	//
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Shell command to execute as the first command of the task.
	//
	// Mutually exclusive with `execArgs`.
	// Default: - add steps using `task.exec(command)` or `task.spawn(subtask)`
	//
	// Experimental.
	Exec *string `field:"optional" json:"exec" yaml:"exec"`
	// Shell command to execute as the first command of the task, provided as a list of the program followed by its arguments (an "argv").
	//
	// A convenient alternative to `exec`: arguments with spaces or special
	// characters are passed through as-is, with no quoting needed. The elements
	// are not run through a shell, so environment variables (`$FOO`) are not
	// expanded and other shell features are unavailable.
	// See: TaskStep.execArgs *
	// Mutually exclusive with `exec`.
	//
	// Default: - add steps using `task.execArgs(args)`, `task.exec(command)` or `task.spawn(subtask)`
	//
	// Experimental.
	ExecArgs *[]*string `field:"optional" json:"execArgs" yaml:"execArgs"`
	// Should the provided `exec` shell command receive args passed to the task.
	// See: {@link TaskStepOptions.receiveArgs }
	//
	// Default: false.
	//
	// Experimental.
	ReceiveArgs *bool `field:"optional" json:"receiveArgs" yaml:"receiveArgs"`
	// List of task steps to run.
	// Experimental.
	Steps *[]*TaskStep `field:"optional" json:"steps" yaml:"steps"`
}

