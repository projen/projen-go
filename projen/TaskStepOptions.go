package projen


// Options for task steps.
// Experimental.
type TaskStepOptions struct {
	// A list of fixed arguments always passed to the step.
	//
	// Useful to re-use existing tasks without having to re-define the whole task.\
	// Fixed args are always passed to the step, even if `receiveArgs` is `false`
	// and are always passed before any args the task is called with.
	//
	// If the step executes a shell commands, args are passed through at the end of the `exec` shell command.\
	// The position of the args can be changed by including the marker `$@` inside the command string.
	//
	// If the step spawns a subtask, args are passed to the subtask.
	// The subtask must define steps receiving args for this to have any effect.
	//
	// If the step calls a builtin script, args are passed to the script.
	// It is up to the script to use or discard the arguments.
	//
	// Example:
	//   task.spawn("deploy", { args: ["--force"] });
	//
	// Default: - no arguments are passed to the step.
	//
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
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
	//
	// Values in this map can be simple, literal values or shell expressions that will be evaluated at runtime e.g. `$(echo "foo")`.
	//
	// Example:
	//   { "foo": "bar", "boo": "$(echo baz)" }
	//
	// Default: - no environment variables defined in step.
	//
	// Experimental.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Step name.
	// Default: - no name.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Capture this step's (trimmed) stdout into an environment variable of this name, visible to all later steps of the task run.
	//
	// For `spawn` steps the
	// spawned subtask's combined stdout is captured.
	//
	// Set only when the step runs (a skipped step leaves it unset) and always
	// overwrites. The step's output still streams live.
	// Default: - stdout is not captured.
	//
	// Experimental.
	OutputEnv *string `field:"optional" json:"outputEnv" yaml:"outputEnv"`
	// Should this step receive args passed to the task.
	//
	// If `true`, args are passed through at the end of the `exec` shell command.\
	// The position of the args can be changed by including the marker `$@` inside the command string.
	//
	// If the marker is explicitly double-quoted ("$@") arguments will be wrapped in double quotes, approximating
	// the whitespace preserving behavior of bash variable expansion.
	//
	// If the step spawns a subtask, args are passed to the subtask.
	// The subtask must define steps receiving args for this to have any effect.
	//
	// Example:
	//   task.exec("echo Hello $@ World!", { receiveArgs: true });
	//
	// Default: false.
	//
	// Experimental.
	ReceiveArgs *bool `field:"optional" json:"receiveArgs" yaml:"receiveArgs"`
	// The shell used to run this step, overriding the task/project shell.
	// See: {@link TaskCommonOptions.shell }
	//
	// Default: - the task's (or project's) shell.
	//
	// Experimental.
	Shell TaskShell `field:"optional" json:"shell" yaml:"shell"`
}

