package typescript


// See: https://kulshekhar.github.io/ts-jest/docs/getting-started/options
//
// Experimental.
type TsJestTransformOptions struct {
	// Custom TypeScript AST transformers.
	// Default: auto.
	//
	// Experimental.
	AstTransformers *map[string]interface{} `field:"optional" json:"astTransformers" yaml:"astTransformers"`
	// Babel(Jest) related configuration.
	// Default: TsJestBabelConfig.disabled()
	//
	// Experimental.
	BabelConfig TsJestBabelConfig `field:"optional" json:"babelConfig" yaml:"babelConfig"`
	// TypeScript module to use as compiler.
	// Default: "typescript".
	//
	// Experimental.
	Compiler *string `field:"optional" json:"compiler" yaml:"compiler"`
	// Diagnostics related configuration.
	// Default: TsJestDiagnostics.all()
	//
	// Experimental.
	Diagnostics TsJestDiagnostics `field:"optional" json:"diagnostics" yaml:"diagnostics"`
	// Run ts-jest tests with this TSConfig isolatedModules setting.
	//
	// You'll lose type-checking ability and some features such as const enum, but in the case you plan on using Jest with the cache disabled (jest --no-cache), your tests will then run much faster.
	// See: https://kulshekhar.github.io/ts-jest/docs/getting-started/options/isolatedModules
	//
	// Default: false.
	//
	// Experimental.
	IsolatedModules *bool `field:"optional" json:"isolatedModules" yaml:"isolatedModules"`
	// Files which will become modules returning self content.
	// Default: disabled.
	//
	// Experimental.
	StringifyContentPathRegex *string `field:"optional" json:"stringifyContentPathRegex" yaml:"stringifyContentPathRegex"`
	// TypeScript compiler related configuration.
	// Default: - Your project's `tsconfigDev` file.
	//
	// Experimental.
	Tsconfig TsJestTsconfig `field:"optional" json:"tsconfig" yaml:"tsconfig"`
	// Enable ESM support.
	// Default: auto.
	//
	// Experimental.
	UseEsm *bool `field:"optional" json:"useEsm" yaml:"useEsm"`
}

