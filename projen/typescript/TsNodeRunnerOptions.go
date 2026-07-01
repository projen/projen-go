package typescript


// Options for the ts-node runner.
// Experimental.
type TsNodeRunnerOptions struct {
	// Whether to use SWC for transpilation.
	//
	// This will disable type-checking.
	// See: https://github.com/TypeStrong/ts-node#swc
	//
	// Default: false.
	//
	// Experimental.
	Swc *bool `field:"optional" json:"swc" yaml:"swc"`
	// Path to the tsconfig file to use.
	// Default: - ts-node default discovery.
	//
	// Experimental.
	Tsconfig *string `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Whether to type-check the script during executing.
	// See: https://github.com/TypeStrong/ts-node#typecheck
	//
	// Default: true.
	//
	// Experimental.
	TypeCheck *bool `field:"optional" json:"typeCheck" yaml:"typeCheck"`
}

