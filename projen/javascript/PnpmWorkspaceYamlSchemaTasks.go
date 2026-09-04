package javascript


// Experimental.
type PnpmWorkspaceYamlSchemaTasks struct {
	// A positive integer limiting how many instances of this named task may run across workspace projects at once.
	//
	// This limit is separate from workspaceConcurrency.
	// Experimental.
	Concurrency *float64 `field:"optional" json:"concurrency" yaml:"concurrency"`
	// Tasks this task depends on.
	//
	// Each entry is either the task in the same project (e.g. `build`) or the task in each selected workspace dependency of the project (e.g. `^build`).
	// Experimental.
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
}

