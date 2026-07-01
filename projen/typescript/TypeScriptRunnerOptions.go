package typescript


// The options that can be adjusted on any {@link TypeScriptRunner}, regardless of its runtime.
//
// This is the override type for {@link TypeScriptRunner.copy}. Runtime-specific
// creation options (such as `swc` for ts-node) are set via the static factories
// and are preserved - but cannot be changed - by `copy`, so a copy can never
// change the runtime of a runner.
// Experimental.
type TypeScriptRunnerOptions struct {
	// Path to the tsconfig file to use.
	// Experimental.
	Tsconfig *string `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Whether to type-check the entrypoint.
	// Experimental.
	TypeCheck *bool `field:"optional" json:"typeCheck" yaml:"typeCheck"`
}

