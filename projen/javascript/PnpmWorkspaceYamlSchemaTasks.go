package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaTasks struct {
	// Set to false to opt a task with declared outputs back out of the `pnpm pipeline` cache.
	//
	// Read by pnpm pipeline only.
	// Experimental.
	Cache *bool `field:"optional" json:"cache" yaml:"cache"`
	// Points Cargo's target and build directories at this path so a Cargo task keeps its local build state between runs and between git worktrees.
	//
	// The directory must be relative to the project and ignored by git. Read by pnpm pipeline only.
	// Experimental.
	CargoTargetDir *string `field:"optional" json:"cargoTargetDir" yaml:"cargoTargetDir"`
	// A positive integer limiting how many instances of this named task may run across workspace projects at once.
	//
	// This limit is separate from workspaceConcurrency.
	// Experimental.
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Assigns this task to a named concurrency group whose machine-wide limit is configured under concurrencyGroups.
	//
	// Added in pnpm v12.5.0.
	// Experimental.
	ConcurrencyGroup *string `field:"optional" json:"concurrencyGroup" yaml:"concurrencyGroup"`
	// Tasks this task depends on.
	//
	// Each entry is either the task in the same project (e.g. `build`) or the task in each selected workspace dependency of the project (e.g. `^build`).
	// Experimental.
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Environment variable names whose values take part in the task's cache key.
	//
	// The values are hashed, never stored. Read by pnpm pipeline only.
	// Experimental.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// Globs narrowing what goes into the task's cache key.
	//
	// Without it, every tracked file of the project counts. An entry prefixed with `+` adds to that default instead of replacing it. Read by pnpm pipeline only.
	// Experimental.
	Inputs *[]*string `field:"optional" json:"inputs" yaml:"inputs"`
	// Globs, relative to the project directory, naming the files the task produces.
	//
	// Their presence makes the task cacheable by `pnpm pipeline`; declaring `outputs: []` states the task produces no files. Read by pnpm pipeline only.
	// Experimental.
	Outputs *[]*string `field:"optional" json:"outputs" yaml:"outputs"`
	// Waiting tasks take available slots in order of descending priority, with higher values running first.
	//
	// Ties are broken by arrival order. Added in pnpm v12.6.0.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

