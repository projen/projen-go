package typescript


// Options for the tsx runner.
// Experimental.
type TsxRunnerOptions struct {
	// Path to the tsconfig file to use.
	//
	// When specified, will use this tsconfig for running tsx and type-checking (if enabled).
	// Default: - tsx/typescript default discovery.
	//
	// Experimental.
	Tsconfig *string `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Whether to type-check the entrypoint before executing.
	//
	// Because tsx does not type check code, you may want to enable this for additional type safety.
	// When enabled, runs `tsc --noEmit`, using the provided tsconfig.
	// Default: false.
	//
	// Experimental.
	TypeCheck *bool `field:"optional" json:"typeCheck" yaml:"typeCheck"`
}

