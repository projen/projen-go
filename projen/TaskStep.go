package projen


// A single step within a task.
//
// The step could either be  the execution of a
// shell command or execution of a sub-task, by name.
//
// The `tasks.json` (manifest) form of a step. {@link TaskStepOptions} is the
// form used to define steps (via `task.exec()` etc.); they differ only in the
// rendered `shell` field.
// Experimental.
type TaskStep struct {
	// A list of fixed arguments always passed to the step.
	// See: {@link TaskStepOptions.args }
	//
	// Default: - no arguments are passed to the step.
	//
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The name of a built-in task to execute.
	//
	// Built-in tasks are node.js programs baked into the projen module and as
	// component runtime helpers.
	//
	// The name is a path relative to the projen lib/ directory (without the .task.js extension).
	// For example, if your built in builtin task is under `src/release/resolve-version.task.ts`,
	// then this would be `release/resolve-version`.
	// Default: - do not execute a builtin task.
	//
	// Experimental.
	Builtin *string `field:"optional" json:"builtin" yaml:"builtin"`
	// A shell command which determines if the this step should be executed.
	//
	// If
	// the program exits with a zero exit code, the step will be executed. A non-zero
	// code means the step will be skipped (subsequent task steps will still be evaluated/executed).
	// Experimental.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The working directory for this step.
	// Default: - determined by the task.
	//
	// Experimental.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// Defines environment variables for the execution of this step (`exec` and `builtin` only).
	// See: {@link TaskStepOptions.env }
	//
	// Default: - no environment variables defined in step.
	//
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Shell command to execute.
	//
	// The whole command is a single shell string. To pass arguments as a list
	// instead - without having to quote spaces or other characters yourself -
	// use `execArgs`.
	// Default: - don't execute a shell command.
	//
	// Experimental.
	Exec *string `field:"optional" json:"exec" yaml:"exec"`
	// Shell command to execute, provided as a list of the program followed by its arguments (an "argv").
	//
	// Often more convenient than `exec`: each element is passed to the
	// program as-is, so arguments with spaces or special characters don't need
	// quoting. Fixed (`args`) or received (`receiveArgs`) arguments are inserted
	// wherever a `$@` element appears, or appended at the end if there is none.
	//
	// The elements are not run through a shell, so environment variables (`$FOO`)
	// are not expanded and other shell features are unavailable. Use `exec`
	// if you need them.
	//
	// Mutually exclusive with `exec`.
	//
	// Example:
	//   { execArgs: ["echo", "hello world"] }
	//
	// Default: - don't execute a shell command.
	//
	// Experimental.
	ExecArgs *[]*string `field:"optional" json:"execArgs" yaml:"execArgs"`
	// Step name.
	// Default: - no name.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Capture this step's standard output and expose it to later steps as an environment variable with this name.
	// See: {@link TaskStepOptions.outputEnv }
	//
	// Default: - stdout is not captured.
	//
	// Experimental.
	OutputEnv *string `field:"optional" json:"outputEnv" yaml:"outputEnv"`
	// Should this step receive args passed to the task.
	// See: {@link TaskStepOptions.receiveArgs }
	//
	// Default: false.
	//
	// Experimental.
	ReceiveArgs *bool `field:"optional" json:"receiveArgs" yaml:"receiveArgs"`
	// Print a message.
	// Default: - don't say anything.
	//
	// Experimental.
	Say *string `field:"optional" json:"say" yaml:"say"`
	// The step shell in `tasks.json` form: a keyword (`"projen"` or `"system"`) or an invocation argument list.
	// Default: - the task's (or project's) shell.
	//
	// Experimental.
	Shell interface{} `field:"optional" json:"shell" yaml:"shell"`
	// Subtask to execute.
	// Default: - don't spawn a subtask.
	//
	// Experimental.
	Spawn *string `field:"optional" json:"spawn" yaml:"spawn"`
}

