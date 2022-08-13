package javascript


// Experimental.
type EslintOptions struct {
	// Directories with source files to lint (e.g. [ "src" ]).
	// Experimental.
	Dirs *[]*string `field:"required" json:"dirs" yaml:"dirs"`
	// Enable import alias for module paths.
	// Experimental.
	AliasExtensions *[]*string `field:"optional" json:"aliasExtensions" yaml:"aliasExtensions"`
	// Enable import alias for module paths.
	// Experimental.
	AliasMap *map[string]*string `field:"optional" json:"aliasMap" yaml:"aliasMap"`
	// Directories with source files that include tests and build tools.
	//
	// These
	// sources are linted but may also import packages from `devDependencies`.
	// Experimental.
	Devdirs *[]*string `field:"optional" json:"devdirs" yaml:"devdirs"`
	// File types that should be linted (e.g. [ ".js", ".ts" ]).
	// Experimental.
	FileExtensions *[]*string `field:"optional" json:"fileExtensions" yaml:"fileExtensions"`
	// List of file patterns that should not be linted, using the same syntax as .gitignore patterns.
	// Experimental.
	IgnorePatterns *[]*string `field:"optional" json:"ignorePatterns" yaml:"ignorePatterns"`
	// Should we lint .projenrc.js.
	// Deprecated: use lintProjenRcFile instead.
	LintProjenRc *bool `field:"optional" json:"lintProjenRc" yaml:"lintProjenRc"`
	// Projenrc file to lint.
	//
	// Use empty string to disable.
	// Experimental.
	LintProjenRcFile *string `field:"optional" json:"lintProjenRcFile" yaml:"lintProjenRcFile"`
	// Enable prettier for code formatting.
	// Experimental.
	Prettier *bool `field:"optional" json:"prettier" yaml:"prettier"`
	// Always try to resolve types under `<root>@types` directory even it doesn't contain any source code.
	//
	// This prevents `import/no-unresolved` eslint errors when importing a `@types/*` module that would otherwise remain unresolved.
	// Experimental.
	TsAlwaysTryTypes *bool `field:"optional" json:"tsAlwaysTryTypes" yaml:"tsAlwaysTryTypes"`
	// Path to `tsconfig.json` which should be used by eslint.
	// Experimental.
	TsconfigPath *string `field:"optional" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Write eslint configuration as YAML instead of JSON.
	// Experimental.
	Yaml *bool `field:"optional" json:"yaml" yaml:"yaml"`
}

