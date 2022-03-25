package javascript

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/projen/projen-go/projen/jsii"

	"github.com/projen/projen-go/projen"
	"github.com/projen/projen-go/projen/build"
	"github.com/projen/projen-go/projen/github"
	"github.com/projen/projen-go/projen/github/workflows"
	"github.com/projen/projen-go/projen/javascript/internal"
	"github.com/projen/projen-go/projen/release"
	"github.com/projen/projen-go/projen/vscode"
)

// Options for `addBundle()`.
// Experimental.
type AddBundleOptions struct {
	// You can mark a file or a package as external to exclude it from your build.
	//
	// Instead of being bundled, the import will be preserved (using require for
	// the iife and cjs formats and using import for the esm format) and will be
	// evaluated at run time instead.
	//
	// This has several uses. First of all, it can be used to trim unnecessary
	// code from your bundle for a code path that you know will never be executed.
	// For example, a package may contain code that only runs in node but you will
	// only be using that package in the browser. It can also be used to import
	// code in node at run time from a package that cannot be bundled. For
	// example, the fsevents package contains a native extension, which esbuild
	// doesn't support.
	// Experimental.
	Externals *[]*string `json:"externals" yaml:"externals"`
	// Include a source map in the bundle.
	// Experimental.
	Sourcemap *bool `json:"sourcemap" yaml:"sourcemap"`
	// In addition to the `bundle:xyz` task, creates `bundle:xyz:watch` task which will invoke the same esbuild command with the `--watch` flag.
	//
	// This can be used
	// to continusouly watch for changes.
	// Experimental.
	WatchTask *bool `json:"watchTask" yaml:"watchTask"`
	// esbuild platform.
	//
	// Example:
	//   "node"
	//
	// Experimental.
	Platform *string `json:"platform" yaml:"platform"`
	// esbuild target.
	//
	// Example:
	//   "node12"
	//
	// Experimental.
	Target *string `json:"target" yaml:"target"`
	// Mark the output file as executable.
	// Experimental.
	Executable *bool `json:"executable" yaml:"executable"`
	// Bundler output path relative to the asset's output directory.
	// Experimental.
	Outfile *string `json:"outfile" yaml:"outfile"`
}

// Experimental.
type ArrowParens string

const (
	// Always include parens.
	//
	// Example: `(x) => x`.
	// Experimental.
	ArrowParens_ALWAYS ArrowParens = "ALWAYS"
	// Omit parens when possible.
	//
	// Example: `x => x`.
	// Experimental.
	ArrowParens_AVOID ArrowParens = "AVOID"
)

// Automatic bump modes.
// Experimental.
type AutoRelease string

const (
	// Automatically bump & release a new version for every commit to "main".
	// Experimental.
	AutoRelease_EVERY_COMMIT AutoRelease = "EVERY_COMMIT"
	// Automatically bump & release a new version on a daily basis.
	// Experimental.
	AutoRelease_DAILY AutoRelease = "DAILY"
)

// Experimental.
type Bundle struct {
	// The task that produces this bundle.
	// Experimental.
	BundleTask projen.Task `json:"bundleTask" yaml:"bundleTask"`
	// Base directory containing the output file (relative to project root).
	// Experimental.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// Location of the output file (relative to project root).
	// Experimental.
	Outfile *string `json:"outfile" yaml:"outfile"`
	// The "watch" task for this bundle.
	// Experimental.
	WatchTask projen.Task `json:"watchTask" yaml:"watchTask"`
}

// Adds support for bundling JavaScript applications and dependencies into a single file.
//
// In the future, this will also supports bundling websites.
// Experimental.
type Bundler interface {
	projen.Component
	// Root bundle directory.
	// Experimental.
	Bundledir() *string
	// Gets or creates the singleton "bundle" task of the project.
	//
	// If the project doesn't have a "bundle" task, it will be created and spawned
	// during the pre-compile phase.
	// Experimental.
	BundleTask() projen.Task
	// The semantic version requirement for `esbuild` (if defined).
	// Experimental.
	EsbuildVersion() *string
	// Experimental.
	Project() projen.Project
	// Adds a task to the project which bundles a specific entrypoint and all of its dependencies into a single javascript output file.
	// Experimental.
	AddBundle(entrypoint *string, options *AddBundleOptions) *Bundle
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Bundler
type jsiiProxy_Bundler struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Bundler) Bundledir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundledir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bundler) BundleTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"bundleTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bundler) EsbuildVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"esbuildVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bundler) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Creates a `Bundler`.
// Experimental.
func NewBundler(project projen.Project, options *BundlerOptions) Bundler {
	_init_.Initialize()

	j := jsiiProxy_Bundler{}

	_jsii_.Create(
		"projen.javascript.Bundler",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Creates a `Bundler`.
// Experimental.
func NewBundler_Override(b Bundler, project projen.Project, options *BundlerOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Bundler",
		[]interface{}{project, options},
		b,
	)
}

// Returns the `Bundler` instance associated with a project or `undefined` if there is no Bundler.
//
// Returns: A bundler.
// Experimental.
func Bundler_Of(project projen.Project) Bundler {
	_init_.Initialize()

	var returns Bundler

	_jsii_.StaticInvoke(
		"projen.javascript.Bundler",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bundler) AddBundle(entrypoint *string, options *AddBundleOptions) *Bundle {
	var returns *Bundle

	_jsii_.Invoke(
		b,
		"addBundle",
		[]interface{}{entrypoint, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bundler) PostSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"postSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Bundler) PreSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"preSynthesize",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Bundler) Synthesize() {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `Bundler`.
// Experimental.
type BundlerOptions struct {
	// Install the `bundle` command as a pre-compile phase.
	// Experimental.
	AddToPreCompile *bool `json:"addToPreCompile" yaml:"addToPreCompile"`
	// Output directory for all bundles.
	// Experimental.
	AssetsDir *string `json:"assetsDir" yaml:"assetsDir"`
	// The semantic version requirement for `esbuild`.
	// Experimental.
	EsbuildVersion *string `json:"esbuildVersion" yaml:"esbuildVersion"`
}

// Options for bundling.
// Experimental.
type BundlingOptions struct {
	// You can mark a file or a package as external to exclude it from your build.
	//
	// Instead of being bundled, the import will be preserved (using require for
	// the iife and cjs formats and using import for the esm format) and will be
	// evaluated at run time instead.
	//
	// This has several uses. First of all, it can be used to trim unnecessary
	// code from your bundle for a code path that you know will never be executed.
	// For example, a package may contain code that only runs in node but you will
	// only be using that package in the browser. It can also be used to import
	// code in node at run time from a package that cannot be bundled. For
	// example, the fsevents package contains a native extension, which esbuild
	// doesn't support.
	// Experimental.
	Externals *[]*string `json:"externals" yaml:"externals"`
	// Include a source map in the bundle.
	// Experimental.
	Sourcemap *bool `json:"sourcemap" yaml:"sourcemap"`
	// In addition to the `bundle:xyz` task, creates `bundle:xyz:watch` task which will invoke the same esbuild command with the `--watch` flag.
	//
	// This can be used
	// to continusouly watch for changes.
	// Experimental.
	WatchTask *bool `json:"watchTask" yaml:"watchTask"`
}

// Experimental.
type CodeArtifactOptions struct {
	// GitHub secret which contains the AWS access key ID to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`npmRegistryUrl` contains AWS CodeArtifact URL).
	// Experimental.
	AccessKeyIdSecret *string `json:"accessKeyIdSecret" yaml:"accessKeyIdSecret"`
	// ARN of AWS role to be assumed prior to get authorization token from AWS CodeArtifact This property must be specified only when publishing to AWS CodeArtifact (`registry` contains AWS CodeArtifact URL).
	// Experimental.
	RoleToAssume *string `json:"roleToAssume" yaml:"roleToAssume"`
	// GitHub secret which contains the AWS secret access key to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`npmRegistryUrl` contains AWS CodeArtifact URL).
	// Experimental.
	SecretAccessKeySecret *string `json:"secretAccessKeySecret" yaml:"secretAccessKeySecret"`
}

// Experimental.
type CoverageThreshold struct {
	// Experimental.
	Branches *float64 `json:"branches" yaml:"branches"`
	// Experimental.
	Functions *float64 `json:"functions" yaml:"functions"`
	// Experimental.
	Lines *float64 `json:"lines" yaml:"lines"`
	// Experimental.
	Statements *float64 `json:"statements" yaml:"statements"`
}

// Experimental.
type EmbeddedLanguageFormatting string

const (
	// Format embedded code if Prettier can automatically identify it.
	// Experimental.
	EmbeddedLanguageFormatting_AUTO EmbeddedLanguageFormatting = "AUTO"
	// Never automatically format embedded code.
	// Experimental.
	EmbeddedLanguageFormatting_OFF EmbeddedLanguageFormatting = "OFF"
)

// Experimental.
type EndOfLine string

const (
	// Maintain existing (mixed values within one file are normalised by looking at what's used after the first line).
	// Experimental.
	EndOfLine_AUTO EndOfLine = "AUTO"
	// Carriage Return character only (\r), used very rarely.
	// Experimental.
	EndOfLine_CR EndOfLine = "CR"
	// Carriage Return + Line Feed characters (\r\n), common on Windows.
	// Experimental.
	EndOfLine_CRLF EndOfLine = "CRLF"
	// Line Feed only (\n), common on Linux and macOS as well as inside git repos.
	// Experimental.
	EndOfLine_LF EndOfLine = "LF"
)

// Represents eslint configuration.
// Experimental.
type Eslint interface {
	projen.Component
	// Direct access to the eslint configuration (escape hatch).
	// Experimental.
	Config() interface{}
	// File patterns that should not be linted.
	// Experimental.
	IgnorePatterns() *[]*string
	// eslint overrides.
	// Experimental.
	Overrides() *[]*EslintOverride
	// Experimental.
	Project() projen.Project
	// eslint rules.
	// Experimental.
	Rules() *map[string]*[]interface{}
	// Adds an `extends` item to the eslint configuration.
	// Experimental.
	AddExtends(extendList ...*string)
	// Do not lint these files.
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Add an eslint override.
	// Experimental.
	AddOverride(override *EslintOverride)
	// Adds an eslint plugin.
	// Experimental.
	AddPlugins(plugins ...*string)
	// Add an eslint rule.
	// Experimental.
	AddRules(rules *map[string]interface{})
	// Add a glob file pattern which allows importing dev dependencies.
	// Experimental.
	AllowDevDeps(pattern *string)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Eslint
type jsiiProxy_Eslint struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Eslint) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) IgnorePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignorePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) Overrides() *[]*EslintOverride {
	var returns *[]*EslintOverride
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Eslint) Rules() *map[string]*[]interface{} {
	var returns *map[string]*[]interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


// Experimental.
func NewEslint(project NodeProject, options *EslintOptions) Eslint {
	_init_.Initialize()

	j := jsiiProxy_Eslint{}

	_jsii_.Create(
		"projen.javascript.Eslint",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewEslint_Override(e Eslint, project NodeProject, options *EslintOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Eslint",
		[]interface{}{project, options},
		e,
	)
}

// Returns the singletone Eslint component of a project or undefined if there is none.
// Experimental.
func Eslint_Of(project projen.Project) Eslint {
	_init_.Initialize()

	var returns Eslint

	_jsii_.StaticInvoke(
		"projen.javascript.Eslint",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Eslint) AddExtends(extendList ...*string) {
	args := []interface{}{}
	for _, a := range extendList {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addExtends",
		args,
	)
}

func (e *jsiiProxy_Eslint) AddIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		e,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (e *jsiiProxy_Eslint) AddOverride(override *EslintOverride) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{override},
	)
}

func (e *jsiiProxy_Eslint) AddPlugins(plugins ...*string) {
	args := []interface{}{}
	for _, a := range plugins {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addPlugins",
		args,
	)
}

func (e *jsiiProxy_Eslint) AddRules(rules *map[string]interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addRules",
		[]interface{}{rules},
	)
}

func (e *jsiiProxy_Eslint) AllowDevDeps(pattern *string) {
	_jsii_.InvokeVoid(
		e,
		"allowDevDeps",
		[]interface{}{pattern},
	)
}

func (e *jsiiProxy_Eslint) PostSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"postSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eslint) PreSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"preSynthesize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Eslint) Synthesize() {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type EslintOptions struct {
	// Directories with source files to lint (e.g. [ "src" ]).
	// Experimental.
	Dirs *[]*string `json:"dirs" yaml:"dirs"`
	// Enable import alias for module paths.
	// Experimental.
	AliasExtensions *[]*string `json:"aliasExtensions" yaml:"aliasExtensions"`
	// Enable import alias for module paths.
	// Experimental.
	AliasMap *map[string]*string `json:"aliasMap" yaml:"aliasMap"`
	// Directories with source files that include tests and build tools.
	//
	// These
	// sources are linted but may also import packages from `devDependencies`.
	// Experimental.
	Devdirs *[]*string `json:"devdirs" yaml:"devdirs"`
	// File types that should be linted (e.g. [ ".js", ".ts" ]).
	// Experimental.
	FileExtensions *[]*string `json:"fileExtensions" yaml:"fileExtensions"`
	// List of file patterns that should not be linted, using the same syntax as .gitignore patterns.
	// Experimental.
	IgnorePatterns *[]*string `json:"ignorePatterns" yaml:"ignorePatterns"`
	// Should we lint .projenrc.js.
	// Experimental.
	LintProjenRc *bool `json:"lintProjenRc" yaml:"lintProjenRc"`
	// Enable prettier for code formatting.
	// Experimental.
	Prettier *bool `json:"prettier" yaml:"prettier"`
	// Always try to resolve types under `<root>@types` directory even it doesn't contain any source code.
	//
	// This prevents `import/no-unresolved` eslint errors when importing a `@types/*` module that would otherwise remain unresolved.
	// Experimental.
	TsAlwaysTryTypes *bool `json:"tsAlwaysTryTypes" yaml:"tsAlwaysTryTypes"`
	// Path to `tsconfig.json` which should be used by eslint.
	// Experimental.
	TsconfigPath *string `json:"tsconfigPath" yaml:"tsconfigPath"`
}

// eslint rules override.
// Experimental.
type EslintOverride struct {
	// Files or file patterns on which to apply the override.
	// Experimental.
	Files *[]*string `json:"files" yaml:"files"`
	// The overriden rules.
	// Experimental.
	Rules *map[string]interface{} `json:"rules" yaml:"rules"`
}

// Experimental.
type HTMLWhitespaceSensitivity string

const (
	// Respect the default value of CSS display property.
	// Experimental.
	HTMLWhitespaceSensitivity_CSS HTMLWhitespaceSensitivity = "CSS"
	// Whitespaces are considered insignificant.
	// Experimental.
	HTMLWhitespaceSensitivity_IGNORE HTMLWhitespaceSensitivity = "IGNORE"
	// Whitespaces are considered significant.
	// Experimental.
	HTMLWhitespaceSensitivity_STRICT HTMLWhitespaceSensitivity = "STRICT"
)

// Experimental.
type HasteConfig struct {
	// Experimental.
	ComputeSha1 *bool `json:"computeSha1" yaml:"computeSha1"`
	// Experimental.
	DefaultPlatform *string `json:"defaultPlatform" yaml:"defaultPlatform"`
	// Experimental.
	HasteImplModulePath *string `json:"hasteImplModulePath" yaml:"hasteImplModulePath"`
	// Experimental.
	Platforms *[]*string `json:"platforms" yaml:"platforms"`
	// Experimental.
	ThrowOnModuleCollision *bool `json:"throwOnModuleCollision" yaml:"throwOnModuleCollision"`
}

// Installs the following npm scripts:.
//
// - `test` will run `jest --passWithNoTests`
// - `test:watch` will run `jest --watch`
// - `test:update` will run `jest -u`.
// Experimental.
type Jest interface {
	// Escape hatch.
	// Experimental.
	Config() interface{}
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Experimental.
	AddReporter(reporter interface{})
	// Experimental.
	AddSnapshotResolver(file *string)
	// Adds a test match pattern.
	// Experimental.
	AddTestMatch(pattern *string)
	// Adds a watch ignore pattern.
	// Experimental.
	AddWatchIgnorePattern(pattern *string)
}

// The jsii proxy struct for Jest
type jsiiProxy_Jest struct {
	_ byte // padding
}

func (j *jsiiProxy_Jest) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


// Experimental.
func NewJest(project NodeProject, options *JestOptions) Jest {
	_init_.Initialize()

	j := jsiiProxy_Jest{}

	_jsii_.Create(
		"projen.javascript.Jest",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewJest_Override(j Jest, project NodeProject, options *JestOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Jest",
		[]interface{}{project, options},
		j,
	)
}

func (j *jsiiProxy_Jest) AddIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_Jest) AddReporter(reporter interface{}) {
	_jsii_.InvokeVoid(
		j,
		"addReporter",
		[]interface{}{reporter},
	)
}

func (j *jsiiProxy_Jest) AddSnapshotResolver(file *string) {
	_jsii_.InvokeVoid(
		j,
		"addSnapshotResolver",
		[]interface{}{file},
	)
}

func (j *jsiiProxy_Jest) AddTestMatch(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addTestMatch",
		[]interface{}{pattern},
	)
}

func (j *jsiiProxy_Jest) AddWatchIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addWatchIgnorePattern",
		[]interface{}{pattern},
	)
}

// Experimental.
type JestConfigOptions struct {
	// This option tells Jest that all imported modules in your tests should be mocked automatically.
	//
	// All modules used in your tests will have a replacement implementation, keeping the API surface.
	// Experimental.
	Automock *bool `json:"automock" yaml:"automock"`
	// By default, Jest runs all tests and produces all errors into the console upon completion.
	//
	// The bail config option can be used here to have Jest stop running tests after n failures.
	// Setting bail to true is the same as setting bail to 1.
	// Experimental.
	Bail interface{} `json:"bail" yaml:"bail"`
	// The directory where Jest should store its cached dependency information.
	// Experimental.
	CacheDirectory *string `json:"cacheDirectory" yaml:"cacheDirectory"`
	// Automatically clear mock calls and instances before every test.
	//
	// Equivalent to calling jest.clearAllMocks() before each test.
	// This does not remove any mock implementation that may have been provided.
	// Experimental.
	ClearMocks *bool `json:"clearMocks" yaml:"clearMocks"`
	// Indicates whether the coverage information should be collected while executing the test.
	//
	// Because this retrofits all executed files with coverage collection statements,
	// it may significantly slow down your tests.
	// Experimental.
	CollectCoverage *bool `json:"collectCoverage" yaml:"collectCoverage"`
	// An array of glob patterns indicating a set of files for which coverage information should be collected.
	// Experimental.
	CollectCoverageFrom *[]*string `json:"collectCoverageFrom" yaml:"collectCoverageFrom"`
	// The directory where Jest should output its coverage files.
	// Experimental.
	CoverageDirectory *string `json:"coverageDirectory" yaml:"coverageDirectory"`
	// An array of regexp pattern strings that are matched against all file paths before executing the test.
	//
	// If the file path matches any of the patterns, coverage information will be skipped.
	// Experimental.
	CoveragePathIgnorePatterns *[]*string `json:"coveragePathIgnorePatterns" yaml:"coveragePathIgnorePatterns"`
	// Indicates which provider should be used to instrument code for coverage.
	//
	// Allowed values are babel (default) or v8.
	// Experimental.
	CoverageProvider *string `json:"coverageProvider" yaml:"coverageProvider"`
	// A list of reporter names that Jest uses when writing coverage reports.
	//
	// Any istanbul reporter can be used.
	// Experimental.
	CoverageReporters *[]*string `json:"coverageReporters" yaml:"coverageReporters"`
	// Specify the global coverage thresholds.
	//
	// This will be used to configure minimum threshold enforcement
	// for coverage results. Thresholds can be specified as global, as a glob, and as a directory or file path.
	// If thresholds aren't met, jest will fail.
	// Experimental.
	CoverageThreshold *CoverageThreshold `json:"coverageThreshold" yaml:"coverageThreshold"`
	// This option allows the use of a custom dependency extractor.
	//
	// It must be a node module that exports an object with an extract function.
	// Experimental.
	DependencyExtractor *string `json:"dependencyExtractor" yaml:"dependencyExtractor"`
	// Allows for a label to be printed alongside a test while it is running.
	// Experimental.
	DisplayName interface{} `json:"displayName" yaml:"displayName"`
	// Make calling deprecated APIs throw helpful error messages.
	//
	// Useful for easing the upgrade process.
	// Experimental.
	ErrorOnDeprecated *bool `json:"errorOnDeprecated" yaml:"errorOnDeprecated"`
	// Test files run inside a vm, which slows calls to global context properties (e.g. Math). With this option you can specify extra properties to be defined inside the vm for faster lookups.
	// Experimental.
	ExtraGlobals *[]*string `json:"extraGlobals" yaml:"extraGlobals"`
	// Test files are normally ignored from collecting code coverage.
	//
	// With this option, you can overwrite this behavior and include otherwise ignored files in code coverage.
	// Experimental.
	ForceCoverageMatch *[]*string `json:"forceCoverageMatch" yaml:"forceCoverageMatch"`
	// A set of global variables that need to be available in all test environments.
	// Experimental.
	Globals interface{} `json:"globals" yaml:"globals"`
	// This option allows the use of a custom global setup module which exports an async function that is triggered once before all test suites.
	//
	// This function gets Jest's globalConfig object as a parameter.
	// Experimental.
	GlobalSetup *string `json:"globalSetup" yaml:"globalSetup"`
	// This option allows the use of a custom global teardown module which exports an async function that is triggered once after all test suites.
	//
	// This function gets Jest's globalConfig object as a parameter.
	// Experimental.
	GlobalTeardown *string `json:"globalTeardown" yaml:"globalTeardown"`
	// This will be used to configure the behavior of jest-haste-map, Jest's internal file crawler/cache system.
	// Experimental.
	Haste *HasteConfig `json:"haste" yaml:"haste"`
	// Insert Jest's globals (expect, test, describe, beforeEach etc.) into the global environment. If you set this to false, you should import from @jest/globals.
	// Experimental.
	InjectGlobals *bool `json:"injectGlobals" yaml:"injectGlobals"`
	// A number limiting the number of tests that are allowed to run at the same time when using test.concurrent. Any test above this limit will be queued and executed once a slot is released.
	// Experimental.
	MaxConcurrency *float64 `json:"maxConcurrency" yaml:"maxConcurrency"`
	// Specifies the maximum number of workers the worker-pool will spawn for running tests.
	//
	// In single run mode,
	// this defaults to the number of the cores available on your machine minus one for the main thread
	// In watch mode, this defaults to half of the available cores on your machine.
	// For environments with variable CPUs available, you can use percentage based configuration: "maxWorkers": "50%".
	// Experimental.
	MaxWorkers interface{} `json:"maxWorkers" yaml:"maxWorkers"`
	// An array of directory names to be searched recursively up from the requiring module's location.
	//
	// Setting this option will override the default, if you wish to still search node_modules for packages
	// include it along with any other options: ["node_modules", "bower_components"].
	// Experimental.
	ModuleDirectories *[]*string `json:"moduleDirectories" yaml:"moduleDirectories"`
	// An array of file extensions your modules use.
	//
	// If you require modules without specifying a file extension,
	// these are the extensions Jest will look for, in left-to-right order.
	// Experimental.
	ModuleFileExtensions *[]*string `json:"moduleFileExtensions" yaml:"moduleFileExtensions"`
	// A map from regular expressions to module names or to arrays of module names that allow to stub out resources, like images or styles with a single module.
	// Experimental.
	ModuleNameMapper *map[string]interface{} `json:"moduleNameMapper" yaml:"moduleNameMapper"`
	// An array of regexp pattern strings that are matched against all module paths before those paths are to be considered 'visible' to the module loader.
	//
	// If a given module's path matches any of the patterns,
	// it will not be require()-able in the test environment.
	// Experimental.
	ModulePathIgnorePatterns *[]*string `json:"modulePathIgnorePatterns" yaml:"modulePathIgnorePatterns"`
	// An alternative API to setting the NODE_PATH env variable, modulePaths is an array of absolute paths to additional locations to search when resolving modules.
	//
	// Use the <rootDir> string token to include
	// the path to your project's root directory. Example: ["<rootDir>/app/"].
	// Experimental.
	ModulePaths *[]*string `json:"modulePaths" yaml:"modulePaths"`
	// Activates notifications for test results.
	// Experimental.
	Notify *bool `json:"notify" yaml:"notify"`
	// Specifies notification mode.
	//
	// Requires notify: true.
	// Experimental.
	NotifyMode *string `json:"notifyMode" yaml:"notifyMode"`
	// A preset that is used as a base for Jest's configuration.
	//
	// A preset should point to an npm module
	// that has a jest-preset.json or jest-preset.js file at the root.
	// Experimental.
	Preset *string `json:"preset" yaml:"preset"`
	// Sets the path to the prettier node module used to update inline snapshots.
	// Experimental.
	PrettierPath *string `json:"prettierPath" yaml:"prettierPath"`
	// When the projects configuration is provided with an array of paths or glob patterns, Jest will run tests in all of the specified projects at the same time.
	//
	// This is great for monorepos or
	// when working on multiple projects at the same time.
	// Experimental.
	Projects *[]interface{} `json:"projects" yaml:"projects"`
	// Use this configuration option to add custom reporters to Jest.
	//
	// A custom reporter is a class
	// that implements onRunStart, onTestStart, onTestResult, onRunComplete methods that will be
	// called when any of those events occurs.
	// Experimental.
	Reporters *[]interface{} `json:"reporters" yaml:"reporters"`
	// Automatically reset mock state before every test.
	//
	// Equivalent to calling jest.resetAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed but
	// does not restore their initial implementation.
	// Experimental.
	ResetMocks *bool `json:"resetMocks" yaml:"resetMocks"`
	// By default, each test file gets its own independent module registry.
	//
	// Enabling resetModules
	// goes a step further and resets the module registry before running each individual test.
	// Experimental.
	ResetModules *bool `json:"resetModules" yaml:"resetModules"`
	// This option allows the use of a custom resolver.
	//
	// https://jestjs.io/docs/en/configuration#resolver-string
	// Experimental.
	Resolver *string `json:"resolver" yaml:"resolver"`
	// Automatically restore mock state before every test.
	//
	// Equivalent to calling jest.restoreAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed and
	// restores their initial implementation.
	// Experimental.
	RestoreMocks *bool `json:"restoreMocks" yaml:"restoreMocks"`
	// The root directory that Jest should scan for tests and modules within.
	//
	// If you put your Jest
	// config inside your package.json and want the root directory to be the root of your repo, the
	// value for this config param will default to the directory of the package.json.
	// Experimental.
	RootDir *string `json:"rootDir" yaml:"rootDir"`
	// A list of paths to directories that Jest should use to search for files in.
	// Experimental.
	Roots *[]*string `json:"roots" yaml:"roots"`
	// This option allows you to use a custom runner instead of Jest's default test runner.
	// Experimental.
	Runner *string `json:"runner" yaml:"runner"`
	// A list of paths to modules that run some code to configure or set up the testing environment.
	//
	// Each setupFile will be run once per test file. Since every test runs in its own environment,
	// these scripts will be executed in the testing environment immediately before executing the
	// test code itself.
	// Experimental.
	SetupFiles *[]*string `json:"setupFiles" yaml:"setupFiles"`
	// A list of paths to modules that run some code to configure or set up the testing framework before each test file in the suite is executed.
	//
	// Since setupFiles executes before the test
	// framework is installed in the environment, this script file presents you the opportunity of
	// running some code immediately after the test framework has been installed in the environment.
	// Experimental.
	SetupFilesAfterEnv *[]*string `json:"setupFilesAfterEnv" yaml:"setupFilesAfterEnv"`
	// The number of seconds after which a test is considered as slow and reported as such in the results.
	// Experimental.
	SlowTestThreshold *float64 `json:"slowTestThreshold" yaml:"slowTestThreshold"`
	// The path to a module that can resolve test<->snapshot path.
	//
	// This config option lets you customize
	// where Jest stores snapshot files on disk.
	// Experimental.
	SnapshotResolver *string `json:"snapshotResolver" yaml:"snapshotResolver"`
	// A list of paths to snapshot serializer modules Jest should use for snapshot testing.
	// Experimental.
	SnapshotSerializers *[]*string `json:"snapshotSerializers" yaml:"snapshotSerializers"`
	// The test environment that will be used for testing.
	//
	// The default environment in Jest is a
	// browser-like environment through jsdom. If you are building a node service, you can use the node
	// option to use a node-like environment instead.
	// Experimental.
	TestEnvironment *string `json:"testEnvironment" yaml:"testEnvironment"`
	// Test environment options that will be passed to the testEnvironment.
	//
	// The relevant options depend on the environment.
	// Experimental.
	TestEnvironmentOptions interface{} `json:"testEnvironmentOptions" yaml:"testEnvironmentOptions"`
	// The exit code Jest returns on test failure.
	// Experimental.
	TestFailureExitCode *float64 `json:"testFailureExitCode" yaml:"testFailureExitCode"`
	// The glob patterns Jest uses to detect test files.
	//
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Experimental.
	TestMatch *[]*string `json:"testMatch" yaml:"testMatch"`
	// An array of regexp pattern strings that are matched against all test paths before executing the test.
	//
	// If the test path matches any of the patterns, it will be skipped.
	// Experimental.
	TestPathIgnorePatterns *[]*string `json:"testPathIgnorePatterns" yaml:"testPathIgnorePatterns"`
	// The pattern or patterns Jest uses to detect test files.
	//
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Experimental.
	TestRegex interface{} `json:"testRegex" yaml:"testRegex"`
	// This option allows the use of a custom results processor.
	// Experimental.
	TestResultsProcessor *string `json:"testResultsProcessor" yaml:"testResultsProcessor"`
	// This option allows the use of a custom test runner.
	//
	// The default is jasmine2. A custom test runner
	// can be provided by specifying a path to a test runner implementation.
	// Experimental.
	TestRunner *string `json:"testRunner" yaml:"testRunner"`
	// This option allows you to use a custom sequencer instead of Jest's default.
	//
	// Sort may optionally return a Promise.
	// Experimental.
	TestSequencer *string `json:"testSequencer" yaml:"testSequencer"`
	// Default timeout of a test in milliseconds.
	// Experimental.
	TestTimeout *float64 `json:"testTimeout" yaml:"testTimeout"`
	// This option sets the URL for the jsdom environment.
	//
	// It is reflected in properties such as location.href.
	// Experimental.
	TestURL *string `json:"testURL" yaml:"testURL"`
	// Setting this value to legacy or fake allows the use of fake timers for functions such as setTimeout.
	//
	// Fake timers are useful when a piece of code sets a long timeout that we don't want to wait for in a test.
	// Experimental.
	Timers *string `json:"timers" yaml:"timers"`
	// A map from regular expressions to paths to transformers.
	//
	// A transformer is a module that provides a
	// synchronous function for transforming source files.
	// Experimental.
	Transform *map[string]interface{} `json:"transform" yaml:"transform"`
	// An array of regexp pattern strings that are matched against all source file paths before transformation.
	//
	// If the test path matches any of the patterns, it will not be transformed.
	// Experimental.
	TransformIgnorePatterns *[]*string `json:"transformIgnorePatterns" yaml:"transformIgnorePatterns"`
	// An array of regexp pattern strings that are matched against all modules before the module loader will automatically return a mock for them.
	//
	// If a module's path matches any of the patterns in this list, it
	// will not be automatically mocked by the module loader.
	// Experimental.
	UnmockedModulePathPatterns *[]*string `json:"unmockedModulePathPatterns" yaml:"unmockedModulePathPatterns"`
	// Indicates whether each individual test should be reported during the run.
	//
	// All errors will also
	// still be shown on the bottom after execution. Note that if there is only one test file being run
	// it will default to true.
	// Experimental.
	Verbose *bool `json:"verbose" yaml:"verbose"`
	// Whether to use watchman for file crawling.
	// Experimental.
	Watchman *bool `json:"watchman" yaml:"watchman"`
	// An array of RegExp patterns that are matched against all source file paths before re-running tests in watch mode.
	//
	// If the file path matches any of the patterns, when it is updated, it will not trigger
	// a re-run of tests.
	// Experimental.
	WatchPathIgnorePatterns *[]*string `json:"watchPathIgnorePatterns" yaml:"watchPathIgnorePatterns"`
	// Experimental.
	WatchPlugins *map[string]interface{} `json:"watchPlugins" yaml:"watchPlugins"`
}

// Experimental.
type JestOptions struct {
	// Path to JSON config file for Jest.
	// Experimental.
	ConfigFilePath *string `json:"configFilePath" yaml:"configFilePath"`
	// Collect coverage.
	//
	// Deprecated.
	// Deprecated: use jestConfig.collectCoverage
	Coverage *bool `json:"coverage" yaml:"coverage"`
	// Include the `text` coverage reporter, which means that coverage summary is printed at the end of the jest execution.
	// Experimental.
	CoverageText *bool `json:"coverageText" yaml:"coverageText"`
	// Defines `testPathIgnorePatterns` and `coveragePathIgnorePatterns`.
	// Deprecated: use jestConfig.coveragePathIgnorePatterns or jestConfig.testPathIgnorePatterns respectively
	IgnorePatterns *[]*string `json:"ignorePatterns" yaml:"ignorePatterns"`
	// Jest configuration.
	// Experimental.
	JestConfig *JestConfigOptions `json:"jestConfig" yaml:"jestConfig"`
	// The version of jest to use.
	// Experimental.
	JestVersion *string `json:"jestVersion" yaml:"jestVersion"`
	// Result processing with jest-junit.
	//
	// Output directory is `test-reports/`.
	// Experimental.
	JunitReporting *bool `json:"junitReporting" yaml:"junitReporting"`
	// Preserve the default Jest reporter when additional reporters are added.
	// Experimental.
	PreserveDefaultReporters *bool `json:"preserveDefaultReporters" yaml:"preserveDefaultReporters"`
}

// Represents the npm `package.json` file.
// Experimental.
type NodePackage interface {
	projen.Component
	// Allow project to take library dependencies.
	// Experimental.
	AllowLibraryDependencies() *bool
	// Options for publishing npm package to AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions() *CodeArtifactOptions
	// The module's entrypoint (e.g. `lib/index.js`).
	// Experimental.
	Entrypoint() *string
	// Renders `yarn install` or `npm install` with lockfile update (not frozen).
	// Experimental.
	InstallAndUpdateLockfileCommand() *string
	// Returns the command to execute in order to install all dependencies (always frozen).
	// Experimental.
	InstallCommand() *string
	// The SPDX license of this module.
	//
	// `undefined` if this package is not licensed.
	// Experimental.
	License() *string
	// The name of the lock file.
	// Experimental.
	LockFile() *string
	// Deprecated: use `addField(x, y)`.
	Manifest() interface{}
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion() *string
	// npm package access level.
	// Experimental.
	NpmAccess() NpmAccess
	// The npm registry host (e.g. `registry.npmjs.org`).
	// Experimental.
	NpmRegistry() *string
	// npm registry (e.g. `https://registry.npmjs.org`). Use `npmRegistryHost` to get just the host name.
	// Experimental.
	NpmRegistryUrl() *string
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret() *string
	// The package manager to use.
	// Experimental.
	PackageManager() NodePackageManager
	// The name of the npm package.
	// Experimental.
	PackageName() *string
	// Experimental.
	Project() projen.Project
	// The command which executes "projen".
	// Experimental.
	ProjenCommand() *string
	// Experimental.
	AddBin(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Experimental.
	AddBundledDeps(deps ...*string)
	// Defines normal dependencies.
	// Experimental.
	AddDeps(deps ...*string)
	// Defines development/test dependencies.
	// Experimental.
	AddDevDeps(deps ...*string)
	// Adds an `engines` requirement to your package.
	// Experimental.
	AddEngine(engine *string, version *string)
	// Directly set fields in `package.json`.
	// Experimental.
	AddField(name *string, value interface{})
	// Adds keywords to package.json (deduplicated).
	// Experimental.
	AddKeywords(keywords ...*string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Experimental.
	AddPeerDeps(deps ...*string)
	// Sets the package version.
	// Experimental.
	AddVersion(version *string)
	// Indicates if a script by the given name is defined.
	// Deprecated: Use `project.tasks.tryFind(name)`
	HasScript(name *string) *bool
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Removes the npm script (always successful).
	// Experimental.
	RemoveScript(name *string)
	// Render a package manager specific command to upgrade all requested dependencies.
	// Experimental.
	RenderUpgradePackagesCommand(exclude *[]*string, include *[]*string) *string
	// Override the contents of an npm package.json script.
	// Experimental.
	SetScript(name *string, command *string)
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for NodePackage
type jsiiProxy_NodePackage struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NodePackage) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) CodeArtifactOptions() *CodeArtifactOptions {
	var returns *CodeArtifactOptions
	_jsii_.Get(
		j,
		"codeArtifactOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) InstallAndUpdateLockfileCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installAndUpdateLockfileCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) InstallCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"installCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) License() *string {
	var returns *string
	_jsii_.Get(
		j,
		"license",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) LockFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmAccess() NpmAccess {
	var returns NpmAccess
	_jsii_.Get(
		j,
		"npmAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmRegistry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"npmRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmRegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"npmRegistryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) NpmTokenSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"npmTokenSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) PackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) PackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodePackage) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodePackage(project projen.Project, options *NodePackageOptions) NodePackage {
	_init_.Initialize()

	j := jsiiProxy_NodePackage{}

	_jsii_.Create(
		"projen.javascript.NodePackage",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNodePackage_Override(n NodePackage, project projen.Project, options *NodePackageOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NodePackage",
		[]interface{}{project, options},
		n,
	)
}

func (n *jsiiProxy_NodePackage) AddBin(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		n,
		"addBin",
		[]interface{}{bins},
	)
}

func (n *jsiiProxy_NodePackage) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addBundledDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDevDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddEngine(engine *string, version *string) {
	_jsii_.InvokeVoid(
		n,
		"addEngine",
		[]interface{}{engine, version},
	)
}

func (n *jsiiProxy_NodePackage) AddField(name *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addField",
		[]interface{}{name, value},
	)
}

func (n *jsiiProxy_NodePackage) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addKeywords",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addPeerDeps",
		args,
	)
}

func (n *jsiiProxy_NodePackage) AddVersion(version *string) {
	_jsii_.InvokeVoid(
		n,
		"addVersion",
		[]interface{}{version},
	)
}

func (n *jsiiProxy_NodePackage) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		n,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodePackage) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodePackage) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodePackage) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		n,
		"removeScript",
		[]interface{}{name},
	)
}

func (n *jsiiProxy_NodePackage) RenderUpgradePackagesCommand(exclude *[]*string, include *[]*string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"renderUpgradePackagesCommand",
		[]interface{}{exclude, include},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodePackage) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		n,
		"setScript",
		[]interface{}{name, command},
	)
}

func (n *jsiiProxy_NodePackage) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

// The node package manager to use.
// Experimental.
type NodePackageManager string

const (
	// Use `yarn` as the package manager.
	// Experimental.
	NodePackageManager_YARN NodePackageManager = "YARN"
	// Use `npm` as the package manager.
	// Experimental.
	NodePackageManager_NPM NodePackageManager = "NPM"
	// Use `pnpm` as the package manager.
	// Experimental.
	NodePackageManager_PNPM NodePackageManager = "PNPM"
)

// Experimental.
type NodePackageOptions struct {
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `json:"allowLibraryDependencies" yaml:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName" yaml:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization *bool `json:"authorOrganization" yaml:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `json:"authorUrl" yaml:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `json:"autoDetectBin" yaml:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `json:"bin" yaml:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `json:"bugsEmail" yaml:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
	BugsUrl *string `json:"bugsUrl" yaml:"bugsUrl"`
	// List of dependencies to bundle into this module.
	//
	// These modules will be
	// added both to the `dependencies` section and `bundledDependencies` section of
	// your `package.json`.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps *[]*string `json:"bundledDeps" yaml:"bundledDeps"`
	// Options for publishing npm package to AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *CodeArtifactOptions `json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'express', 'lodash', 'foo@^2' ]
	//
	// Experimental.
	Deps *[]*string `json:"deps" yaml:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Build dependencies for this module.
	//
	// These dependencies will only be
	// available in your build environment but will not be fetched when this
	// module is consumed.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'typescript', '@types/express' ]
	//
	// Experimental.
	DevDeps *[]*string `json:"devDeps" yaml:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `json:"entrypoint" yaml:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `json:"homepage" yaml:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `json:"keywords" yaml:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `json:"license" yaml:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `json:"licensed" yaml:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `json:"maxNodeVersion" yaml:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `json:"minNodeVersion" yaml:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess" yaml:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead.
	NpmRegistry *string `json:"npmRegistry" yaml:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `json:"npmRegistryUrl" yaml:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager" yaml:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `json:"packageName" yaml:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions *PeerDependencyOptions `json:"peerDependencyOptions" yaml:"peerDependencyOptions"`
	// Peer dependencies for this module.
	//
	// Dependencies listed here are required to
	// be installed (and satisfied) by the _consumer_ of this library. Using peer
	// dependencies allows you to ensure that only a single module of a certain
	// library exists in the `node_modules` tree of your consumers.
	//
	// Note that prior to npm@7, peer dependencies are _not_ automatically
	// installed, which means that adding peer dependencies to a library will be a
	// breaking change for your customers.
	//
	// Unless `peerDependencyOptions.pinnedDevDependency` is disabled (it is
	// enabled by default), projen will automatically add a dev dependency with a
	// pinned version for each peer dependency. This will ensure that you build &
	// test your module against the lowest peer version required.
	// Experimental.
	PeerDeps *[]*string `json:"peerDeps" yaml:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `json:"repositoryDirectory" yaml:"repositoryDirectory"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts *map[string]*string `json:"scripts" yaml:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `json:"stability" yaml:"stability"`
}

// Node.js project.
// Experimental.
type NodeProject interface {
	github.GitHubProject
	// Deprecated: use `package.allowLibraryDependencies`
	AllowLibraryDependencies() *bool
	// The build output directory.
	//
	// An npm tarball will be created under the `js`
	// subdirectory. For example, if this is set to `dist` (the default), the npm
	// tarball will be placed under `dist/js/boom-boom-1.2.3.tg`.
	// Experimental.
	ArtifactsDirectory() *string
	// The location of the npm tarball after build (`${artifactsDirectory}/js`).
	// Experimental.
	ArtifactsJavascriptDirectory() *string
	// Auto approve set up for this project.
	// Deprecated.
	AutoApprove() github.AutoApprove
	// Automatic PR merges.
	// Experimental.
	AutoMerge() github.AutoMerge
	// Experimental.
	BuildTask() projen.Task
	// The PR build GitHub workflow.
	//
	// `undefined` if `buildWorkflow` is disabled.
	// Experimental.
	BuildWorkflow() build.BuildWorkflow
	// The job ID of the build workflow.
	// Experimental.
	BuildWorkflowJobId() *string
	// Experimental.
	Bundler() Bundler
	// Experimental.
	CompileTask() projen.Task
	// Returns all the components within this project.
	// Experimental.
	Components() *[]projen.Component
	// This is the "default" task, the one that executes "projen".
	//
	// Undefined if
	// the project is being ejected.
	// Experimental.
	DefaultTask() projen.Task
	// Project dependencies.
	// Experimental.
	Deps() projen.Dependencies
	// Access for .devcontainer.json (used for GitHub Codespaces).
	//
	// This will be `undefined` if devContainer boolean is false.
	// Deprecated.
	DevContainer() vscode.DevContainer
	// Whether or not the project is being ejected.
	// Experimental.
	Ejected() *bool
	// Deprecated: use `package.entrypoint`
	Entrypoint() *string
	// All files in this project.
	// Experimental.
	Files() *[]projen.FileBase
	// The .gitattributes file for this repository.
	// Experimental.
	Gitattributes() projen.GitAttributesFile
	// Access all github components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Github() github.GitHub
	// .gitignore.
	// Experimental.
	Gitignore() projen.IgnoreFile
	// Access for Gitpod.
	//
	// This will be `undefined` if gitpod boolean is false.
	// Deprecated.
	Gitpod() projen.Gitpod
	// The options used when this project is bootstrapped via `projen new`.
	//
	// It
	// includes the original set of options passed to the CLI and also the JSII
	// FQN of the project type.
	// Experimental.
	InitProject() *projen.InitProject
	// The Jest configuration (if enabled).
	// Experimental.
	Jest() Jest
	// Logging utilities.
	// Experimental.
	Logger() projen.Logger
	// Deprecated: use `package.addField(x, y)`
	Manifest() interface{}
	// Maximum node version required by this pacakge.
	// Experimental.
	MaxNodeVersion() *string
	// Minimum node.js version required by this package.
	// Experimental.
	MinNodeVersion() *string
	// Project name.
	// Experimental.
	Name() *string
	// The .npmignore file.
	// Experimental.
	Npmignore() projen.IgnoreFile
	// Absolute output directory of this project.
	// Experimental.
	Outdir() *string
	// API for managing the node package.
	// Experimental.
	Package() NodePackage
	// The package manager to use.
	// Deprecated: use `package.packageManager`
	PackageManager() NodePackageManager
	// Experimental.
	PackageTask() projen.Task
	// A parent project.
	//
	// If undefined, this is the root project.
	// Experimental.
	Parent() projen.Project
	// Experimental.
	PostCompileTask() projen.Task
	// Experimental.
	PreCompileTask() projen.Task
	// Experimental.
	Prettier() Prettier
	// Manages the build process of the project.
	// Experimental.
	ProjectBuild() projen.ProjectBuild
	// Deprecated.
	ProjectType() projen.ProjectType
	// The command to use in order to run the projen CLI.
	// Experimental.
	ProjenCommand() *string
	// Package publisher.
	//
	// This will be `undefined` if the project does not have a
	// release workflow.
	// Deprecated: use `release.publisher`.
	Publisher() release.Publisher
	// Release management.
	// Experimental.
	Release() release.Release
	// The root project.
	// Experimental.
	Root() projen.Project
	// The command to use to run scripts (e.g. `yarn run` or `npm run` depends on the package manager).
	// Experimental.
	RunScriptCommand() *string
	// Project tasks.
	// Experimental.
	Tasks() projen.Tasks
	// Experimental.
	TestTask() projen.Task
	// The upgrade workflow.
	// Experimental.
	UpgradeWorkflow() UpgradeDependencies
	// Access all VSCode components.
	//
	// This will be `undefined` for subprojects.
	// Deprecated.
	Vscode() vscode.VsCode
	// Experimental.
	AddBins(bins *map[string]*string)
	// Defines bundled dependencies.
	//
	// Bundled dependencies will be added as normal dependencies as well as to the
	// `bundledDependencies` section of your `package.json`.
	// Experimental.
	AddBundledDeps(deps ...*string)
	// DEPRECATED.
	// Deprecated: use `project.compileTask.exec()`
	AddCompileCommand(commands ...*string)
	// Defines normal dependencies.
	// Experimental.
	AddDeps(deps ...*string)
	// Defines development/test dependencies.
	// Experimental.
	AddDevDeps(deps ...*string)
	// Exclude the matching files from pre-synth cleanup.
	//
	// Can be used when, for example, some
	// source files include the projen marker and we don't want them to be erased during synth.
	// Experimental.
	AddExcludeFromCleanup(globs ...*string)
	// Directly set fields in `package.json`.
	// Experimental.
	AddFields(fields *map[string]interface{})
	// Adds a .gitignore pattern.
	// Experimental.
	AddGitIgnore(pattern *string)
	// Adds keywords to package.json (deduplicated).
	// Experimental.
	AddKeywords(keywords ...*string)
	// Exclude these files from the bundled package.
	//
	// Implemented by project types based on the
	// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
	// Experimental.
	AddPackageIgnore(pattern *string)
	// Defines peer dependencies.
	//
	// When adding peer dependencies, a devDependency will also be added on the
	// pinned version of the declared peer. This will ensure that you are testing
	// your code against the minimum version required from your consumers.
	// Experimental.
	AddPeerDeps(deps ...*string)
	// Adds a new task to this project.
	//
	// This will fail if the project already has
	// a task with this name.
	// Experimental.
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	// DEPRECATED.
	// Deprecated: use `project.testTask.exec()`
	AddTestCommand(commands ...*string)
	// Prints a "tip" message during synthesis.
	// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
	AddTip(message *string)
	// Marks the provided file(s) as being generated.
	//
	// This is achieved using the
	// github-linguist attributes. Generated files do not count against the
	// repository statistics and language breakdown.
	// See: https://github.com/github/linguist/blob/master/docs/overrides.md
	//
	// Deprecated.
	AnnotateGenerated(glob *string)
	// Indicates if a script by the name name is defined.
	// Experimental.
	HasScript(name *string) *bool
	// Called after all components are synthesized.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before all components are synthesized.
	// Experimental.
	PreSynthesize()
	// Removes the npm script (always successful).
	// Experimental.
	RemoveScript(name *string)
	// Removes a task from a project.
	//
	// Returns: The `Task` that was removed, otherwise `undefined`.
	// Experimental.
	RemoveTask(name *string) projen.Task
	// Returns the set of workflow steps which should be executed to bootstrap a workflow.
	//
	// Returns: Job steps.
	// Experimental.
	RenderWorkflowSetup(options *RenderWorkflowSetupOptions) *[]*workflows.JobStep
	// Returns the shell command to execute in order to run a task.
	//
	// This will
	// typically be `npx projen TASK`.
	// Experimental.
	RunTaskCommand(task projen.Task) *string
	// Replaces the contents of an npm package.json script.
	// Experimental.
	SetScript(name *string, command *string)
	// Synthesize all project files into `outdir`.
	//
	// 1. Call "this.preSynthesize()"
	// 2. Delete all generated files
	// 3. Synthesize all sub-projects
	// 4. Synthesize all components of this project
	// 5. Call "postSynthesize()" for all components of this project
	// 6. Call "this.postSynthesize()"
	// Experimental.
	Synth()
	// Finds a file at the specified relative path within this project and all its subprojects.
	//
	// Returns: a `FileBase` or undefined if there is no file in that path.
	// Experimental.
	TryFindFile(filePath *string) projen.FileBase
	// Finds a json file by name.
	// Deprecated: use `tryFindObjectFile`.
	TryFindJsonFile(filePath *string) projen.JsonFile
	// Finds an object file (like JsonFile, YamlFile, etc.) by name.
	// Experimental.
	TryFindObjectFile(filePath *string) projen.ObjectFile
}

// The jsii proxy struct for NodeProject
type jsiiProxy_NodeProject struct {
	internal.Type__githubGitHubProject
}

func (j *jsiiProxy_NodeProject) AllowLibraryDependencies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowLibraryDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) ArtifactsDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) ArtifactsJavascriptDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactsJavascriptDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) AutoApprove() github.AutoApprove {
	var returns github.AutoApprove
	_jsii_.Get(
		j,
		"autoApprove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) AutoMerge() github.AutoMerge {
	var returns github.AutoMerge
	_jsii_.Get(
		j,
		"autoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) BuildTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"buildTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) BuildWorkflow() build.BuildWorkflow {
	var returns build.BuildWorkflow
	_jsii_.Get(
		j,
		"buildWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) BuildWorkflowJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildWorkflowJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Bundler() Bundler {
	var returns Bundler
	_jsii_.Get(
		j,
		"bundler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) CompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"compileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Components() *[]projen.Component {
	var returns *[]projen.Component
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) DefaultTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"defaultTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Deps() projen.Dependencies {
	var returns projen.Dependencies
	_jsii_.Get(
		j,
		"deps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) DevContainer() vscode.DevContainer {
	var returns vscode.DevContainer
	_jsii_.Get(
		j,
		"devContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Ejected() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ejected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Entrypoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Files() *[]projen.FileBase {
	var returns *[]projen.FileBase
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Gitattributes() projen.GitAttributesFile {
	var returns projen.GitAttributesFile
	_jsii_.Get(
		j,
		"gitattributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Github() github.GitHub {
	var returns github.GitHub
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Gitignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"gitignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Gitpod() projen.Gitpod {
	var returns projen.Gitpod
	_jsii_.Get(
		j,
		"gitpod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) InitProject() *projen.InitProject {
	var returns *projen.InitProject
	_jsii_.Get(
		j,
		"initProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Jest() Jest {
	var returns Jest
	_jsii_.Get(
		j,
		"jest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Logger() projen.Logger {
	var returns projen.Logger
	_jsii_.Get(
		j,
		"logger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Manifest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) MaxNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) MinNodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minNodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Npmignore() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"npmignore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Package() NodePackage {
	var returns NodePackage
	_jsii_.Get(
		j,
		"package",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) PackageManager() NodePackageManager {
	var returns NodePackageManager
	_jsii_.Get(
		j,
		"packageManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) PackageTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"packageTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Parent() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) PostCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) PreCompileTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"preCompileTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Prettier() Prettier {
	var returns Prettier
	_jsii_.Get(
		j,
		"prettier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) ProjectBuild() projen.ProjectBuild {
	var returns projen.ProjectBuild
	_jsii_.Get(
		j,
		"projectBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) ProjectType() projen.ProjectType {
	var returns projen.ProjectType
	_jsii_.Get(
		j,
		"projectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) ProjenCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projenCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Publisher() release.Publisher {
	var returns release.Publisher
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Release() release.Release {
	var returns release.Release
	_jsii_.Get(
		j,
		"release",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Root() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"root",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) RunScriptCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runScriptCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Tasks() projen.Tasks {
	var returns projen.Tasks
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) TestTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"testTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) UpgradeWorkflow() UpgradeDependencies {
	var returns UpgradeDependencies
	_jsii_.Get(
		j,
		"upgradeWorkflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeProject) Vscode() vscode.VsCode {
	var returns vscode.VsCode
	_jsii_.Get(
		j,
		"vscode",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodeProject(options *NodeProjectOptions) NodeProject {
	_init_.Initialize()

	j := jsiiProxy_NodeProject{}

	_jsii_.Create(
		"projen.javascript.NodeProject",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewNodeProject_Override(n NodeProject, options *NodeProjectOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NodeProject",
		[]interface{}{options},
		n,
	)
}

func NodeProject_DEFAULT_TASK() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"projen.javascript.NodeProject",
		"DEFAULT_TASK",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NodeProject) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		n,
		"addBins",
		[]interface{}{bins},
	)
}

func (n *jsiiProxy_NodeProject) AddBundledDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addBundledDeps",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddCompileCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addCompileCommand",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDeps",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddDevDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDevDeps",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddExcludeFromCleanup(globs ...*string) {
	args := []interface{}{}
	for _, a := range globs {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addExcludeFromCleanup",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addFields",
		[]interface{}{fields},
	)
}

func (n *jsiiProxy_NodeProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

func (n *jsiiProxy_NodeProject) AddKeywords(keywords ...*string) {
	args := []interface{}{}
	for _, a := range keywords {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addKeywords",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

func (n *jsiiProxy_NodeProject) AddPeerDeps(deps ...*string) {
	args := []interface{}{}
	for _, a := range deps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addPeerDeps",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddTask(name *string, props *projen.TaskOptions) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		n,
		"addTask",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProject) AddTestCommand(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addTestCommand",
		args,
	)
}

func (n *jsiiProxy_NodeProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		n,
		"addTip",
		[]interface{}{message},
	)
}

func (n *jsiiProxy_NodeProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		n,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

func (n *jsiiProxy_NodeProject) HasScript(name *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		n,
		"hasScript",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeProject) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		n,
		"removeScript",
		[]interface{}{name},
	)
}

func (n *jsiiProxy_NodeProject) RemoveTask(name *string) projen.Task {
	var returns projen.Task

	_jsii_.Invoke(
		n,
		"removeTask",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProject) RenderWorkflowSetup(options *RenderWorkflowSetupOptions) *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep

	_jsii_.Invoke(
		n,
		"renderWorkflowSetup",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProject) RunTaskCommand(task projen.Task) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"runTaskCommand",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProject) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		n,
		"setScript",
		[]interface{}{name, command},
	)
}

func (n *jsiiProxy_NodeProject) Synth() {
	_jsii_.InvokeVoid(
		n,
		"synth",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NodeProject) TryFindFile(filePath *string) projen.FileBase {
	var returns projen.FileBase

	_jsii_.Invoke(
		n,
		"tryFindFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProject) TryFindJsonFile(filePath *string) projen.JsonFile {
	var returns projen.JsonFile

	_jsii_.Invoke(
		n,
		"tryFindJsonFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeProject) TryFindObjectFile(filePath *string) projen.ObjectFile {
	var returns projen.ObjectFile

	_jsii_.Invoke(
		n,
		"tryFindObjectFile",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Experimental.
type NodeProjectOptions struct {
	// This is the name of your project.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `json:"logging" yaml:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir" yaml:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `json:"parent" yaml:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand" yaml:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson" yaml:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions" yaml:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions" yaml:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions" yaml:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `json:"clobber" yaml:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `json:"devContainer" yaml:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `json:"github" yaml:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `json:"githubOptions" yaml:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod" yaml:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify" yaml:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `json:"mergifyOptions" yaml:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level.
	ProjectType projen.ProjectType `json:"projectType" yaml:"projectType"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Experimental.
	ProjenTokenSecret *string `json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// The README setup.
	//
	// Example:
	//   "{ filename: 'readme.md', contents: '# title' }"
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `json:"readme" yaml:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `json:"stale" yaml:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `json:"staleOptions" yaml:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode" yaml:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `json:"allowLibraryDependencies" yaml:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail" yaml:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName" yaml:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization *bool `json:"authorOrganization" yaml:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `json:"authorUrl" yaml:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `json:"autoDetectBin" yaml:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `json:"bin" yaml:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `json:"bugsEmail" yaml:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
	BugsUrl *string `json:"bugsUrl" yaml:"bugsUrl"`
	// List of dependencies to bundle into this module.
	//
	// These modules will be
	// added both to the `dependencies` section and `bundledDependencies` section of
	// your `package.json`.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	// Experimental.
	BundledDeps *[]*string `json:"bundledDeps" yaml:"bundledDeps"`
	// Options for publishing npm package to AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *CodeArtifactOptions `json:"codeArtifactOptions" yaml:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'express', 'lodash', 'foo@^2' ]
	//
	// Experimental.
	Deps *[]*string `json:"deps" yaml:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description *string `json:"description" yaml:"description"`
	// Build dependencies for this module.
	//
	// These dependencies will only be
	// available in your build environment but will not be fetched when this
	// module is consumed.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// Example:
	//   [ 'typescript', '@types/express' ]
	//
	// Experimental.
	DevDeps *[]*string `json:"devDeps" yaml:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `json:"entrypoint" yaml:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `json:"homepage" yaml:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `json:"keywords" yaml:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `json:"license" yaml:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `json:"licensed" yaml:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `json:"maxNodeVersion" yaml:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `json:"minNodeVersion" yaml:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess" yaml:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead.
	NpmRegistry *string `json:"npmRegistry" yaml:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `json:"npmRegistryUrl" yaml:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `json:"npmTokenSecret" yaml:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager" yaml:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `json:"packageName" yaml:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions *PeerDependencyOptions `json:"peerDependencyOptions" yaml:"peerDependencyOptions"`
	// Peer dependencies for this module.
	//
	// Dependencies listed here are required to
	// be installed (and satisfied) by the _consumer_ of this library. Using peer
	// dependencies allows you to ensure that only a single module of a certain
	// library exists in the `node_modules` tree of your consumers.
	//
	// Note that prior to npm@7, peer dependencies are _not_ automatically
	// installed, which means that adding peer dependencies to a library will be a
	// breaking change for your customers.
	//
	// Unless `peerDependencyOptions.pinnedDevDependency` is disabled (it is
	// enabled by default), projen will automatically add a dev dependency with a
	// pinned version for each peer dependency. This will ensure that you build &
	// test your module against the lowest peer version required.
	// Experimental.
	PeerDeps *[]*string `json:"peerDeps" yaml:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `json:"repository" yaml:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `json:"repositoryDirectory" yaml:"repositoryDirectory"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts *map[string]*string `json:"scripts" yaml:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `json:"stability" yaml:"stability"`
	// Version requirement of `publib` which is used to publish modules to npm.
	// Experimental.
	JsiiReleaseVersion *string `json:"jsiiReleaseVersion" yaml:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Experimental.
	MajorVersion *float64 `json:"majorVersion" yaml:"majorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Experimental.
	NpmDistTag *string `json:"npmDistTag" yaml:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `json:"postBuildSteps" yaml:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Experimental.
	Prerelease *string `json:"prerelease" yaml:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Experimental.
	PublishDryRun *bool `json:"publishDryRun" yaml:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `json:"publishTasks" yaml:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Experimental.
	ReleaseBranches *map[string]*release.BranchOptions `json:"releaseBranches" yaml:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `json:"releaseEveryCommit" yaml:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Experimental.
	ReleaseFailureIssue *bool `json:"releaseFailureIssue" yaml:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Experimental.
	ReleaseFailureIssueLabel *string `json:"releaseFailureIssueLabel" yaml:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `json:"releaseSchedule" yaml:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	ReleaseTagPrefix *string `json:"releaseTagPrefix" yaml:"releaseTagPrefix"`
	// The release trigger to use.
	// Experimental.
	ReleaseTrigger release.ReleaseTrigger `json:"releaseTrigger" yaml:"releaseTrigger"`
	// The name of the default release workflow.
	// Experimental.
	ReleaseWorkflowName *string `json:"releaseWorkflowName" yaml:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Experimental.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `json:"releaseWorkflowSetupSteps" yaml:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Experimental.
	VersionrcOptions *map[string]interface{} `json:"versionrcOptions" yaml:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage *string `json:"workflowContainerImage" yaml:"workflowContainerImage"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `json:"workflowRunsOn" yaml:"workflowRunsOn"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch *string `json:"defaultReleaseBranch" yaml:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveProjenUpgrades *bool `json:"autoApproveProjenUpgrades" yaml:"autoApproveProjenUpgrades"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveUpgrades *bool `json:"autoApproveUpgrades" yaml:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow *bool `json:"buildWorkflow" yaml:"buildWorkflow"`
	// Build workflow triggers.
	// Experimental.
	BuildWorkflowTriggers *workflows.Triggers `json:"buildWorkflowTriggers" yaml:"buildWorkflowTriggers"`
	// Options for `Bundler`.
	// Experimental.
	BundlerOptions *BundlerOptions `json:"bundlerOptions" yaml:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov *bool `json:"codeCov" yaml:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret *string `json:"codeCovTokenSecret" yaml:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner *string `json:"copyrightOwner" yaml:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod *string `json:"copyrightPeriod" yaml:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Experimental.
	Dependabot *bool `json:"dependabot" yaml:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions *github.DependabotOptions `json:"dependabotOptions" yaml:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Experimental.
	DepsUpgrade *bool `json:"depsUpgrade" yaml:"depsUpgrade"`
	// Options for depsUpgrade.
	// Experimental.
	DepsUpgradeOptions *UpgradeDependenciesOptions `json:"depsUpgradeOptions" yaml:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore *[]*string `json:"gitignore" yaml:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest *bool `json:"jest" yaml:"jest"`
	// Jest options.
	// Experimental.
	JestOptions *JestOptions `json:"jestOptions" yaml:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Experimental.
	MutableBuild *bool `json:"mutableBuild" yaml:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `json:"npmignore" yaml:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled *bool `json:"npmignoreEnabled" yaml:"npmignoreEnabled"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Experimental.
	Package *bool `json:"package" yaml:"package"`
	// Setup prettier.
	// Experimental.
	Prettier *bool `json:"prettier" yaml:"prettier"`
	// Prettier options.
	// Experimental.
	PrettierOptions *PrettierOptions `json:"prettierOptions" yaml:"prettierOptions"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency *bool `json:"projenDevDependency" yaml:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Experimental.
	ProjenrcJs *bool `json:"projenrcJs" yaml:"projenrcJs"`
	// Options for .projenrc.js.
	// Experimental.
	ProjenrcJsOptions *ProjenrcOptions `json:"projenrcJsOptions" yaml:"projenrcJsOptions"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Deprecated: use `autoApproveProjenUpgrades`.
	ProjenUpgradeAutoMerge *bool `json:"projenUpgradeAutoMerge" yaml:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule *[]*string `json:"projenUpgradeSchedule" yaml:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	//
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	//
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	//
	// To create a personal access token see https://github.com/settings/tokens
	// Deprecated: use `githubTokenSecret` instead.
	ProjenUpgradeSecret *string `json:"projenUpgradeSecret" yaml:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion *string `json:"projenVersion" yaml:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate *bool `json:"pullRequestTemplate" yaml:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents *[]*string `json:"pullRequestTemplateContents" yaml:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Experimental.
	Release *bool `json:"release" yaml:"release"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm *bool `json:"releaseToNpm" yaml:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `json:"releaseWorkflow" yaml:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps *[]*workflows.JobStep `json:"workflowBootstrapSteps" yaml:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Experimental.
	WorkflowGitIdentity *github.GitIdentity `json:"workflowGitIdentity" yaml:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion *string `json:"workflowNodeVersion" yaml:"workflowNodeVersion"`
}

// Npm package access level.
// Experimental.
type NpmAccess string

const (
	// Package is public.
	// Experimental.
	NpmAccess_PUBLIC NpmAccess = "PUBLIC"
	// Package can only be accessed with credentials.
	// Experimental.
	NpmAccess_RESTRICTED NpmAccess = "RESTRICTED"
)

// File representing the local NPM config in .npmrc.
// Experimental.
type NpmConfig interface {
	projen.Component
	// Experimental.
	Project() projen.Project
	// configure a generic property.
	// Experimental.
	AddConfig(name *string, value *string)
	// configure a scoped registry.
	// Experimental.
	AddRegistry(url *string, scope *string)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for NpmConfig
type jsiiProxy_NpmConfig struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_NpmConfig) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewNpmConfig(project NodeProject, options *NpmConfigOptions) NpmConfig {
	_init_.Initialize()

	j := jsiiProxy_NpmConfig{}

	_jsii_.Create(
		"projen.javascript.NpmConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewNpmConfig_Override(n NpmConfig, project NodeProject, options *NpmConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.NpmConfig",
		[]interface{}{project, options},
		n,
	)
}

func (n *jsiiProxy_NpmConfig) AddConfig(name *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"addConfig",
		[]interface{}{name, value},
	)
}

func (n *jsiiProxy_NpmConfig) AddRegistry(url *string, scope *string) {
	_jsii_.InvokeVoid(
		n,
		"addRegistry",
		[]interface{}{url, scope},
	)
}

func (n *jsiiProxy_NpmConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NpmConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NpmConfig) Synthesize() {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		nil, // no parameters
	)
}

// Options to configure the local NPM config.
// Experimental.
type NpmConfigOptions struct {
	// URL of the registry mirror to use.
	//
	// You can change this or add scoped registries using the addRegistry method.
	// Experimental.
	Registry *string `json:"registry" yaml:"registry"`
}

// Experimental.
type PeerDependencyOptions struct {
	// Automatically add a pinned dev dependency.
	// Experimental.
	PinnedDevDependency *bool `json:"pinnedDevDependency" yaml:"pinnedDevDependency"`
}

// Represents prettier configuration.
// Experimental.
type Prettier interface {
	projen.Component
	// The .prettierIgnore file.
	// Experimental.
	IgnoreFile() projen.IgnoreFile
	// Access to the Prettieroverrides to extend those.
	// Experimental.
	Overrides() *[]*PrettierOverride
	// Experimental.
	Project() projen.Project
	// Direct access to the prettier settings.
	// Experimental.
	Settings() *PrettierSettings
	// Defines Prettier ignore Patterns these patterns will be added to the file .prettierignore.
	// Experimental.
	AddIgnorePattern(pattern *string)
	// Add a prettier override.
	// See: https://prettier.io/docs/en/configuration.html#configuration-overrides
	//
	// Experimental.
	AddOverride(override *PrettierOverride)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Prettier
type jsiiProxy_Prettier struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Prettier) IgnoreFile() projen.IgnoreFile {
	var returns projen.IgnoreFile
	_jsii_.Get(
		j,
		"ignoreFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Overrides() *[]*PrettierOverride {
	var returns *[]*PrettierOverride
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prettier) Settings() *PrettierSettings {
	var returns *PrettierSettings
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrettier(project NodeProject, options *PrettierOptions) Prettier {
	_init_.Initialize()

	j := jsiiProxy_Prettier{}

	_jsii_.Create(
		"projen.javascript.Prettier",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewPrettier_Override(p Prettier, project NodeProject, options *PrettierOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Prettier",
		[]interface{}{project, options},
		p,
	)
}

// Experimental.
func Prettier_Of(project projen.Project) Prettier {
	_init_.Initialize()

	var returns Prettier

	_jsii_.StaticInvoke(
		"projen.javascript.Prettier",
		"of",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prettier) AddIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		p,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

func (p *jsiiProxy_Prettier) AddOverride(override *PrettierOverride) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{override},
	)
}

func (p *jsiiProxy_Prettier) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prettier) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Prettier) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Options for Prettier.
// Experimental.
type PrettierOptions struct {
	// Defines an .prettierIgnore file.
	// Experimental.
	IgnoreFile *bool `json:"ignoreFile" yaml:"ignoreFile"`
	// Provide a list of patterns to override prettier configuration.
	// See: https://prettier.io/docs/en/configuration.html#configuration-overrides
	//
	// Experimental.
	Overrides *[]*PrettierOverride `json:"overrides" yaml:"overrides"`
	// Prettier settings.
	// Experimental.
	Settings *PrettierSettings `json:"settings" yaml:"settings"`
}

// Experimental.
type PrettierOverride struct {
	// Include these files in this override.
	// Experimental.
	Files interface{} `json:"files" yaml:"files"`
	// The options to apply for this override.
	// Experimental.
	Settings *PrettierSettings `json:"settings" yaml:"settings"`
	// Exclude these files from this override.
	// Experimental.
	ExcludeFiles interface{} `json:"excludeFiles" yaml:"excludeFiles"`
}

// Options to set in Prettier directly or through overrides.
// See: https://prettier.io/docs/en/options.html
//
// Experimental.
type PrettierSettings struct {
	// Include parentheses around a sole arrow function parameter.
	// Experimental.
	ArrowParens ArrowParens `json:"arrowParens" yaml:"arrowParens"`
	// Put > of opening tags on the last line instead of on a new line.
	// Experimental.
	BracketSameLine *bool `json:"bracketSameLine" yaml:"bracketSameLine"`
	// Print spaces between brackets.
	// Experimental.
	BracketSpacing *bool `json:"bracketSpacing" yaml:"bracketSpacing"`
	// Print (to stderr) where a cursor at the given position would move to after formatting.
	//
	// This option cannot be used with --range-start and --range-end.
	// Experimental.
	CursorOffset *float64 `json:"cursorOffset" yaml:"cursorOffset"`
	// Control how Prettier formats quoted code embedded in the file.
	// Experimental.
	EmbeddedLanguageFormatting EmbeddedLanguageFormatting `json:"embeddedLanguageFormatting" yaml:"embeddedLanguageFormatting"`
	// Which end of line characters to apply.
	// Experimental.
	EndOfLine EndOfLine `json:"endOfLine" yaml:"endOfLine"`
	// Specify the input filepath.
	//
	// This will be used to do parser inference.
	// Experimental.
	Filepath *string `json:"filepath" yaml:"filepath"`
	// How to handle whitespaces in HTML.
	// Experimental.
	HtmlWhitespaceSensitivity HTMLWhitespaceSensitivity `json:"htmlWhitespaceSensitivity" yaml:"htmlWhitespaceSensitivity"`
	// Insert @format pragma into file's first docblock comment.
	// Experimental.
	InsertPragma *bool `json:"insertPragma" yaml:"insertPragma"`
	// Use single quotes in JSX.
	// Experimental.
	JsxSingleQuote *bool `json:"jsxSingleQuote" yaml:"jsxSingleQuote"`
	// Which parser to use.
	// Experimental.
	Parser *string `json:"parser" yaml:"parser"`
	// Add a plugin.
	//
	// Multiple plugins can be passed as separate `--plugin`s.
	// Experimental.
	Plugins *[]*string `json:"plugins" yaml:"plugins"`
	// Custom directory that contains prettier plugins in node_modules subdirectory.
	//
	// Overrides default behavior when plugins are searched relatively to the location of
	// Prettier.
	// Multiple values are accepted.
	// Experimental.
	PluginSearchDirs *[]*string `json:"pluginSearchDirs" yaml:"pluginSearchDirs"`
	// The line length where Prettier will try wrap.
	// Experimental.
	PrintWidth *float64 `json:"printWidth" yaml:"printWidth"`
	// How to wrap prose.
	// Experimental.
	ProseWrap ProseWrap `json:"proseWrap" yaml:"proseWrap"`
	// Change when properties in objects are quoted.
	// Experimental.
	QuoteProps QuoteProps `json:"quoteProps" yaml:"quoteProps"`
	// Format code ending at a given character offset (exclusive).
	//
	// The range will extend forwards to the end of the selected statement.
	// This option cannot be used with --cursor-offset.
	// Experimental.
	RangeEnd *float64 `json:"rangeEnd" yaml:"rangeEnd"`
	// Format code starting at a given character offset.
	//
	// The range will extend backwards to the start of the first line containing the selected
	// statement.
	// This option cannot be used with --cursor-offset.
	// Experimental.
	RangeStart *float64 `json:"rangeStart" yaml:"rangeStart"`
	// Require either '@prettier' or '@format' to be present in the file's first docblock comment in order for it to be formatted.
	// Experimental.
	RequirePragma *bool `json:"requirePragma" yaml:"requirePragma"`
	// Print semicolons.
	// Experimental.
	Semi *bool `json:"semi" yaml:"semi"`
	// Use single quotes instead of double quotes.
	// Experimental.
	SingleQuote *bool `json:"singleQuote" yaml:"singleQuote"`
	// Number of spaces per indentation level.
	// Experimental.
	TabWidth *float64 `json:"tabWidth" yaml:"tabWidth"`
	// Print trailing commas wherever possible when multi-line.
	// Experimental.
	TrailingComma TrailingComma `json:"trailingComma" yaml:"trailingComma"`
	// Indent with tabs instead of spaces.
	// Experimental.
	UseTabs *bool `json:"useTabs" yaml:"useTabs"`
	// Indent script and style tags in Vue files.
	// Experimental.
	VueIndentScriptAndStyle *bool `json:"vueIndentScriptAndStyle" yaml:"vueIndentScriptAndStyle"`
}

// Sets up a javascript project to use TypeScript for projenrc.
// Experimental.
type Projenrc interface {
	projen.Component
	// Experimental.
	Project() projen.Project
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for Projenrc
type jsiiProxy_Projenrc struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_Projenrc) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewProjenrc(project projen.Project, options *ProjenrcOptions) Projenrc {
	_init_.Initialize()

	j := jsiiProxy_Projenrc{}

	_jsii_.Create(
		"projen.javascript.Projenrc",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewProjenrc_Override(p Projenrc, project projen.Project, options *ProjenrcOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.Projenrc",
		[]interface{}{project, options},
		p,
	)
}

func (p *jsiiProxy_Projenrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Projenrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Projenrc) Synthesize() {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		nil, // no parameters
	)
}

// Experimental.
type ProjenrcOptions struct {
	// The name of the projenrc file.
	// Experimental.
	Filename *string `json:"filename" yaml:"filename"`
}

// Experimental.
type ProseWrap string

const (
	// Wrap prose if it exceeds the print width.
	// Experimental.
	ProseWrap_ALWAYS ProseWrap = "ALWAYS"
	// Do not wrap prose.
	// Experimental.
	ProseWrap_NEVER ProseWrap = "NEVER"
	// Wrap prose as-is.
	// Experimental.
	ProseWrap_PRESERVE ProseWrap = "PRESERVE"
)

// Experimental.
type QuoteProps string

const (
	// Only add quotes around object properties where required.
	// Experimental.
	QuoteProps_ASNEEDED QuoteProps = "ASNEEDED"
	// If at least one property in an object requires quotes, quote all properties.
	// Experimental.
	QuoteProps_CONSISTENT QuoteProps = "CONSISTENT"
	// Respect the input use of quotes in object properties.
	// Experimental.
	QuoteProps_PRESERVE QuoteProps = "PRESERVE"
)

// Options for `renderInstallSteps()`.
// Experimental.
type RenderWorkflowSetupOptions struct {
	// Should the pacakge lockfile be updated?
	// Experimental.
	Mutable *bool `json:"mutable" yaml:"mutable"`
}

// Experimental.
type TrailingComma string

const (
	// Trailing commas wherever possible (including function arguments).
	// Experimental.
	TrailingComma_ALL TrailingComma = "ALL"
	// Trailing commas where valid in ES5 (objects, arrays, etc.).
	// Experimental.
	TrailingComma_ES5 TrailingComma = "ES5"
	// No trailing commas.
	// Experimental.
	TrailingComma_NONE TrailingComma = "NONE"
)

// Experimental.
type TypeScriptCompilerOptions struct {
	// Allow JavaScript files to be compiled.
	// Experimental.
	AllowJs *bool `json:"allowJs" yaml:"allowJs"`
	// Allow default imports from modules with no default export.
	//
	// This does not affect code emit, just typechecking.
	// Experimental.
	AllowSyntheticDefaultImports *bool `json:"allowSyntheticDefaultImports" yaml:"allowSyntheticDefaultImports"`
	// Ensures that your files are parsed in the ECMAScript strict mode, and emit “use strict” for each source file.
	// Experimental.
	AlwaysStrict *bool `json:"alwaysStrict" yaml:"alwaysStrict"`
	// Lets you set a base directory to resolve non-absolute module names.
	//
	// You can define a root folder where you can do absolute file resolution.
	// Experimental.
	BaseUrl *string `json:"baseUrl" yaml:"baseUrl"`
	// To be specified along with the above.
	// Experimental.
	Declaration *bool `json:"declaration" yaml:"declaration"`
	// Offers a way to configure the root directory for where declaration files are emitted.
	// Experimental.
	DeclarationDir *string `json:"declarationDir" yaml:"declarationDir"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	//
	// Decorators are a language feature which hasn’t yet been fully ratified into the JavaScript specification.
	// This means that the implementation version in TypeScript may differ from the implementation in JavaScript when it it decided by TC39.
	// You can find out more about decorator support in TypeScript in the handbook.
	// See: https://www.typescriptlang.org/docs/handbook/decorators.html
	//
	// Experimental.
	EmitDecoratorMetadata *bool `json:"emitDecoratorMetadata" yaml:"emitDecoratorMetadata"`
	// Emit __importStar and __importDefault helpers for runtime babel ecosystem compatibility and enable --allowSyntheticDefaultImports for typesystem compatibility.
	// Experimental.
	EsModuleInterop *bool `json:"esModuleInterop" yaml:"esModuleInterop"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	// Experimental.
	ExperimentalDecorators *bool `json:"experimentalDecorators" yaml:"experimentalDecorators"`
	// Disallow inconsistently-cased references to the same file.
	// Experimental.
	ForceConsistentCasingInFileNames *bool `json:"forceConsistentCasingInFileNames" yaml:"forceConsistentCasingInFileNames"`
	// When set, instead of writing out a .js.map file to provide source maps, TypeScript will embed the source map content in the .js files.
	// Experimental.
	InlineSourceMap *bool `json:"inlineSourceMap" yaml:"inlineSourceMap"`
	// When set, TypeScript will include the original content of the .ts file as an embedded string in the source map. This is often useful in the same cases as inlineSourceMap.
	// Experimental.
	InlineSources *bool `json:"inlineSources" yaml:"inlineSources"`
	// Perform additional checks to ensure that separate compilation (such as with transpileModule or @babel/plugin-transform-typescript) would be safe.
	// Experimental.
	IsolatedModules *bool `json:"isolatedModules" yaml:"isolatedModules"`
	// Support JSX in .tsx files: "react", "preserve", "react-native" etc.
	// Experimental.
	Jsx TypeScriptJsxMode `json:"jsx" yaml:"jsx"`
	// Reference for type definitions / libraries to use (eg.
	//
	// ES2016, ES5, ES2018).
	// Experimental.
	Lib *[]*string `json:"lib" yaml:"lib"`
	// Sets the module system for the program.
	//
	// See https://www.typescriptlang.org/docs/handbook/modules.html#ambient-modules.
	// Experimental.
	Module *string `json:"module" yaml:"module"`
	// Determine how modules get resolved.
	//
	// Either "Node" for Node.js/io.js style resolution, or "Classic".
	// Experimental.
	ModuleResolution TypeScriptModuleResolution `json:"moduleResolution" yaml:"moduleResolution"`
	// Do not emit outputs.
	// Experimental.
	NoEmit *bool `json:"noEmit" yaml:"noEmit"`
	// Do not emit compiler output files like JavaScript source code, source-maps or declarations if any errors were reported.
	// Experimental.
	NoEmitOnError *bool `json:"noEmitOnError" yaml:"noEmitOnError"`
	// Report errors for fallthrough cases in switch statements.
	//
	// Ensures that any non-empty
	// case inside a switch statement includes either break or return. This means you won’t
	// accidentally ship a case fallthrough bug.
	// Experimental.
	NoFallthroughCasesInSwitch *bool `json:"noFallthroughCasesInSwitch" yaml:"noFallthroughCasesInSwitch"`
	// In some cases where no type annotations are present, TypeScript will fall back to a type of any for a variable when it cannot infer the type.
	// Experimental.
	NoImplicitAny *bool `json:"noImplicitAny" yaml:"noImplicitAny"`
	// When enabled, TypeScript will check all code paths in a function to ensure they return a value.
	// Experimental.
	NoImplicitReturns *bool `json:"noImplicitReturns" yaml:"noImplicitReturns"`
	// Raise error on ‘this’ expressions with an implied ‘any’ type.
	// Experimental.
	NoImplicitThis *bool `json:"noImplicitThis" yaml:"noImplicitThis"`
	// Raise error on use of the dot syntax to access fields which are not defined.
	// Experimental.
	NoPropertyAccessFromIndexSignature *bool `json:"noPropertyAccessFromIndexSignature" yaml:"noPropertyAccessFromIndexSignature"`
	// Raise error when accessing indexes on objects with unknown keys defined in index signatures.
	// Experimental.
	NoUncheckedIndexedAccess *bool `json:"noUncheckedIndexedAccess" yaml:"noUncheckedIndexedAccess"`
	// Report errors on unused local variables.
	// Experimental.
	NoUnusedLocals *bool `json:"noUnusedLocals" yaml:"noUnusedLocals"`
	// Report errors on unused parameters in functions.
	// Experimental.
	NoUnusedParameters *bool `json:"noUnusedParameters" yaml:"noUnusedParameters"`
	// Output directory for the compiled files.
	// Experimental.
	OutDir *string `json:"outDir" yaml:"outDir"`
	// A series of entries which re-map imports to lookup locations relative to the baseUrl, there is a larger coverage of paths in the handbook.
	//
	// paths lets you declare how TypeScript should resolve an import in your require/imports.
	// Experimental.
	Paths *map[string]*[]*string `json:"paths" yaml:"paths"`
	// Allows importing modules with a ‘.json’ extension, which is a common practice in node projects. This includes generating a type for the import based on the static JSON shape.
	// Experimental.
	ResolveJsonModule *bool `json:"resolveJsonModule" yaml:"resolveJsonModule"`
	// Specifies the root directory of input files.
	//
	// Only use to control the output directory structure with `outDir`.
	// Experimental.
	RootDir *string `json:"rootDir" yaml:"rootDir"`
	// Skip type checking of all declaration files (*.d.ts).
	// Experimental.
	SkipLibCheck *bool `json:"skipLibCheck" yaml:"skipLibCheck"`
	// The strict flag enables a wide range of type checking behavior that results in stronger guarantees of program correctness.
	//
	// Turning this on is equivalent to enabling all of the strict mode family
	// options, which are outlined below. You can then turn off individual strict mode family checks as
	// needed.
	// Experimental.
	Strict *bool `json:"strict" yaml:"strict"`
	// When strictNullChecks is false, null and undefined are effectively ignored by the language.
	//
	// This can lead to unexpected errors at runtime.
	// When strictNullChecks is true, null and undefined have their own distinct types and you’ll
	// get a type error if you try to use them where a concrete value is expected.
	// Experimental.
	StrictNullChecks *bool `json:"strictNullChecks" yaml:"strictNullChecks"`
	// When set to true, TypeScript will raise an error when a class property was declared but not set in the constructor.
	// Experimental.
	StrictPropertyInitialization *bool `json:"strictPropertyInitialization" yaml:"strictPropertyInitialization"`
	// Do not emit declarations for code that has an @internal annotation in it’s JSDoc comment.
	// Experimental.
	StripInternal *bool `json:"stripInternal" yaml:"stripInternal"`
	// Modern browsers support all ES6 features, so ES6 is a good choice.
	//
	// You might choose to set
	// a lower target if your code is deployed to older environments, or a higher target if your
	// code is guaranteed to run in newer environments.
	// Experimental.
	Target *string `json:"target" yaml:"target"`
}

// Determines how JSX should get transformed into valid JavaScript.
// See: https://www.typescriptlang.org/docs/handbook/jsx.html
//
// Experimental.
type TypeScriptJsxMode string

const (
	// Keeps the JSX as part of the output to be further consumed by another transform step (e.g. Babel).
	// Experimental.
	TypeScriptJsxMode_PRESERVE TypeScriptJsxMode = "PRESERVE"
	// Converts JSX syntax into React.createElement, does not need to go through a JSX transformation before use, and the output will have a .js file extension.
	// Experimental.
	TypeScriptJsxMode_REACT TypeScriptJsxMode = "REACT"
	// Keeps all JSX like 'preserve' mode, but output will have a .js extension.
	// Experimental.
	TypeScriptJsxMode_REACT_NATIVE TypeScriptJsxMode = "REACT_NATIVE"
	// Passes `key` separately from props and always passes `children` as props (since React 17).
	// See: https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-1.html#react-17-jsx-factories
	//
	// Experimental.
	TypeScriptJsxMode_REACT_JSX TypeScriptJsxMode = "REACT_JSX"
	// Same as `REACT_JSX` with additional debug data.
	// Experimental.
	TypeScriptJsxMode_REACT_JSXDEV TypeScriptJsxMode = "REACT_JSXDEV"
)

// Determines how modules get resolved.
// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html
//
// Experimental.
type TypeScriptModuleResolution string

const (
	// TypeScript's former default resolution strategy.
	// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html#classic
	//
	// Experimental.
	TypeScriptModuleResolution_CLASSIC TypeScriptModuleResolution = "CLASSIC"
	// Resolution strategy which attempts to mimic the Node.js module resolution strategy at runtime.
	// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html#node
	//
	// Experimental.
	TypeScriptModuleResolution_NODE TypeScriptModuleResolution = "NODE"
)

// Experimental.
type TypescriptConfig interface {
	// Experimental.
	CompilerOptions() *TypeScriptCompilerOptions
	// Experimental.
	Exclude() *[]*string
	// Experimental.
	File() projen.JsonFile
	// Experimental.
	FileName() *string
	// Experimental.
	Include() *[]*string
	// Experimental.
	AddExclude(pattern *string)
	// Experimental.
	AddInclude(pattern *string)
}

// The jsii proxy struct for TypescriptConfig
type jsiiProxy_TypescriptConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_TypescriptConfig) CompilerOptions() *TypeScriptCompilerOptions {
	var returns *TypeScriptCompilerOptions
	_jsii_.Get(
		j,
		"compilerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) File() projen.JsonFile {
	var returns projen.JsonFile
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TypescriptConfig) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}


// Experimental.
func NewTypescriptConfig(project NodeProject, options *TypescriptConfigOptions) TypescriptConfig {
	_init_.Initialize()

	j := jsiiProxy_TypescriptConfig{}

	_jsii_.Create(
		"projen.javascript.TypescriptConfig",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewTypescriptConfig_Override(t TypescriptConfig, project NodeProject, options *TypescriptConfigOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.TypescriptConfig",
		[]interface{}{project, options},
		t,
	)
}

func (t *jsiiProxy_TypescriptConfig) AddExclude(pattern *string) {
	_jsii_.InvokeVoid(
		t,
		"addExclude",
		[]interface{}{pattern},
	)
}

func (t *jsiiProxy_TypescriptConfig) AddInclude(pattern *string) {
	_jsii_.InvokeVoid(
		t,
		"addInclude",
		[]interface{}{pattern},
	)
}

// Experimental.
type TypescriptConfigOptions struct {
	// Compiler options to use.
	// Experimental.
	CompilerOptions *TypeScriptCompilerOptions `json:"compilerOptions" yaml:"compilerOptions"`
	// Filters results from the "include" option.
	// Experimental.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// Experimental.
	FileName *string `json:"fileName" yaml:"fileName"`
	// Specifies a list of glob patterns that match TypeScript files to be included in compilation.
	// Experimental.
	Include *[]*string `json:"include" yaml:"include"`
}

// Upgrade node project dependencies.
// Experimental.
type UpgradeDependencies interface {
	projen.Component
	// Container definitions for the upgrade workflow.
	// Experimental.
	ContainerOptions() *workflows.ContainerOptions
	// Experimental.
	SetContainerOptions(val *workflows.ContainerOptions)
	// Whether or not projen is also upgraded in this workflow,.
	// Experimental.
	IgnoresProjen() *bool
	// A task run after the upgrade task.
	// Experimental.
	PostUpgradeTask() projen.Task
	// Experimental.
	Project() projen.Project
	// The upgrade task.
	// Experimental.
	UpgradeTask() projen.Task
	// The workflows that execute the upgrades.
	//
	// One workflow per branch.
	// Experimental.
	Workflows() *[]github.GithubWorkflow
	// Add steps to execute a successful build.
	// Experimental.
	AddPostBuildSteps(steps ...*workflows.JobStep)
	// Called after synthesis.
	//
	// Order is *not* guaranteed.
	// Experimental.
	PostSynthesize()
	// Called before synthesis.
	// Experimental.
	PreSynthesize()
	// Synthesizes files to the project output directory.
	// Experimental.
	Synthesize()
}

// The jsii proxy struct for UpgradeDependencies
type jsiiProxy_UpgradeDependencies struct {
	internal.Type__projenComponent
}

func (j *jsiiProxy_UpgradeDependencies) ContainerOptions() *workflows.ContainerOptions {
	var returns *workflows.ContainerOptions
	_jsii_.Get(
		j,
		"containerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) IgnoresProjen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignoresProjen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) PostUpgradeTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"postUpgradeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) UpgradeTask() projen.Task {
	var returns projen.Task
	_jsii_.Get(
		j,
		"upgradeTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpgradeDependencies) Workflows() *[]github.GithubWorkflow {
	var returns *[]github.GithubWorkflow
	_jsii_.Get(
		j,
		"workflows",
		&returns,
	)
	return returns
}


// Experimental.
func NewUpgradeDependencies(project NodeProject, options *UpgradeDependenciesOptions) UpgradeDependencies {
	_init_.Initialize()

	j := jsiiProxy_UpgradeDependencies{}

	_jsii_.Create(
		"projen.javascript.UpgradeDependencies",
		[]interface{}{project, options},
		&j,
	)

	return &j
}

// Experimental.
func NewUpgradeDependencies_Override(u UpgradeDependencies, project NodeProject, options *UpgradeDependenciesOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"projen.javascript.UpgradeDependencies",
		[]interface{}{project, options},
		u,
	)
}

func (j *jsiiProxy_UpgradeDependencies) SetContainerOptions(val *workflows.ContainerOptions) {
	_jsii_.Set(
		j,
		"containerOptions",
		val,
	)
}

func (u *jsiiProxy_UpgradeDependencies) AddPostBuildSteps(steps ...*workflows.JobStep) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addPostBuildSteps",
		args,
	)
}

func (u *jsiiProxy_UpgradeDependencies) PostSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"postSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpgradeDependencies) PreSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"preSynthesize",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpgradeDependencies) Synthesize() {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		nil, // no parameters
	)
}

// Options for `UpgradeDependencies`.
// Experimental.
type UpgradeDependenciesOptions struct {
	// List of package names to exclude during the upgrade.
	// Experimental.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// Whether or not to ignore projen upgrades.
	// Experimental.
	IgnoreProjen *bool `json:"ignoreProjen" yaml:"ignoreProjen"`
	// List of package names to include during the upgrade.
	// Experimental.
	Include *[]*string `json:"include" yaml:"include"`
	// Title of the pull request to use (should be all lower-case).
	// Experimental.
	PullRequestTitle *string `json:"pullRequestTitle" yaml:"pullRequestTitle"`
	// Add Signed-off-by line by the committer at the end of the commit log message.
	// Experimental.
	Signoff *bool `json:"signoff" yaml:"signoff"`
	// The name of the task that will be created.
	//
	// This will also be the workflow name.
	// Experimental.
	TaskName *string `json:"taskName" yaml:"taskName"`
	// Include a github workflow for creating PR's that upgrades the required dependencies, either by manual dispatch, or by a schedule.
	//
	// If this is `false`, only a local projen task is created, which can be executed manually to
	// upgrade the dependencies.
	// Experimental.
	Workflow *bool `json:"workflow" yaml:"workflow"`
	// Options for the github workflow.
	//
	// Only applies if `workflow` is true.
	// Experimental.
	WorkflowOptions *UpgradeDependenciesWorkflowOptions `json:"workflowOptions" yaml:"workflowOptions"`
}

// How often to check for new versions and raise pull requests for version upgrades.
// Experimental.
type UpgradeDependenciesSchedule interface {
	// Experimental.
	Cron() *[]*string
}

// The jsii proxy struct for UpgradeDependenciesSchedule
type jsiiProxy_UpgradeDependenciesSchedule struct {
	_ byte // padding
}

func (j *jsiiProxy_UpgradeDependenciesSchedule) Cron() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cron",
		&returns,
	)
	return returns
}


// Create a schedule from a raw cron expression.
// Experimental.
func UpgradeDependenciesSchedule_Expressions(cron *[]*string) UpgradeDependenciesSchedule {
	_init_.Initialize()

	var returns UpgradeDependenciesSchedule

	_jsii_.StaticInvoke(
		"projen.javascript.UpgradeDependenciesSchedule",
		"expressions",
		[]interface{}{cron},
		&returns,
	)

	return returns
}

func UpgradeDependenciesSchedule_DAILY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"DAILY",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_MONTHLY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"MONTHLY",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_NEVER() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"NEVER",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_WEEKDAY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"WEEKDAY",
		&returns,
	)
	return returns
}

func UpgradeDependenciesSchedule_WEEKLY() UpgradeDependenciesSchedule {
	_init_.Initialize()
	var returns UpgradeDependenciesSchedule
	_jsii_.StaticGet(
		"projen.javascript.UpgradeDependenciesSchedule",
		"WEEKLY",
		&returns,
	)
	return returns
}

// Options for `UpgradeDependencies.workflowOptions`.
// Experimental.
type UpgradeDependenciesWorkflowOptions struct {
	// Assignees to add on the PR.
	// Experimental.
	Assignees *[]*string `json:"assignees" yaml:"assignees"`
	// List of branches to create PR's for.
	// Experimental.
	Branches *[]*string `json:"branches" yaml:"branches"`
	// Job container options.
	// Experimental.
	Container *workflows.ContainerOptions `json:"container" yaml:"container"`
	// The git identity to use for commits.
	// Experimental.
	GitIdentity *github.GitIdentity `json:"gitIdentity" yaml:"gitIdentity"`
	// Labels to apply on the PR.
	// Experimental.
	Labels *[]*string `json:"labels" yaml:"labels"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn" yaml:"runsOn"`
	// Schedule to run on.
	// Experimental.
	Schedule UpgradeDependenciesSchedule `json:"schedule" yaml:"schedule"`
	// Which secret to use when creating the PR.
	//
	// When using the default github token, PR's created by this workflow
	// will not trigger any subsequent workflows (i.e the build workflow).
	// This is why this workflow also runs 'build' by default, and manually updates
	// the status check of the PR.
	//
	// If you pass a token that has the `workflow` permissions, you can skip running
	// build in this workflow by specifying `rebuild: false`.
	// See: https://github.com/peter-evans/create-pull-request/issues/48
	//
	// Experimental.
	Secret *string `json:"secret" yaml:"secret"`
}

