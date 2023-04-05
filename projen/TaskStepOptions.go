// CDK for software projects
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
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The working directory for this step.
	// Experimental.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// Step name.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Should this step receive args passed to the task.
	//
	// If `true`, args are passed through at the end of the `exec` shell command.\
	// The position of the args can be changed by including the marker `$@` inside the command string.
	//
	// If the step spawns a subtask, args are passed to the subtask.
	// The subtask must define steps receiving args for this to have any effect.
	//
	// Example:
	//   task.exec("echo Hello $@ World!", { receiveArgs: true });
	//
	// Experimental.
	ReceiveArgs *bool `field:"optional" json:"receiveArgs" yaml:"receiveArgs"`
}

