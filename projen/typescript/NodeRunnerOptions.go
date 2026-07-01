package typescript


// Options for the native Node.js TypeScript runner.
// Experimental.
type NodeRunnerOptions struct {
	// Whether to also enable `--experimental-transform-types`.
	// Default: false.
	//
	// Experimental.
	ExperimentalTransformTypes *bool `field:"optional" json:"experimentalTransformTypes" yaml:"experimentalTransformTypes"`
	// Path to the tsconfig file for type-checking.
	//
	// When specified, will use this tsconfig for type-checking (if enabled).
	// Default: - typescript default discovery.
	//
	// Experimental.
	Tsconfig *string `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Whether to type-check the entrypoint before executing.
	//
	// Because the native Node.js TypeScript does not type check code,
	// you may want to enable this for additional type safety.
	// When enabled, runs `tsc --noEmit`, using the provided tsconfig.
	// Default: false.
	//
	// Experimental.
	TypeCheck *bool `field:"optional" json:"typeCheck" yaml:"typeCheck"`
}

