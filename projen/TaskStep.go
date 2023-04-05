// CDK for software projects
package projen


// A single step within a task.
//
// The step could either be  the execution of a
// shell command or execution of a sub-task, by name.
// Experimental.
type TaskStep struct {
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
	// The name of a built-in task to execute.
	//
	// Built-in tasks are node.js programs baked into the projen module and as
	// component runtime helpers.
	//
	// The name is a path relative to the projen lib/ directory (without the .task.js extension).
	// For example, if your built in builtin task is under `src/release/resolve-version.task.ts`,
	// then this would be `release/resolve-version`.
	// Experimental.
	Builtin *string `field:"optional" json:"builtin" yaml:"builtin"`
	// Shell command to execute.
	// Experimental.
	Exec *string `field:"optional" json:"exec" yaml:"exec"`
	// Print a message.
	// Experimental.
	Say *string `field:"optional" json:"say" yaml:"say"`
	// Subtask to execute.
	// Experimental.
	Spawn *string `field:"optional" json:"spawn" yaml:"spawn"`
}

