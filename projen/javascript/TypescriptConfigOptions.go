package javascript


// Experimental.
type TypescriptConfigOptions struct {
	// Compiler options to use.
	// Experimental.
	CompilerOptions *TypeScriptCompilerOptions `field:"required" json:"compilerOptions" yaml:"compilerOptions"`
	// Filters results from the "include" option.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Base `tsconfig.json` configuration(s) to inherit from.
	// Experimental.
	Extends TypescriptConfigExtends `field:"optional" json:"extends" yaml:"extends"`
	// Experimental.
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Specifies a list of glob patterns that match TypeScript files to be included in compilation.
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

