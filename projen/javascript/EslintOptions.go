package javascript


// Experimental.
type EslintOptions struct {
	// Files or glob patterns or directories with source files to lint (e.g. [ "src" ]).
	// Experimental.
	Dirs *[]*string `field:"required" json:"dirs" yaml:"dirs"`
	// Enable import alias for module paths.
	// Default: undefined.
	//
	// Experimental.
	AliasExtensions *[]*string `field:"optional" json:"aliasExtensions" yaml:"aliasExtensions"`
	// Enable import alias for module paths.
	// Default: undefined.
	//
	// Experimental.
	AliasMap *map[string]*string `field:"optional" json:"aliasMap" yaml:"aliasMap"`
	// Files or glob patterns or directories with source files that include tests and build tools.
	//
	// These sources are linted but may also import packages from `devDependencies`.
	// Default: [].
	//
	// Experimental.
	Devdirs *[]*string `field:"optional" json:"devdirs" yaml:"devdirs"`
	// File types that should be linted (e.g. [ ".js", ".ts" ]).
	// Default: [".ts"]
	//
	// Experimental.
	FileExtensions *[]*string `field:"optional" json:"fileExtensions" yaml:"fileExtensions"`
	// List of file patterns that should not be linted, using the same syntax as .gitignore patterns.
	// Default: [ '*.js', '*.d.ts', 'node_modules/', '*.generated.ts', 'coverage' ]
	//
	// Experimental.
	IgnorePatterns *[]*string `field:"optional" json:"ignorePatterns" yaml:"ignorePatterns"`
	// Should we lint .projenrc.js.
	// Default: true.
	//
	// Deprecated: set to `false` to remove any automatic rules and add manually.
	LintProjenRc *bool `field:"optional" json:"lintProjenRc" yaml:"lintProjenRc"`
	// Projenrc file to lint.
	//
	// Use empty string to disable.
	// Default: PROJEN_RC.
	//
	// Deprecated: provide as `devdirs`.
	LintProjenRcFile *string `field:"optional" json:"lintProjenRcFile" yaml:"lintProjenRcFile"`
	// Enable prettier for code formatting.
	// Default: false.
	//
	// Experimental.
	Prettier *bool `field:"optional" json:"prettier" yaml:"prettier"`
	// Always try to resolve types under `<root>@types` directory even it doesn't contain any source code.
	//
	// This prevents `import/no-unresolved` eslint errors when importing a `@types/*` module that would otherwise remain unresolved.
	// Default: true.
	//
	// Experimental.
	TsAlwaysTryTypes *bool `field:"optional" json:"tsAlwaysTryTypes" yaml:"tsAlwaysTryTypes"`
	// Path to `tsconfig.json` which should be used by eslint.
	// Default: "./tsconfig.json"
	//
	// Experimental.
	TsconfigPath *string `field:"optional" json:"tsconfigPath" yaml:"tsconfigPath"`
	// Write eslint configuration as YAML instead of JSON.
	// Default: false.
	//
	// Experimental.
	Yaml *bool `field:"optional" json:"yaml" yaml:"yaml"`
}

