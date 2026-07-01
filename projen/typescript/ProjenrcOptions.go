package typescript


// Experimental.
type ProjenrcOptions struct {
	// The name of the projenrc file.
	// Default: ".projenrc.ts"
	//
	// Experimental.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// A directory tree that may contain *.ts files that can be referenced from your projenrc typescript file.
	// Default: "projenrc".
	//
	// Experimental.
	ProjenCodeDir *string `field:"optional" json:"projenCodeDir" yaml:"projenCodeDir"`
	// The runner to use for executing TypeScript files.
	// Default: - the project's runner.
	//
	// Experimental.
	Runner TypeScriptRunner `field:"optional" json:"runner" yaml:"runner"`
	// Whether to use `SWC` for ts-node.
	// Default: false.
	//
	// Deprecated: Use `runner: TypeScriptRunner.tsNode({ swc: true })` instead.
	Swc *bool `field:"optional" json:"swc" yaml:"swc"`
}

