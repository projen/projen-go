package typescript


// Options for the native Node.js TypeScript runner.
// Experimental.
type NodeRunnerOptions struct {
	// Whether to also enable `--experimental-transform-types`.
	// Default: false.
	//
	// Deprecated: This flag has been removed from Node.js 26. Use `transformTypes` instead.
	ExperimentalTransformTypes *bool `field:"optional" json:"experimentalTransformTypes" yaml:"experimentalTransformTypes"`
	// Whether to enable transformation of TypeScript-only syntax (e.g. enums, namespaces).
	//
	// Uses `amaro` (the TypeScript transformer used internally by Node.js) as an
	// external loader via `--import=amaro/transform`. Adds a dependency on the
	// `amaro` package and enables `--enable-source-maps` to preserve accurate
	// stack traces.
	// See: https://github.com/nodejs/amaro
	//
	// Default: false.
	//
	// Experimental.
	TransformTypes *bool `field:"optional" json:"transformTypes" yaml:"transformTypes"`
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

