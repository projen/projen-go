package javascript


// Experimental.
type TypeScriptCompilerOptions struct {
	// Suppress arbitrary extension import errors with the assumption that a bundler will be handling it.
	// See: https://www.typescriptlang.org/tsconfig#allowArbitraryExtensions
	//
	// Default: undefined.
	//
	// Experimental.
	AllowArbitraryExtensions *bool `field:"optional" json:"allowArbitraryExtensions" yaml:"allowArbitraryExtensions"`
	// Allows TypeScript files to import each other with TypeScript-specific extensions (`.ts`, `.mts`, `.tsx`). Requires `noEmit` or `emitDeclarationOnly`.
	// Default: undefined.
	//
	// Experimental.
	AllowImportingTsExtensions *bool `field:"optional" json:"allowImportingTsExtensions" yaml:"allowImportingTsExtensions"`
	// Allow JavaScript files to be compiled.
	// Default: false.
	//
	// Experimental.
	AllowJs *bool `field:"optional" json:"allowJs" yaml:"allowJs"`
	// Allow default imports from modules with no default export.
	//
	// This does not affect code emit, just typechecking.
	// Experimental.
	AllowSyntheticDefaultImports *bool `field:"optional" json:"allowSyntheticDefaultImports" yaml:"allowSyntheticDefaultImports"`
	// Allow Unreachable Code.
	//
	// When:
	//
	// - `undefined` (default) provide suggestions as warnings to editors
	// - `true` unreachable code is ignored
	// - `false` raises compiler errors about unreachable code
	//
	// These warnings are only about code which is provably unreachable due to the use of JavaScript syntax.
	// See: https://www.typescriptlang.org/tsconfig#allowUnreachableCode
	//
	// Experimental.
	AllowUnreachableCode *bool `field:"optional" json:"allowUnreachableCode" yaml:"allowUnreachableCode"`
	// Allow Unused Labels.
	//
	// When:
	//
	// - `undefined` (default) provide suggestions as warnings to editors
	// - `true` unused labels are ignored
	// - `false` raises compiler errors about unused labels
	//
	// Labels are very rare in JavaScript and typically indicate an attempt to write an object literal:
	//
	// ```ts
	// function verifyAge(age: number) {
	//   // Forgot 'return' statement
	//   if (age > 18) {
	//     verified: true;
	// //  ^^^^^^^^ Unused label.
	//   }
	// }
	// ```.
	// See: https://www.typescriptlang.org/tsconfig#allowUnusedLabels
	//
	// Experimental.
	AllowUnusedLabels *bool `field:"optional" json:"allowUnusedLabels" yaml:"allowUnusedLabels"`
	// Ensures that your files are parsed in the ECMAScript strict mode, and emit “use strict” for each source file.
	// Default: true.
	//
	// Experimental.
	AlwaysStrict *bool `field:"optional" json:"alwaysStrict" yaml:"alwaysStrict"`
	// Lets you set a base directory to resolve non-absolute module names.
	//
	// You can define a root folder where you can do absolute file resolution.
	// Experimental.
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Check JS.
	//
	// Works in tandem with [allowJs](https://www.typescriptlang.org/tsconfig#allowJs). When checkJs is enabled then
	// errors are reported in JavaScript files. This is the equivalent of including //
	// See: https://www.typescriptlang.org/tsconfig#checkJs
	//
	// Experimental.
	CheckJs *bool `field:"optional" json:"checkJs" yaml:"checkJs"`
	// List of additional conditions that should succeed when TypeScript resolves from an `exports` or `imports` field of a `package.json`.
	// See: https://www.typescriptlang.org/tsconfig#customConditions
	//
	// Default: undefined.
	//
	// Experimental.
	CustomConditions *[]*string `field:"optional" json:"customConditions" yaml:"customConditions"`
	// To be specified along with the above.
	// Experimental.
	Declaration *bool `field:"optional" json:"declaration" yaml:"declaration"`
	// Offers a way to configure the root directory for where declaration files are emitted.
	// Experimental.
	DeclarationDir *string `field:"optional" json:"declarationDir" yaml:"declarationDir"`
	// Generates a source map for .d.ts files which map back to the original .ts source file. This will allow editors such as VS Code to go to the original .ts file when using features like Go to Definition.
	// See: {@link https://www.typescriptlang.org/tsconfig#declarationMap}
	//
	// Experimental.
	DeclarationMap *bool `field:"optional" json:"declarationMap" yaml:"declarationMap"`
	// Downleveling is TypeScript’s term for transpiling to an older version of JavaScript.
	//
	// This flag is to enable support for a more accurate implementation of how modern JavaScript iterates through new concepts in older JavaScript runtimes.
	//
	// ECMAScript 6 added several new iteration primitives: the for / of loop (for (el of arr)), Array spread ([a, ...b]), argument spread (fn(...args)), and Symbol.iterator.
	// downlevelIteration allows for these iteration primitives to be used more accurately in ES5 environments if a Symbol.iterator implementation is present.
	// Experimental.
	DownlevelIteration *bool `field:"optional" json:"downlevelIteration" yaml:"downlevelIteration"`
	// Only emit .d.ts files; do not emit .js files.
	// Default: false.
	//
	// Experimental.
	EmitDeclarationOnly *bool `field:"optional" json:"emitDeclarationOnly" yaml:"emitDeclarationOnly"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	//
	// Decorators are a language feature which hasn’t yet been fully ratified into the JavaScript specification.
	// This means that the implementation version in TypeScript may differ from the implementation in JavaScript when it it decided by TC39.
	// You can find out more about decorator support in TypeScript in the handbook.
	// See: https://www.typescriptlang.org/docs/handbook/decorators.html
	//
	// Default: undefined.
	//
	// Experimental.
	EmitDecoratorMetadata *bool `field:"optional" json:"emitDecoratorMetadata" yaml:"emitDecoratorMetadata"`
	// Emit __importStar and __importDefault helpers for runtime babel ecosystem compatibility and enable --allowSyntheticDefaultImports for typesystem compatibility.
	// Default: false.
	//
	// Experimental.
	EsModuleInterop *bool `field:"optional" json:"esModuleInterop" yaml:"esModuleInterop"`
	// Specifies that optional property types should be interpreted exactly as written, meaning that `| undefined` is not added to the type Available with TypeScript 4.4 and newer.
	// Default: false.
	//
	// Experimental.
	ExactOptionalPropertyTypes *bool `field:"optional" json:"exactOptionalPropertyTypes" yaml:"exactOptionalPropertyTypes"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	// Default: true.
	//
	// Experimental.
	ExperimentalDecorators *bool `field:"optional" json:"experimentalDecorators" yaml:"experimentalDecorators"`
	// Disallow inconsistently-cased references to the same file.
	// Default: false.
	//
	// Experimental.
	ForceConsistentCasingInFileNames *bool `field:"optional" json:"forceConsistentCasingInFileNames" yaml:"forceConsistentCasingInFileNames"`
	// This flag works because you can use `import type` to explicitly create an `import` statement which should never be emitted into JavaScript.
	// See: https://www.typescriptlang.org/tsconfig#importsNotUsedAsValues
	//
	// Default: "remove".
	//
	// Experimental.
	ImportsNotUsedAsValues TypeScriptImportsNotUsedAsValues `field:"optional" json:"importsNotUsedAsValues" yaml:"importsNotUsedAsValues"`
	// Tells TypeScript to save information about the project graph from the last compilation to files stored on disk.
	//
	// This creates a series of .tsbuildinfo files in the same folder as your compilation output.
	// They are not used by your JavaScript at runtime and can be safely deleted.
	// You can read more about the flag in the 3.4 release notes.
	// See: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-4.html#faster-subsequent-builds-with-the---incremental-flag
	//
	// To control which folders you want to the files to be built to, use the config option tsBuildInfoFile.
	//
	// Experimental.
	Incremental *bool `field:"optional" json:"incremental" yaml:"incremental"`
	// When set, instead of writing out a .js.map file to provide source maps, TypeScript will embed the source map content in the .js files.
	// Default: true.
	//
	// Experimental.
	InlineSourceMap *bool `field:"optional" json:"inlineSourceMap" yaml:"inlineSourceMap"`
	// When set, TypeScript will include the original content of the .ts file as an embedded string in the source map. This is often useful in the same cases as inlineSourceMap.
	// Default: true.
	//
	// Experimental.
	InlineSources *bool `field:"optional" json:"inlineSources" yaml:"inlineSources"`
	// Perform additional checks to ensure that separate compilation (such as with transpileModule or.
	// Default: false.
	//
	// Experimental.
	IsolatedModules *bool `field:"optional" json:"isolatedModules" yaml:"isolatedModules"`
	// Support JSX in .tsx files: "react", "preserve", "react-native" etc.
	// Default: undefined.
	//
	// Experimental.
	Jsx TypeScriptJsxMode `field:"optional" json:"jsx" yaml:"jsx"`
	// Declares the module specifier to be used for importing the jsx and jsxs factory functions when using jsx.
	// Default: undefined.
	//
	// Experimental.
	JsxImportSource *string `field:"optional" json:"jsxImportSource" yaml:"jsxImportSource"`
	// Reference for type definitions / libraries to use (eg.
	//
	// ES2016, ES5, ES2018).
	// Default: [ "es2018" ].
	//
	// Experimental.
	Lib *[]*string `field:"optional" json:"lib" yaml:"lib"`
	// Sets the module system for the program.
	//
	// See https://www.typescriptlang.org/docs/handbook/modules.html#ambient-modules.
	// Default: "CommonJS".
	//
	// Experimental.
	Module *string `field:"optional" json:"module" yaml:"module"`
	// This setting controls how TypeScript determines whether a file is a [script or a module](https://www.typescriptlang.org/docs/handbook/modules/theory.html#scripts-and-modules-in-javascript).
	// Default: "auto".
	//
	// Experimental.
	ModuleDetection TypeScriptModuleDetection `field:"optional" json:"moduleDetection" yaml:"moduleDetection"`
	// Determine how modules get resolved.
	//
	// Either "Node" for Node.js/io.js style resolution, or "Classic".
	// Default: "node".
	//
	// Experimental.
	ModuleResolution TypeScriptModuleResolution `field:"optional" json:"moduleResolution" yaml:"moduleResolution"`
	// Do not emit outputs.
	// Default: false.
	//
	// Experimental.
	NoEmit *bool `field:"optional" json:"noEmit" yaml:"noEmit"`
	// Do not emit compiler output files like JavaScript source code, source-maps or declarations if any errors were reported.
	// Default: true.
	//
	// Experimental.
	NoEmitOnError *bool `field:"optional" json:"noEmitOnError" yaml:"noEmitOnError"`
	// Report errors for fallthrough cases in switch statements.
	//
	// Ensures that any non-empty
	// case inside a switch statement includes either break or return. This means you won’t
	// accidentally ship a case fallthrough bug.
	// Default: true.
	//
	// Experimental.
	NoFallthroughCasesInSwitch *bool `field:"optional" json:"noFallthroughCasesInSwitch" yaml:"noFallthroughCasesInSwitch"`
	// In some cases where no type annotations are present, TypeScript will fall back to a type of any for a variable when it cannot infer the type.
	// Default: true.
	//
	// Experimental.
	NoImplicitAny *bool `field:"optional" json:"noImplicitAny" yaml:"noImplicitAny"`
	// Using `noImplicitOverride`, you can ensure that sub-classes never go out of sync as they are required to explicitly declare that they are overriding a member using the `override` keyword.
	//
	// This also improves readability of the programmer's intent.
	//
	// Available with TypeScript 4.3 and newer.
	// Default: false.
	//
	// Experimental.
	NoImplicitOverride *bool `field:"optional" json:"noImplicitOverride" yaml:"noImplicitOverride"`
	// When enabled, TypeScript will check all code paths in a function to ensure they return a value.
	// Default: true.
	//
	// Experimental.
	NoImplicitReturns *bool `field:"optional" json:"noImplicitReturns" yaml:"noImplicitReturns"`
	// Raise error on ‘this’ expressions with an implied ‘any’ type.
	// Default: true.
	//
	// Experimental.
	NoImplicitThis *bool `field:"optional" json:"noImplicitThis" yaml:"noImplicitThis"`
	// Raise error on use of the dot syntax to access fields which are not defined.
	// Default: true.
	//
	// Experimental.
	NoPropertyAccessFromIndexSignature *bool `field:"optional" json:"noPropertyAccessFromIndexSignature" yaml:"noPropertyAccessFromIndexSignature"`
	// Raise error when accessing indexes on objects with unknown keys defined in index signatures.
	// Default: true.
	//
	// Experimental.
	NoUncheckedIndexedAccess *bool `field:"optional" json:"noUncheckedIndexedAccess" yaml:"noUncheckedIndexedAccess"`
	// Report errors on unused local variables.
	// Default: true.
	//
	// Experimental.
	NoUnusedLocals *bool `field:"optional" json:"noUnusedLocals" yaml:"noUnusedLocals"`
	// Report errors on unused parameters in functions.
	// Default: true.
	//
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
	// Default: true.
	//
	// Experimental.
	ResolveJsonModule *bool `field:"optional" json:"resolveJsonModule" yaml:"resolveJsonModule"`
	// Forces TypeScript to consult the `exports` field of `package.json` files if it ever reads from a package in `node_modules`.
	// Default: true.
	//
	// Experimental.
	ResolvePackageJsonExports *bool `field:"optional" json:"resolvePackageJsonExports" yaml:"resolvePackageJsonExports"`
	// Forces TypeScript to consult the `imports` field of `package.json` when performing a lookup that begins with `#` from a file that has a `package.json` as an ancestor.
	// Default: undefined.
	//
	// Experimental.
	ResolvePackageJsonImports *bool `field:"optional" json:"resolvePackageJsonImports" yaml:"resolvePackageJsonImports"`
	// Specifies the root directory of input files.
	//
	// Only use to control the output directory structure with `outDir`.
	// Experimental.
	RootDir *string `field:"optional" json:"rootDir" yaml:"rootDir"`
	// Skip type checking of all declaration files (*.d.ts).
	// Default: false.
	//
	// Experimental.
	SkipLibCheck *bool `field:"optional" json:"skipLibCheck" yaml:"skipLibCheck"`
	// Enables the generation of sourcemap files.
	// Default: undefined.
	//
	// Experimental.
	SourceMap *bool `field:"optional" json:"sourceMap" yaml:"sourceMap"`
	// Specify the location where a debugger should locate TypeScript files instead of relative source locations.
	// Default: undefined.
	//
	// Experimental.
	SourceRoot *string `field:"optional" json:"sourceRoot" yaml:"sourceRoot"`
	// The strict flag enables a wide range of type checking behavior that results in stronger guarantees of program correctness.
	//
	// Turning this on is equivalent to enabling all of the strict mode family
	// options, which are outlined below. You can then turn off individual strict mode family checks as
	// needed.
	// Default: true.
	//
	// Experimental.
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	// When strictNullChecks is false, null and undefined are effectively ignored by the language.
	//
	// This can lead to unexpected errors at runtime.
	// When strictNullChecks is true, null and undefined have their own distinct types and you’ll
	// get a type error if you try to use them where a concrete value is expected.
	// Default: true.
	//
	// Experimental.
	StrictNullChecks *bool `field:"optional" json:"strictNullChecks" yaml:"strictNullChecks"`
	// When set to true, TypeScript will raise an error when a class property was declared but not set in the constructor.
	// Default: true.
	//
	// Experimental.
	StrictPropertyInitialization *bool `field:"optional" json:"strictPropertyInitialization" yaml:"strictPropertyInitialization"`
	// Do not emit declarations for code that has an `@internal` annotation in it’s JSDoc comment.
	// Default: true.
	//
	// Experimental.
	StripInternal *bool `field:"optional" json:"stripInternal" yaml:"stripInternal"`
	// Modern browsers support all ES6 features, so ES6 is a good choice.
	//
	// You might choose to set
	// a lower target if your code is deployed to older environments, or a higher target if your
	// code is guaranteed to run in newer environments.
	// Default: "ES2018".
	//
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// This setting lets you specify a file for storing incremental compilation information as a part of composite projects which enables faster building of larger TypeScript codebases.
	//
	// You can read more about composite projects in the handbook.
	// Experimental.
	TsBuildInfoFile *string `field:"optional" json:"tsBuildInfoFile" yaml:"tsBuildInfoFile"`
	// If typeRoots is specified, only packages under typeRoots will be included.
	// See: https://www.typescriptlang.org/tsconfig/#typeRoots
	//
	// Experimental.
	TypeRoots *[]*string `field:"optional" json:"typeRoots" yaml:"typeRoots"`
	// If types is specified, only packages listed will be included in the global scope.
	// See: https://www.typescriptlang.org/tsconfig#types
	//
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Change the type of the variable in a catch clause from any to unknown Available with TypeScript 4.4 and newer.
	// Default: true.
	//
	// Experimental.
	UseUnknownInCatchVariables *bool `field:"optional" json:"useUnknownInCatchVariables" yaml:"useUnknownInCatchVariables"`
	// Simplifies TypeScript's handling of import/export `type` modifiers.
	// See: https://www.typescriptlang.org/tsconfig#verbatimModuleSyntax
	//
	// Default: undefined.
	//
	// Experimental.
	VerbatimModuleSyntax *bool `field:"optional" json:"verbatimModuleSyntax" yaml:"verbatimModuleSyntax"`
}

