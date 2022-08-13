// CDK for software projects
package projen


// A single step within a task.
//
// The step could either be  the execution of a
// shell command or execution of a sub-task, by name.
// Experimental.
type TaskStep struct {
	// The working directory for this step.
	// Experimental.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// Step name.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
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

