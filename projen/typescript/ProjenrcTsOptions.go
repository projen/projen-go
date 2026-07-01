package typescript


// Experimental.
type ProjenrcTsOptions struct {
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
	// Default: TypeScriptRunner.tsNode()
	//
	// Experimental.
	Runner TypeScriptRunner `field:"optional" json:"runner" yaml:"runner"`
	// The name of the tsconfig file that will be used by the runner when compiling projen source files.
	// Default: "tsconfig.projen.json"
	//
	// Deprecated: Use `runner` to configure the tsconfigFileName directly.
	TsconfigFileName *string `field:"optional" json:"tsconfigFileName" yaml:"tsconfigFileName"`
}

