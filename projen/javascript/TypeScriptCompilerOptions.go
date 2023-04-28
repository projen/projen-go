package javascript


// Experimental.
type TypeScriptCompilerOptions struct {
	// Suppress arbitrary extension import errors with the assumption that a bundler will be handling it.
	// See: https://www.typescriptlang.org/tsconfig#allowArbitraryExtensions
	//
	// Experimental.
	AllowArbitraryExtensions *bool `field:"optional" json:"allowArbitraryExtensions" yaml:"allowArbitraryExtensions"`
	// Allows TypeScript files to import each other with TypeScript-specific extensions (`.ts`, `.mts`, `.tsx`). Requires `noEmit` or `emitDeclarationOnly`.
	// Experimental.
	AllowImportingTsExtensions *bool `field:"optional" json:"allowImportingTsExtensions" yaml:"allowImportingTsExtensions"`
	// Allow JavaScript files to be compiled.
	// Experimental.
	AllowJs *bool `field:"optional" json:"allowJs" yaml:"allowJs"`
	// Allow default imports from modules with no default export.
	//
	// This does not affect code emit, just typechecking.
	// Experimental.
	AllowSyntheticDefaultImports *bool `field:"optional" json:"allowSyntheticDefaultImports" yaml:"allowSyntheticDefaultImports"`
	// Ensures that your files are parsed in the ECMAScript strict mode, and emit “use strict” for each source file.
	// Experimental.
	AlwaysStrict *bool `field:"optional" json:"alwaysStrict" yaml:"alwaysStrict"`
	// Lets you set a base directory to resolve non-absolute module names.
	//
	// You can define a root folder where you can do absolute file resolution.
	// Experimental.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// List of additional conditions that should succeed when TypeScript resolves from an `exports` or `imports` field of a `package.json`.
	// See: https://www.typescriptlang.org/tsconfig#customConditions
	//
	// Experimental.
	CustomConditions *[]*string `field:"optional" json:"customConditions" yaml:"customConditions"`
	// To be specified along with the above.
	// Experimental.
	Declaration *bool `field:"optional" json:"declaration" yaml:"declaration"`
	// Offers a way to configure the root directory for where declaration files are emitted.
	// Experimental.
	DeclarationDir *string `field:"optional" json:"declarationDir" yaml:"declarationDir"`
	// Only emit .d.ts files; do not emit .js files.
	// Experimental.
	EmitDeclarationOnly *bool `field:"optional" json:"emitDeclarationOnly" yaml:"emitDeclarationOnly"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	//
	// Decorators are a language feature which hasn’t yet been fully ratified into the JavaScript specification.
	// This means that the implementation version in TypeScript may differ from the implementation in JavaScript when it it decided by TC39.
	// You can find out more about decorator support in TypeScript in the handbook.
	// See: https://www.typescriptlang.org/docs/handbook/decorators.html
	//
	// Experimental.
	EmitDecoratorMetadata *bool `field:"optional" json:"emitDecoratorMetadata" yaml:"emitDecoratorMetadata"`
	// Emit __importStar and __importDefault helpers for runtime babel ecosystem compatibility and enable --allowSyntheticDefaultImports for typesystem compatibility.
	// Experimental.
	EsModuleInterop *bool `field:"optional" json:"esModuleInterop" yaml:"esModuleInterop"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	// Experimental.
	ExperimentalDecorators *bool `field:"optional" json:"experimentalDecorators" yaml:"experimentalDecorators"`
	// Disallow inconsistently-cased references to the same file.
	// Experimental.
	ForceConsistentCasingInFileNames *bool `field:"optional" json:"forceConsistentCasingInFileNames" yaml:"forceConsistentCasingInFileNames"`
	// This flag works because you can use `import type` to explicitly create an `import` statement which should never be emitted into JavaScript.
	// See: https://www.typescriptlang.org/tsconfig#importsNotUsedAsValues
	//
	// Experimental.
	ImportsNotUsedAsValues TypeScriptImportsNotUsedAsValues `field:"optional" json:"importsNotUsedAsValues" yaml:"importsNotUsedAsValues"`
	// When set, instead of writing out a .js.map file to provide source maps, TypeScript will embed the source map content in the .js files.
	// Experimental.
	InlineSourceMap *bool `field:"optional" json:"inlineSourceMap" yaml:"inlineSourceMap"`
	// When set, TypeScript will include the original content of the .ts file as an embedded string in the source map. This is often useful in the same cases as inlineSourceMap.
	// Experimental.
	InlineSources *bool `field:"optional" json:"inlineSources" yaml:"inlineSources"`
	// Perform additional checks to ensure that separate compilation (such as with transpileModule or @babel/plugin-transform-typescript) would be safe.
	// Experimental.
	IsolatedModules *bool `field:"optional" json:"isolatedModules" yaml:"isolatedModules"`
	// Support JSX in .tsx files: "react", "preserve", "react-native" etc.
	// Experimental.
	Jsx TypeScriptJsxMode `field:"optional" json:"jsx" yaml:"jsx"`
	// Declares the module specifier to be used for importing the jsx and jsxs factory functions when using jsx.
	// Experimental.
	JsxImportSource *string `field:"optional" json:"jsxImportSource" yaml:"jsxImportSource"`
	// Reference for type definitions / libraries to use (eg.
	//
	// ES2016, ES5, ES2018).
	// Experimental.
	Lib *[]*string `field:"optional" json:"lib" yaml:"lib"`
	// Sets the module system for the program.
	//
	// See https://www.typescriptlang.org/docs/handbook/modules.html#ambient-modules.
	// Experimental.
	Module *string `field:"optional" json:"module" yaml:"module"`
	// Determine how modules get resolved.
	//
	// Either "Node" for Node.js/io.js style resolution, or "Classic".
	// Experimental.
	ModuleResolution TypeScriptModuleResolution `field:"optional" json:"moduleResolution" yaml:"moduleResolution"`
	// Do not emit outputs.
	// Experimental.
	NoEmit *bool `field:"optional" json:"noEmit" yaml:"noEmit"`
	// Do not emit compiler output files like JavaScript source code, source-maps or declarations if any errors were reported.
	// Experimental.
	NoEmitOnError *bool `field:"optional" json:"noEmitOnError" yaml:"noEmitOnError"`
	// Report errors for fallthrough cases in switch statements.
	//
	// Ensures that any non-empty
	// case inside a switch statement includes either break or return. This means you won’t
	// accidentally ship a case fallthrough bug.
	// Experimental.
	NoFallthroughCasesInSwitch *bool `field:"optional" json:"noFallthroughCasesInSwitch" yaml:"noFallthroughCasesInSwitch"`
	// In some cases where no type annotations are present, TypeScript will fall back to a type of any for a variable when it cannot infer the type.
	// Experimental.
	NoImplicitAny *bool `field:"optional" json:"noImplicitAny" yaml:"noImplicitAny"`
	// Using `noImplicitOverride`, you can ensure that sub-classes never go out of sync as they are required to explicitly declare that they are overriding a member using the `override` keyword.
	//
	// This also improves readability of the programmer's intent.
	//
	// Available with TypeScript 4.3 and newer.
	// Experimental.
	NoImplicitOverride *bool `field:"optional" json:"noImplicitOverride" yaml:"noImplicitOverride"`
	// When enabled, TypeScript will check all code paths in a function to ensure they return a value.
	// Experimental.
	NoImplicitReturns *bool `field:"optional" json:"noImplicitReturns" yaml:"noImplicitReturns"`
	// Raise error on ‘this’ expressions with an implied ‘any’ type.
	// Experimental.
	NoImplicitThis *bool `field:"optional" json:"noImplicitThis" yaml:"noImplicitThis"`
	// Raise error on use of the dot syntax to access fields which are not defined.
	// Experimental.
	NoPropertyAccessFromIndexSignature *bool `field:"optional" json:"noPropertyAccessFromIndexSignature" yaml:"noPropertyAccessFromIndexSignature"`
	// Raise error when accessing indexes on objects with unknown keys defined in index signatures.
	// Experimental.
	NoUncheckedIndexedAccess *bool `field:"optional" json:"noUncheckedIndexedAccess" yaml:"noUncheckedIndexedAccess"`
	// Report errors on unused local variables.
	// Experimental.
	NoUnusedLocals *bool `field:"optional" json:"noUnusedLocals" yaml:"noUnusedLocals"`
	// Report errors on unused parameters in functions.
	// Experimental.
	NoUnusedParameters *bool `field:"optional" json:"noUnusedParameters" yaml:"noUnusedParameters"`
	// Output directory for the compiled files.
	// Experimental.
	OutDir *string `field:"optional" json:"outDir" yaml:"outDir"`
	// A series of entries which re-map imports to lookup locations relative to the baseUrl, there is a larger coverage of paths in the handbook.
	//
	// paths lets you declare how TypeScript should resolve an import in your require/imports.
	// Experimental.
	Paths *map[string]*[]*string `field:"optional" json:"paths" yaml:"paths"`
	// Allows importing modules with a ‘.json’ extension, which is a common practice in node projects. This includes generating a type for the import based on the static JSON shape.
	// Experimental.
	ResolveJsonModule *bool `field:"optional" json:"resolveJsonModule" yaml:"resolveJsonModule"`
	// Forces TypeScript to consult the `exports` field of `package.json` files if it ever reads from a package in `node_modules`.
	// Experimental.
	ResolvePackageJsonExports *bool `field:"optional" json:"resolvePackageJsonExports" yaml:"resolvePackageJsonExports"`
	// Forces TypeScript to consult the `imports` field of `package.json` when performing a lookup that begins with `#` from a file that has a `package.json` as an ancestor.
	// Experimental.
	ResolvePackageJsonImports *bool `field:"optional" json:"resolvePackageJsonImports" yaml:"resolvePackageJsonImports"`
	// Specifies the root directory of input files.
	//
	// Only use to control the output directory structure with `outDir`.
	// Experimental.
	RootDir *string `field:"optional" json:"rootDir" yaml:"rootDir"`
	// Skip type checking of all declaration files (*.d.ts).
	// Experimental.
	SkipLibCheck *bool `field:"optional" json:"skipLibCheck" yaml:"skipLibCheck"`
	// Enables the generation of sourcemap files.
	// Experimental.
	SourceMap *bool `field:"optional" json:"sourceMap" yaml:"sourceMap"`
	// Specify the location where a debugger should locate TypeScript files instead of relative source locations.
	// Experimental.
	SourceRoot *string `field:"optional" json:"sourceRoot" yaml:"sourceRoot"`
	// The strict flag enables a wide range of type checking behavior that results in stronger guarantees of program correctness.
	//
	// Turning this on is equivalent to enabling all of the strict mode family
	// options, which are outlined below. You can then turn off individual strict mode family checks as
	// needed.
	// Experimental.
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	// When strictNullChecks is false, null and undefined are effectively ignored by the language.
	//
	// This can lead to unexpected errors at runtime.
	// When strictNullChecks is true, null and undefined have their own distinct types and you’ll
	// get a type error if you try to use them where a concrete value is expected.
	// Experimental.
	StrictNullChecks *bool `field:"optional" json:"strictNullChecks" yaml:"strictNullChecks"`
	// When set to true, TypeScript will raise an error when a class property was declared but not set in the constructor.
	// Experimental.
	StrictPropertyInitialization *bool `field:"optional" json:"strictPropertyInitialization" yaml:"strictPropertyInitialization"`
	// Do not emit declarations for code that has an @internal annotation in it’s JSDoc comment.
	// Experimental.
	StripInternal *bool `field:"optional" json:"stripInternal" yaml:"stripInternal"`
	// Modern browsers support all ES6 features, so ES6 is a good choice.
	//
	// You might choose to set
	// a lower target if your code is deployed to older environments, or a higher target if your
	// code is guaranteed to run in newer environments.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Simplifies TypeScript's handling of import/export `type` modifiers.
	// See: https://www.typescriptlang.org/tsconfig#verbatimModuleSyntax
	//
	// Experimental.
	VerbatimModuleSyntax *bool `field:"optional" json:"verbatimModuleSyntax" yaml:"verbatimModuleSyntax"`
}

