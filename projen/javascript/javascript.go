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
	Externals *[]*string `json:"externals"`
	// Include a source map in the bundle.
	// Experimental.
	Sourcemap *bool `json:"sourcemap"`
	// In addition to the `bundle:xyz` task, creates `bundle:xyz:watch` task which will invoke the same esbuild command with the `--watch` flag.
	//
	// This can be used
	// to continusouly watch for changes.
	// Experimental.
	WatchTask *bool `json:"watchTask"`
	// esbuild platform.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Platform *string `json:"platform"`
	// esbuild target.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Target *string `json:"target"`
}

// Automatic bump modes.
// Experimental.
type AutoRelease string

const (
	AutoRelease_EVERY_COMMIT AutoRelease = "EVERY_COMMIT"
	AutoRelease_DAILY AutoRelease = "DAILY"
)

// Experimental.
type Bundle struct {
	// The task that produces this bundle.
	// Experimental.
	BundleTask projen.Task `json:"bundleTask"`
	// Location of the output file (relative to project root).
	// Experimental.
	Outfile *string `json:"outfile"`
	// The "watch" task for this bundle.
	// Experimental.
	WatchTask projen.Task `json:"watchTask"`
}

// Adds support for bundling JavaScript applications and dependencies into a single file.
//
// In the future, this will also supports bundling websites.
// Experimental.
type Bundler interface {
	projen.Component
	Bundledir() *string
	BundleTask() projen.Task
	EsbuildVersion() *string
	Project() projen.Project
	AddBundle(entrypoint *string, options *AddBundleOptions) *Bundle
	PostSynthesize()
	PreSynthesize()
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
// Returns: A bundler
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

// Adds a task to the project which bundles a specific entrypoint and all of its dependencies into a single javascript output file.
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (b *jsiiProxy_Bundler) PostSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (b *jsiiProxy_Bundler) PreSynthesize() {
	_jsii_.InvokeVoid(
		b,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	AddToPreCompile *bool `json:"addToPreCompile"`
	// Output directory for all bundles.
	// Experimental.
	AssetsDir *string `json:"assetsDir"`
	// The semantic version requirement for `esbuild`.
	// Experimental.
	EsbuildVersion *string `json:"esbuildVersion"`
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
	Externals *[]*string `json:"externals"`
	// Include a source map in the bundle.
	// Experimental.
	Sourcemap *bool `json:"sourcemap"`
	// In addition to the `bundle:xyz` task, creates `bundle:xyz:watch` task which will invoke the same esbuild command with the `--watch` flag.
	//
	// This can be used
	// to continusouly watch for changes.
	// Experimental.
	WatchTask *bool `json:"watchTask"`
}

// Experimental.
type CodeArtifactOptions struct {
	// GitHub secret which contains the AWS access key ID to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`npmRegistryUrl` contains AWS CodeArtifact URL).
	// Experimental.
	AccessKeyIdSecret *string `json:"accessKeyIdSecret"`
	// ARN of AWS role to be assumed prior to get authorization token from AWS CodeArtifact This property must be specified only when publishing to AWS CodeArtifact (`registry` contains AWS CodeArtifact URL).
	// Experimental.
	RoleToAssume *string `json:"roleToAssume"`
	// GitHub secret which contains the AWS secret access key to use when publishing packages to AWS CodeArtifact.
	//
	// This property must be specified only when publishing to AWS CodeArtifact (`npmRegistryUrl` contains AWS CodeArtifact URL).
	// Experimental.
	SecretAccessKeySecret *string `json:"secretAccessKeySecret"`
}

// Experimental.
type CoverageThreshold struct {
	// Experimental.
	Branches *float64 `json:"branches"`
	// Experimental.
	Functions *float64 `json:"functions"`
	// Experimental.
	Lines *float64 `json:"lines"`
	// Experimental.
	Statements *float64 `json:"statements"`
}

// Represents eslint configuration.
// Experimental.
type Eslint interface {
	projen.Component
	Config() interface{}
	IgnorePatterns() *[]*string
	Overrides() *[]*EslintOverride
	Project() projen.Project
	Rules() *map[string]*[]interface{}
	AddIgnorePattern(pattern *string)
	AddOverride(override *EslintOverride)
	AddRules(rules *map[string]interface{})
	AllowDevDeps(pattern *string)
	PostSynthesize()
	PreSynthesize()
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

// Do not lint these files.
// Experimental.
func (e *jsiiProxy_Eslint) AddIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		e,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

// Add an eslint override.
// Experimental.
func (e *jsiiProxy_Eslint) AddOverride(override *EslintOverride) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{override},
	)
}

// Add an eslint rule.
// Experimental.
func (e *jsiiProxy_Eslint) AddRules(rules *map[string]interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addRules",
		[]interface{}{rules},
	)
}

// Add a glob file pattern which allows importing dev dependencies.
// Experimental.
func (e *jsiiProxy_Eslint) AllowDevDeps(pattern *string) {
	_jsii_.InvokeVoid(
		e,
		"allowDevDeps",
		[]interface{}{pattern},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (e *jsiiProxy_Eslint) PostSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (e *jsiiProxy_Eslint) PreSynthesize() {
	_jsii_.InvokeVoid(
		e,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Dirs *[]*string `json:"dirs"`
	// Enable import alias for module paths.
	// Experimental.
	AliasExtensions *[]*string `json:"aliasExtensions"`
	// Enable import alias for module paths.
	// Experimental.
	AliasMap *map[string]*string `json:"aliasMap"`
	// Directories with source files that include tests and build tools.
	//
	// These
	// sources are linted but may also import packages from `devDependencies`.
	// Experimental.
	Devdirs *[]*string `json:"devdirs"`
	// File types that should be linted (e.g. [ ".js", ".ts" ]).
	// Experimental.
	FileExtensions *[]*string `json:"fileExtensions"`
	// List of file patterns that should not be linted, using the same syntax as .gitignore patterns.
	// Experimental.
	IgnorePatterns *[]*string `json:"ignorePatterns"`
	// Should we lint .projenrc.js.
	// Experimental.
	LintProjenRc *bool `json:"lintProjenRc"`
	// Enable prettier for code formatting.
	// Experimental.
	Prettier *bool `json:"prettier"`
	// Always try to resolve types under `<root>@types` directory even it doesn't contain any source code.
	//
	// This prevents `import/no-unresolved` eslint errors when importing a `@types/*` module that would otherwise remain unresolved.
	// Experimental.
	TsAlwaysTryTypes *bool `json:"tsAlwaysTryTypes"`
	// Path to `tsconfig.json` which should be used by eslint.
	// Experimental.
	TsconfigPath *string `json:"tsconfigPath"`
}

// eslint rules override.
// Experimental.
type EslintOverride struct {
	// Files or file patterns on which to apply the override.
	// Experimental.
	Files *[]*string `json:"files"`
	// The overriden rules.
	// Experimental.
	Rules *map[string]interface{} `json:"rules"`
}

// Experimental.
type HasteConfig struct {
	// Experimental.
	ComputeSha1 *bool `json:"computeSha1"`
	// Experimental.
	DefaultPlatform *string `json:"defaultPlatform"`
	// Experimental.
	HasteImplModulePath *string `json:"hasteImplModulePath"`
	// Experimental.
	Platforms *[]*string `json:"platforms"`
	// Experimental.
	ThrowOnModuleCollision *bool `json:"throwOnModuleCollision"`
}

// Installs the following npm scripts:.
//
// - `test` will run `jest --passWithNoTests`
// - `test:watch` will run `jest --watch`
// - `test:update` will run `jest -u`
// Experimental.
type Jest interface {
	Config() interface{}
	AddIgnorePattern(pattern *string)
	AddReporter(reporter interface{})
	AddSnapshotResolver(file *string)
	AddTestMatch(pattern *string)
	AddTypeScriptSupport(tsconfig TypescriptConfig)
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

// Experimental.
func (j *jsiiProxy_Jest) AddIgnorePattern(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addIgnorePattern",
		[]interface{}{pattern},
	)
}

// Experimental.
func (j *jsiiProxy_Jest) AddReporter(reporter interface{}) {
	_jsii_.InvokeVoid(
		j,
		"addReporter",
		[]interface{}{reporter},
	)
}

// Experimental.
func (j *jsiiProxy_Jest) AddSnapshotResolver(file *string) {
	_jsii_.InvokeVoid(
		j,
		"addSnapshotResolver",
		[]interface{}{file},
	)
}

// Adds a test match pattern.
// Experimental.
func (j *jsiiProxy_Jest) AddTestMatch(pattern *string) {
	_jsii_.InvokeVoid(
		j,
		"addTestMatch",
		[]interface{}{pattern},
	)
}

// Configures jest for TypeScript.
// Experimental.
func (j *jsiiProxy_Jest) AddTypeScriptSupport(tsconfig TypescriptConfig) {
	_jsii_.InvokeVoid(
		j,
		"addTypeScriptSupport",
		[]interface{}{tsconfig},
	)
}

// Adds a watch ignore pattern.
// Experimental.
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
	// All modules used in your tests will have a replacement implementation, keeping the API surface
	// Experimental.
	Automock *bool `json:"automock"`
	// By default, Jest runs all tests and produces all errors into the console upon completion.
	//
	// The bail config option can be used here to have Jest stop running tests after n failures.
	// Setting bail to true is the same as setting bail to 1.
	// Experimental.
	Bail interface{} `json:"bail"`
	// The directory where Jest should store its cached dependency information.
	// Experimental.
	CacheDirectory *string `json:"cacheDirectory"`
	// Automatically clear mock calls and instances before every test.
	//
	// Equivalent to calling jest.clearAllMocks() before each test.
	// This does not remove any mock implementation that may have been provided
	// Experimental.
	ClearMocks *bool `json:"clearMocks"`
	// Indicates whether the coverage information should be collected while executing the test.
	//
	// Because this retrofits all executed files with coverage collection statements,
	// it may significantly slow down your tests
	// Experimental.
	CollectCoverage *bool `json:"collectCoverage"`
	// An array of glob patterns indicating a set of files for which coverage information should be collected.
	// Experimental.
	CollectCoverageFrom *[]*string `json:"collectCoverageFrom"`
	// The directory where Jest should output its coverage files.
	// Experimental.
	CoverageDirectory *string `json:"coverageDirectory"`
	// An array of regexp pattern strings that are matched against all file paths before executing the test.
	//
	// If the file path matches any of the patterns, coverage information will be skipped
	// Experimental.
	CoveragePathIgnorePatterns *[]*string `json:"coveragePathIgnorePatterns"`
	// Indicates which provider should be used to instrument code for coverage.
	//
	// Allowed values are babel (default) or v8
	// Experimental.
	CoverageProvider *string `json:"coverageProvider"`
	// A list of reporter names that Jest uses when writing coverage reports.
	//
	// Any istanbul reporter can be used
	// Experimental.
	CoverageReporters *[]*string `json:"coverageReporters"`
	// Specify the global coverage thresholds.
	//
	// This will be used to configure minimum threshold enforcement
	// for coverage results. Thresholds can be specified as global, as a glob, and as a directory or file path.
	// If thresholds aren't met, jest will fail.
	// Experimental.
	CoverageThreshold *CoverageThreshold `json:"coverageThreshold"`
	// This option allows the use of a custom dependency extractor.
	//
	// It must be a node module that exports an object with an extract function
	// Experimental.
	DependencyExtractor *string `json:"dependencyExtractor"`
	// Allows for a label to be printed alongside a test while it is running.
	// Experimental.
	DisplayName interface{} `json:"displayName"`
	// Make calling deprecated APIs throw helpful error messages.
	//
	// Useful for easing the upgrade process.
	// Experimental.
	ErrorOnDeprecated *bool `json:"errorOnDeprecated"`
	// Test files run inside a vm, which slows calls to global context properties (e.g. Math). With this option you can specify extra properties to be defined inside the vm for faster lookups.
	// Experimental.
	ExtraGlobals *[]*string `json:"extraGlobals"`
	// Test files are normally ignored from collecting code coverage.
	//
	// With this option, you can overwrite this behavior and include otherwise ignored files in code coverage.
	// Experimental.
	ForceCoverageMatch *[]*string `json:"forceCoverageMatch"`
	// A set of global variables that need to be available in all test environments.
	// Experimental.
	Globals interface{} `json:"globals"`
	// This option allows the use of a custom global setup module which exports an async function that is triggered once before all test suites.
	//
	// This function gets Jest's globalConfig object as a parameter.
	// Experimental.
	GlobalSetup *string `json:"globalSetup"`
	// This option allows the use of a custom global teardown module which exports an async function that is triggered once after all test suites.
	//
	// This function gets Jest's globalConfig object as a parameter.
	// Experimental.
	GlobalTeardown *string `json:"globalTeardown"`
	// This will be used to configure the behavior of jest-haste-map, Jest's internal file crawler/cache system.
	// Experimental.
	Haste *HasteConfig `json:"haste"`
	// Insert Jest's globals (expect, test, describe, beforeEach etc.) into the global environment. If you set this to false, you should import from @jest/globals.
	// Experimental.
	InjectGlobals *bool `json:"injectGlobals"`
	// A number limiting the number of tests that are allowed to run at the same time when using test.concurrent. Any test above this limit will be queued and executed once a slot is released.
	// Experimental.
	MaxConcurrency *float64 `json:"maxConcurrency"`
	// An array of directory names to be searched recursively up from the requiring module's location.
	//
	// Setting this option will override the default, if you wish to still search node_modules for packages
	// include it along with any other options: ["node_modules", "bower_components"]
	// Experimental.
	ModuleDirectories *[]*string `json:"moduleDirectories"`
	// An array of file extensions your modules use.
	//
	// If you require modules without specifying a file extension,
	// these are the extensions Jest will look for, in left-to-right order.
	// Experimental.
	ModuleFileExtensions *[]*string `json:"moduleFileExtensions"`
	// A map from regular expressions to module names or to arrays of module names that allow to stub out resources, like images or styles with a single module.
	// Experimental.
	ModuleNameMapper *map[string]interface{} `json:"moduleNameMapper"`
	// An array of regexp pattern strings that are matched against all module paths before those paths are to be considered 'visible' to the module loader.
	//
	// If a given module's path matches any of the patterns,
	// it will not be require()-able in the test environment.
	// Experimental.
	ModulePathIgnorePatterns *[]*string `json:"modulePathIgnorePatterns"`
	// An alternative API to setting the NODE_PATH env variable, modulePaths is an array of absolute paths to additional locations to search when resolving modules.
	//
	// Use the <rootDir> string token to include
	// the path to your project's root directory. Example: ["<rootDir>/app/"].
	// Experimental.
	ModulePaths *[]*string `json:"modulePaths"`
	// Activates notifications for test results.
	// Experimental.
	Notify *bool `json:"notify"`
	// Specifies notification mode.
	//
	// Requires notify: true
	// Experimental.
	NotifyMode *string `json:"notifyMode"`
	// A preset that is used as a base for Jest's configuration.
	//
	// A preset should point to an npm module
	// that has a jest-preset.json or jest-preset.js file at the root.
	// Experimental.
	Preset *string `json:"preset"`
	// Sets the path to the prettier node module used to update inline snapshots.
	// Experimental.
	PrettierPath *string `json:"prettierPath"`
	// When the projects configuration is provided with an array of paths or glob patterns, Jest will run tests in all of the specified projects at the same time.
	//
	// This is great for monorepos or
	// when working on multiple projects at the same time.
	// Experimental.
	Projects *[]interface{} `json:"projects"`
	// Use this configuration option to add custom reporters to Jest.
	//
	// A custom reporter is a class
	// that implements onRunStart, onTestStart, onTestResult, onRunComplete methods that will be
	// called when any of those events occurs.
	// Experimental.
	Reporters *[]interface{} `json:"reporters"`
	// Automatically reset mock state before every test.
	//
	// Equivalent to calling jest.resetAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed but
	// does not restore their initial implementation.
	// Experimental.
	ResetMocks *bool `json:"resetMocks"`
	// By default, each test file gets its own independent module registry.
	//
	// Enabling resetModules
	// goes a step further and resets the module registry before running each individual test.
	// Experimental.
	ResetModules *bool `json:"resetModules"`
	// This option allows the use of a custom resolver.
	//
	// https://jestjs.io/docs/en/configuration#resolver-string
	// Experimental.
	Resolver *string `json:"resolver"`
	// Automatically restore mock state before every test.
	//
	// Equivalent to calling jest.restoreAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed and
	// restores their initial implementation.
	// Experimental.
	RestoreMocks *bool `json:"restoreMocks"`
	// The root directory that Jest should scan for tests and modules within.
	//
	// If you put your Jest
	// config inside your package.json and want the root directory to be the root of your repo, the
	// value for this config param will default to the directory of the package.json.
	// Experimental.
	RootDir *string `json:"rootDir"`
	// A list of paths to directories that Jest should use to search for files in.
	// Experimental.
	Roots *[]*string `json:"roots"`
	// This option allows you to use a custom runner instead of Jest's default test runner.
	// Experimental.
	Runner *string `json:"runner"`
	// A list of paths to modules that run some code to configure or set up the testing environment.
	//
	// Each setupFile will be run once per test file. Since every test runs in its own environment,
	// these scripts will be executed in the testing environment immediately before executing the
	// test code itself.
	// Experimental.
	SetupFiles *[]*string `json:"setupFiles"`
	// A list of paths to modules that run some code to configure or set up the testing framework before each test file in the suite is executed.
	//
	// Since setupFiles executes before the test
	// framework is installed in the environment, this script file presents you the opportunity of
	// running some code immediately after the test framework has been installed in the environment.
	// Experimental.
	SetupFilesAfterEnv *[]*string `json:"setupFilesAfterEnv"`
	// The number of seconds after which a test is considered as slow and reported as such in the results.
	// Experimental.
	SlowTestThreshold *float64 `json:"slowTestThreshold"`
	// The path to a module that can resolve test<->snapshot path.
	//
	// This config option lets you customize
	// where Jest stores snapshot files on disk.
	// Experimental.
	SnapshotResolver *string `json:"snapshotResolver"`
	// A list of paths to snapshot serializer modules Jest should use for snapshot testing.
	// Experimental.
	SnapshotSerializers *[]*string `json:"snapshotSerializers"`
	// The test environment that will be used for testing.
	//
	// The default environment in Jest is a
	// browser-like environment through jsdom. If you are building a node service, you can use the node
	// option to use a node-like environment instead.
	// Experimental.
	TestEnvironment *string `json:"testEnvironment"`
	// Test environment options that will be passed to the testEnvironment.
	//
	// The relevant options depend on the environment.
	// Experimental.
	TestEnvironmentOptions interface{} `json:"testEnvironmentOptions"`
	// The exit code Jest returns on test failure.
	// Experimental.
	TestFailureExitCode *float64 `json:"testFailureExitCode"`
	// The glob patterns Jest uses to detect test files.
	//
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Experimental.
	TestMatch *[]*string `json:"testMatch"`
	// An array of regexp pattern strings that are matched against all test paths before executing the test.
	//
	// If the test path matches any of the patterns, it will be skipped.
	// Experimental.
	TestPathIgnorePatterns *[]*string `json:"testPathIgnorePatterns"`
	// The pattern or patterns Jest uses to detect test files.
	//
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Experimental.
	TestRegex interface{} `json:"testRegex"`
	// This option allows the use of a custom results processor.
	// Experimental.
	TestResultsProcessor *string `json:"testResultsProcessor"`
	// This option allows the use of a custom test runner.
	//
	// The default is jasmine2. A custom test runner
	// can be provided by specifying a path to a test runner implementation.
	// Experimental.
	TestRunner *string `json:"testRunner"`
	// This option allows you to use a custom sequencer instead of Jest's default.
	//
	// Sort may optionally return a Promise.
	// Experimental.
	TestSequencer *string `json:"testSequencer"`
	// Default timeout of a test in milliseconds.
	// Experimental.
	TestTimeout *float64 `json:"testTimeout"`
	// This option sets the URL for the jsdom environment.
	//
	// It is reflected in properties such as location.href.
	// Experimental.
	TestURL *string `json:"testURL"`
	// Setting this value to legacy or fake allows the use of fake timers for functions such as setTimeout.
	//
	// Fake timers are useful when a piece of code sets a long timeout that we don't want to wait for in a test.
	// Experimental.
	Timers *string `json:"timers"`
	// A map from regular expressions to paths to transformers.
	//
	// A transformer is a module that provides a
	// synchronous function for transforming source files.
	// Experimental.
	Transform *map[string]interface{} `json:"transform"`
	// An array of regexp pattern strings that are matched against all source file paths before transformation.
	//
	// If the test path matches any of the patterns, it will not be transformed.
	// Experimental.
	TransformIgnorePatterns *[]*string `json:"transformIgnorePatterns"`
	// An array of regexp pattern strings that are matched against all modules before the module loader will automatically return a mock for them.
	//
	// If a module's path matches any of the patterns in this list, it
	// will not be automatically mocked by the module loader.
	// Experimental.
	UnmockedModulePathPatterns *[]*string `json:"unmockedModulePathPatterns"`
	// Indicates whether each individual test should be reported during the run.
	//
	// All errors will also
	// still be shown on the bottom after execution. Note that if there is only one test file being run
	// it will default to true.
	// Experimental.
	Verbose *bool `json:"verbose"`
	// Whether to use watchman for file crawling.
	// Experimental.
	Watchman *bool `json:"watchman"`
	// An array of RegExp patterns that are matched against all source file paths before re-running tests in watch mode.
	//
	// If the file path matches any of the patterns, when it is updated, it will not trigger
	// a re-run of tests.
	// Experimental.
	WatchPathIgnorePatterns *[]*string `json:"watchPathIgnorePatterns"`
	// Experimental.
	WatchPlugins *map[string]interface{} `json:"watchPlugins"`
}

// Experimental.
type JestOptions struct {
	// Path to JSON config file for Jest.
	// Experimental.
	ConfigFilePath *string `json:"configFilePath"`
	// Collect coverage.
	//
	// Deprecated
	// Deprecated: use jestConfig.collectCoverage
	Coverage *bool `json:"coverage"`
	// Include the `text` coverage reporter, which means that coverage summary is printed at the end of the jest execution.
	// Experimental.
	CoverageText *bool `json:"coverageText"`
	// Defines `testPathIgnorePatterns` and `coveragePathIgnorePatterns`.
	// Deprecated: use jestConfig.coveragePathIgnorePatterns or jestConfig.testPathIgnorePatterns respectively
	IgnorePatterns *[]*string `json:"ignorePatterns"`
	// Jest configuration.
	// Experimental.
	JestConfig *JestConfigOptions `json:"jestConfig"`
	// The version of jest to use.
	// Experimental.
	JestVersion *string `json:"jestVersion"`
	// Result processing with jest-junit.
	//
	// Output directory is `test-reports/`.
	// Experimental.
	JunitReporting *bool `json:"junitReporting"`
	// Preserve the default Jest reporter when additional reporters are added.
	// Experimental.
	PreserveDefaultReporters *bool `json:"preserveDefaultReporters"`
}

// Represents the npm `package.json` file.
// Experimental.
type NodePackage interface {
	projen.Component
	AllowLibraryDependencies() *bool
	CodeArtifactOptions() *CodeArtifactOptions
	Entrypoint() *string
	InstallAndUpdateLockfileCommand() *string
	InstallCommand() *string
	License() *string
	LockFile() *string
	Manifest() interface{}
	MaxNodeVersion() *string
	MinNodeVersion() *string
	NpmAccess() NpmAccess
	NpmRegistry() *string
	NpmRegistryUrl() *string
	NpmTokenSecret() *string
	PackageManager() NodePackageManager
	PackageName() *string
	Project() projen.Project
	ProjenCommand() *string
	AddBin(bins *map[string]*string)
	AddBundledDeps(deps ...*string)
	AddDeps(deps ...*string)
	AddDevDeps(deps ...*string)
	AddEngine(engine *string, version *string)
	AddField(name *string, value interface{})
	AddKeywords(keywords ...*string)
	AddPeerDeps(deps ...*string)
	AddVersion(version *string)
	HasScript(name *string) *bool
	PostSynthesize()
	PreSynthesize()
	RemoveScript(name *string)
	RenderUpgradePackagesCommand(exclude *[]*string, include *[]*string) *string
	SetScript(name *string, command *string)
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

// Experimental.
func (n *jsiiProxy_NodePackage) AddBin(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		n,
		"addBin",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
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

// Defines normal dependencies.
// Experimental.
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

// Defines development/test dependencies.
// Experimental.
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

// Adds an `engines` requirement to your package.
// Experimental.
func (n *jsiiProxy_NodePackage) AddEngine(engine *string, version *string) {
	_jsii_.InvokeVoid(
		n,
		"addEngine",
		[]interface{}{engine, version},
	)
}

// Directly set fields in `package.json`.
// Experimental.
func (n *jsiiProxy_NodePackage) AddField(name *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addField",
		[]interface{}{name, value},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
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

// Defines peer dependencies.
//
// When adding peer dependencies, a devDependency will also be added on the
// pinned version of the declared peer. This will ensure that you are testing
// your code against the minimum version required from your consumers.
// Experimental.
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

// Sets the package version.
// Experimental.
func (n *jsiiProxy_NodePackage) AddVersion(version *string) {
	_jsii_.InvokeVoid(
		n,
		"addVersion",
		[]interface{}{version},
	)
}

// Indicates if a script by the given name is defined.
// Deprecated: Use `project.tasks.tryFind(name)`
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (n *jsiiProxy_NodePackage) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (n *jsiiProxy_NodePackage) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (n *jsiiProxy_NodePackage) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		n,
		"removeScript",
		[]interface{}{name},
	)
}

// Render a package manager specific command to upgrade all requested dependencies.
// Experimental.
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

// Override the contents of an npm package.json script.
// Experimental.
func (n *jsiiProxy_NodePackage) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		n,
		"setScript",
		[]interface{}{name, command},
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	NodePackageManager_YARN NodePackageManager = "YARN"
	NodePackageManager_NPM NodePackageManager = "NPM"
	NodePackageManager_PNPM NodePackageManager = "PNPM"
)

// Experimental.
type NodePackageOptions struct {
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization *bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `json:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `json:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
	BugsUrl *string `json:"bugsUrl"`
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
	BundledDeps *[]*string `json:"bundledDeps"`
	// Options for publishing npm package to AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *CodeArtifactOptions `json:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Deps *[]*string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description *string `json:"description"`
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
	// TODO: EXAMPLE
	//
	// Experimental.
	DevDeps *[]*string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `json:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry *string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `json:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `json:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions *PeerDependencyOptions `json:"peerDependencyOptions"`
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
	PeerDeps *[]*string `json:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `json:"repositoryDirectory"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts *map[string]*string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `json:"stability"`
}

// Node.js project.
// Experimental.
type NodeProject interface {
	github.GitHubProject
	AllowLibraryDependencies() *bool
	Antitamper() *bool
	ArtifactsDirectory() *string
	ArtifactsJavascriptDirectory() *string
	AutoApprove() github.AutoApprove
	AutoMerge() github.AutoMerge
	BuildTask() projen.Task
	BuildWorkflow() build.BuildWorkflow
	BuildWorkflowJobId() *string
	Bundler() Bundler
	CompileTask() projen.Task
	Components() *[]projen.Component
	DefaultTask() projen.Task
	Deps() projen.Dependencies
	DevContainer() vscode.DevContainer
	Entrypoint() *string
	Files() *[]projen.FileBase
	Gitattributes() projen.GitAttributesFile
	Github() github.GitHub
	Gitignore() projen.IgnoreFile
	Gitpod() projen.Gitpod
	InitProject() *projen.InitProject
	InstallWorkflowSteps() *[]*workflows.JobStep
	Jest() Jest
	Logger() projen.Logger
	Manifest() interface{}
	MaxNodeVersion() *string
	MinNodeVersion() *string
	Name() *string
	Npmignore() projen.IgnoreFile
	Outdir() *string
	Package() NodePackage
	PackageManager() NodePackageManager
	PackageTask() projen.Task
	Parent() projen.Project
	PostCompileTask() projen.Task
	PreCompileTask() projen.Task
	ProjectBuild() projen.ProjectBuild
	ProjectType() projen.ProjectType
	ProjenCommand() *string
	Publisher() release.Publisher
	Release() release.Release
	Root() projen.Project
	RunScriptCommand() *string
	Tasks() projen.Tasks
	TestTask() projen.Task
	UpgradeWorkflow() UpgradeDependencies
	Vscode() vscode.VsCode
	AddBins(bins *map[string]*string)
	AddBundledDeps(deps ...*string)
	AddCompileCommand(commands ...*string)
	AddDeps(deps ...*string)
	AddDevDeps(deps ...*string)
	AddExcludeFromCleanup(globs ...*string)
	AddFields(fields *map[string]interface{})
	AddGitIgnore(pattern *string)
	AddKeywords(keywords ...*string)
	AddPackageIgnore(pattern *string)
	AddPeerDeps(deps ...*string)
	AddTask(name *string, props *projen.TaskOptions) projen.Task
	AddTestCommand(commands ...*string)
	AddTip(message *string)
	AnnotateGenerated(glob *string)
	HasScript(name *string) *bool
	PostSynthesize()
	PreSynthesize()
	RemoveScript(name *string)
	RemoveTask(name *string) projen.Task
	RunTaskCommand(task projen.Task) *string
	SetScript(name *string, command *string)
	Synth()
	TryFindFile(filePath *string) projen.FileBase
	TryFindJsonFile(filePath *string) projen.JsonFile
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

func (j *jsiiProxy_NodeProject) Antitamper() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"antitamper",
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

func (j *jsiiProxy_NodeProject) InstallWorkflowSteps() *[]*workflows.JobStep {
	var returns *[]*workflows.JobStep
	_jsii_.Get(
		j,
		"installWorkflowSteps",
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

// Experimental.
func (n *jsiiProxy_NodeProject) AddBins(bins *map[string]*string) {
	_jsii_.InvokeVoid(
		n,
		"addBins",
		[]interface{}{bins},
	)
}

// Defines bundled dependencies.
//
// Bundled dependencies will be added as normal dependencies as well as to the
// `bundledDependencies` section of your `package.json`.
// Experimental.
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

// DEPRECATED.
// Deprecated: use `project.compileTask.exec()`
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

// Defines normal dependencies.
// Experimental.
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

// Defines development/test dependencies.
// Experimental.
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

// Exclude the matching files from pre-synth cleanup.
//
// Can be used when, for example, some
// source files include the projen marker and we don't want them to be erased during synth.
// Experimental.
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

// Directly set fields in `package.json`.
// Experimental.
func (n *jsiiProxy_NodeProject) AddFields(fields *map[string]interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addFields",
		[]interface{}{fields},
	)
}

// Adds a .gitignore pattern.
// Experimental.
func (n *jsiiProxy_NodeProject) AddGitIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
		"addGitIgnore",
		[]interface{}{pattern},
	)
}

// Adds keywords to package.json (deduplicated).
// Experimental.
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

// Exclude these files from the bundled package.
//
// Implemented by project types based on the
// packaging mechanism. For example, `NodeProject` delegates this to `.npmignore`.
// Experimental.
func (n *jsiiProxy_NodeProject) AddPackageIgnore(pattern *string) {
	_jsii_.InvokeVoid(
		n,
		"addPackageIgnore",
		[]interface{}{pattern},
	)
}

// Defines peer dependencies.
//
// When adding peer dependencies, a devDependency will also be added on the
// pinned version of the declared peer. This will ensure that you are testing
// your code against the minimum version required from your consumers.
// Experimental.
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

// Adds a new task to this project.
//
// This will fail if the project already has
// a task with this name.
// Experimental.
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

// DEPRECATED.
// Deprecated: use `project.testTask.exec()`
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

// Prints a "tip" message during synthesis.
// Deprecated: - use `project.logger.info(message)` to show messages during synthesis
func (n *jsiiProxy_NodeProject) AddTip(message *string) {
	_jsii_.InvokeVoid(
		n,
		"addTip",
		[]interface{}{message},
	)
}

// Marks the provided file(s) as being generated.
//
// This is achieved using the
// github-linguist attributes. Generated files do not count against the
// repository statistics and language breakdown.
// See: https://github.com/github/linguist/blob/master/docs/overrides.md
//
// Deprecated.
func (n *jsiiProxy_NodeProject) AnnotateGenerated(glob *string) {
	_jsii_.InvokeVoid(
		n,
		"annotateGenerated",
		[]interface{}{glob},
	)
}

// Indicates if a script by the name name is defined.
// Experimental.
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

// Called after all components are synthesized.
//
// Order is *not* guaranteed.
// Experimental.
func (n *jsiiProxy_NodeProject) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before all components are synthesized.
// Experimental.
func (n *jsiiProxy_NodeProject) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

// Removes the npm script (always successful).
// Experimental.
func (n *jsiiProxy_NodeProject) RemoveScript(name *string) {
	_jsii_.InvokeVoid(
		n,
		"removeScript",
		[]interface{}{name},
	)
}

// Removes a task from a project.
//
// Returns: The `Task` that was removed, otherwise `undefined`.
// Experimental.
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

// Returns the shell command to execute in order to run a task.
//
// This will
// typically be `npx projen TASK`.
// Experimental.
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

// Replaces the contents of an npm package.json script.
// Experimental.
func (n *jsiiProxy_NodeProject) SetScript(name *string, command *string) {
	_jsii_.InvokeVoid(
		n,
		"setScript",
		[]interface{}{name, command},
	)
}

// Synthesize all project files into `outdir`.
//
// 1. Call "this.preSynthesize()"
// 2. Delete all generated files
// 3. Synthesize all sub-projects
// 4. Synthesize all components of this project
// 5. Call "postSynthesize()" for all components of this project
// 6. Call "this.postSynthesize()"
// Experimental.
func (n *jsiiProxy_NodeProject) Synth() {
	_jsii_.InvokeVoid(
		n,
		"synth",
		nil, // no parameters
	)
}

// Finds a file at the specified relative path within this project and all its subprojects.
//
// Returns: a `FileBase` or undefined if there is no file in that path
// Experimental.
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

// Finds a json file by name.
// Deprecated: use `tryFindObjectFile`
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

// Finds an object file (like JsonFile, YamlFile, etc.) by name.
// Experimental.
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
	Name *string `json:"name"`
	// Configure logging options such as verbosity.
	// Experimental.
	Logging *projen.LoggerOptions `json:"logging"`
	// The root directory of the project.
	//
	// Relative to this directory, all files are synthesized.
	//
	// If this project has a parent, this directory is relative to the parent
	// directory and it cannot be the same as the parent or any of it's other
	// sub-projects.
	// Experimental.
	Outdir *string `json:"outdir"`
	// The parent project, if this project is part of a bigger project.
	// Experimental.
	Parent projen.Project `json:"parent"`
	// The shell command to use in order to run the projen CLI.
	//
	// Can be used to customize in special environments.
	// Experimental.
	ProjenCommand *string `json:"projenCommand"`
	// Generate (once) .projenrc.json (in JSON). Set to `false` in order to disable .projenrc.json generation.
	// Experimental.
	ProjenrcJson *bool `json:"projenrcJson"`
	// Options for .projenrc.json.
	// Experimental.
	ProjenrcJsonOptions *projen.ProjenrcOptions `json:"projenrcJsonOptions"`
	// Enable and configure the 'auto approve' workflow.
	// Experimental.
	AutoApproveOptions *github.AutoApproveOptions `json:"autoApproveOptions"`
	// Configure options for automatic merging on GitHub.
	//
	// Has no effect if
	// `github.mergify` is set to false.
	// Experimental.
	AutoMergeOptions *github.AutoMergeOptions `json:"autoMergeOptions"`
	// Add a `clobber` task which resets the repo to origin.
	// Experimental.
	Clobber *bool `json:"clobber"`
	// Add a VSCode development environment (used for GitHub Codespaces).
	// Experimental.
	DevContainer *bool `json:"devContainer"`
	// Enable GitHub integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Github *bool `json:"github"`
	// Options for GitHub integration.
	// Experimental.
	GithubOptions *github.GitHubOptions `json:"githubOptions"`
	// Add a Gitpod development environment.
	// Experimental.
	Gitpod *bool `json:"gitpod"`
	// Whether mergify should be enabled on this repository or not.
	// Deprecated: use `githubOptions.mergify` instead
	Mergify *bool `json:"mergify"`
	// Options for mergify.
	// Deprecated: use `githubOptions.mergifyOptions` instead
	MergifyOptions *github.MergifyOptions `json:"mergifyOptions"`
	// Which type of project this is (library/app).
	// Deprecated: no longer supported at the base project level
	ProjectType projen.ProjectType `json:"projectType"`
	// The README setup.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Readme *projen.SampleReadmeProps `json:"readme"`
	// Auto-close of stale issues and pull request.
	//
	// See `staleOptions` for options.
	// Experimental.
	Stale *bool `json:"stale"`
	// Auto-close stale issues and pull requests.
	//
	// To disable set `stale` to `false`.
	// Experimental.
	StaleOptions *github.StaleOptions `json:"staleOptions"`
	// Enable VSCode integration.
	//
	// Enabled by default for root projects. Disabled for non-root projects.
	// Experimental.
	Vscode *bool `json:"vscode"`
	// Allow the project to include `peerDependencies` and `bundledDependencies`.
	//
	// This is normally only allowed for libraries. For apps, there's no meaning
	// for specifying these.
	// Experimental.
	AllowLibraryDependencies *bool `json:"allowLibraryDependencies"`
	// Author's e-mail.
	// Experimental.
	AuthorEmail *string `json:"authorEmail"`
	// Author's name.
	// Experimental.
	AuthorName *string `json:"authorName"`
	// Author's Organization.
	// Experimental.
	AuthorOrganization *bool `json:"authorOrganization"`
	// Author's URL / Website.
	// Experimental.
	AuthorUrl *string `json:"authorUrl"`
	// Automatically add all executables under the `bin` directory to your `package.json` file under the `bin` section.
	// Experimental.
	AutoDetectBin *bool `json:"autoDetectBin"`
	// Binary programs vended with your module.
	//
	// You can use this option to add/customize how binaries are represented in
	// your `package.json`, but unless `autoDetectBin` is `false`, every
	// executable file under `bin` will automatically be added to this section.
	// Experimental.
	Bin *map[string]*string `json:"bin"`
	// The email address to which issues should be reported.
	// Experimental.
	BugsEmail *string `json:"bugsEmail"`
	// The url to your project's issue tracker.
	// Experimental.
	BugsUrl *string `json:"bugsUrl"`
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
	BundledDeps *[]*string `json:"bundledDeps"`
	// Options for publishing npm package to AWS CodeArtifact.
	// Experimental.
	CodeArtifactOptions *CodeArtifactOptions `json:"codeArtifactOptions"`
	// Runtime dependencies of this module.
	//
	// The recommendation is to only specify the module name here (e.g.
	// `express`). This will behave similar to `yarn add` or `npm install` in the
	// sense that it will add the module as a dependency to your `package.json`
	// file with the latest version (`^`). You can specify semver requirements in
	// the same syntax passed to `npm i` or `yarn add` (e.g. `express@^2`) and
	// this will be what you `package.json` will eventually include.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Deps *[]*string `json:"deps"`
	// The description is just a string that helps people understand the purpose of the package.
	//
	// It can be used when searching for packages in a package manager as well.
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-description
	// Experimental.
	Description *string `json:"description"`
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
	// TODO: EXAMPLE
	//
	// Experimental.
	DevDeps *[]*string `json:"devDeps"`
	// Module entrypoint (`main` in `package.json`).
	//
	// Set to an empty string to not include `main` in your package.json
	// Experimental.
	Entrypoint *string `json:"entrypoint"`
	// Package's Homepage / Website.
	// Experimental.
	Homepage *string `json:"homepage"`
	// Keywords to include in `package.json`.
	// Experimental.
	Keywords *[]*string `json:"keywords"`
	// License's SPDX identifier.
	//
	// See https://github.com/projen/projen/tree/main/license-text for a list of supported licenses.
	// Use the `licensed` option if you want to no license to be specified.
	// Experimental.
	License *string `json:"license"`
	// Indicates if a license should be added.
	// Experimental.
	Licensed *bool `json:"licensed"`
	// Minimum node.js version to require via `engines` (inclusive).
	// Experimental.
	MaxNodeVersion *string `json:"maxNodeVersion"`
	// Minimum Node.js version to require via package.json `engines` (inclusive).
	// Experimental.
	MinNodeVersion *string `json:"minNodeVersion"`
	// Access level of the npm package.
	// Experimental.
	NpmAccess NpmAccess `json:"npmAccess"`
	// The host name of the npm registry to publish to.
	//
	// Cannot be set together with `npmRegistryUrl`.
	// Deprecated: use `npmRegistryUrl` instead
	NpmRegistry *string `json:"npmRegistry"`
	// The base URL of the npm package registry.
	//
	// Must be a URL (e.g. start with "https://" or "http://")
	// Experimental.
	NpmRegistryUrl *string `json:"npmRegistryUrl"`
	// GitHub secret which contains the NPM token to use when publishing packages.
	// Experimental.
	NpmTokenSecret *string `json:"npmTokenSecret"`
	// The Node Package Manager used to execute scripts.
	// Experimental.
	PackageManager NodePackageManager `json:"packageManager"`
	// The "name" in package.json.
	// Experimental.
	PackageName *string `json:"packageName"`
	// Options for `peerDeps`.
	// Experimental.
	PeerDependencyOptions *PeerDependencyOptions `json:"peerDependencyOptions"`
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
	PeerDeps *[]*string `json:"peerDeps"`
	// The repository is the location where the actual code for your package lives.
	//
	// See https://classic.yarnpkg.com/en/docs/package-json/#toc-repository
	// Experimental.
	Repository *string `json:"repository"`
	// If the package.json for your package is not in the root directory (for example if it is part of a monorepo), you can specify the directory in which it lives.
	// Experimental.
	RepositoryDirectory *string `json:"repositoryDirectory"`
	// npm scripts to include.
	//
	// If a script has the same name as a standard script,
	// the standard script will be overwritten.
	// Experimental.
	Scripts *map[string]*string `json:"scripts"`
	// Package's Stability.
	// Experimental.
	Stability *string `json:"stability"`
	// Checks that after build there are no modified files on git.
	// Experimental.
	Antitamper *bool `json:"antitamper"`
	// Version requirement of `jsii-release` which is used to publish modules to npm.
	// Experimental.
	JsiiReleaseVersion *string `json:"jsiiReleaseVersion"`
	// Major version to release from the default branch.
	//
	// If this is specified, we bump the latest version of this major version line.
	// If not specified, we bump the global latest version.
	// Experimental.
	MajorVersion *float64 `json:"majorVersion"`
	// The npmDistTag to use when publishing from the default branch.
	//
	// To set the npm dist-tag for release branches, set the `npmDistTag` property
	// for each branch.
	// Experimental.
	NpmDistTag *string `json:"npmDistTag"`
	// Steps to execute after build as part of the release workflow.
	// Experimental.
	PostBuildSteps *[]*workflows.JobStep `json:"postBuildSteps"`
	// Bump versions from the default branch as pre-releases (e.g. "beta", "alpha", "pre").
	// Experimental.
	Prerelease *string `json:"prerelease"`
	// Instead of actually publishing to package managers, just print the publishing command.
	// Experimental.
	PublishDryRun *bool `json:"publishDryRun"`
	// Define publishing tasks that can be executed manually as well as workflows.
	//
	// Normally, publishing only happens within automated workflows. Enable this
	// in order to create a publishing task for each publishing activity.
	// Experimental.
	PublishTasks *bool `json:"publishTasks"`
	// Defines additional release branches.
	//
	// A workflow will be created for each
	// release branch which will publish releases from commits in this branch.
	// Each release branch _must_ be assigned a major version number which is used
	// to enforce that versions published from that branch always use that major
	// version. If multiple branches are used, the `majorVersion` field must also
	// be provided for the default branch.
	// Experimental.
	ReleaseBranches *map[string]*release.BranchOptions `json:"releaseBranches"`
	// Automatically release new versions every commit to one of branches in `releaseBranches`.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.continuous()` instead
	ReleaseEveryCommit *bool `json:"releaseEveryCommit"`
	// Create a github issue on every failed publishing task.
	// Experimental.
	ReleaseFailureIssue *bool `json:"releaseFailureIssue"`
	// The label to apply to issues indicating publish failures.
	//
	// Only applies if `releaseFailureIssue` is true.
	// Experimental.
	ReleaseFailureIssueLabel *string `json:"releaseFailureIssueLabel"`
	// CRON schedule to trigger new releases.
	// Deprecated: Use `releaseTrigger: ReleaseTrigger.scheduled()` instead
	ReleaseSchedule *string `json:"releaseSchedule"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	ReleaseTagPrefix *string `json:"releaseTagPrefix"`
	// The release trigger to use.
	// Experimental.
	ReleaseTrigger release.ReleaseTrigger `json:"releaseTrigger"`
	// The name of the default release workflow.
	// Experimental.
	ReleaseWorkflowName *string `json:"releaseWorkflowName"`
	// A set of workflow steps to execute in order to setup the workflow container.
	// Experimental.
	ReleaseWorkflowSetupSteps *[]*workflows.JobStep `json:"releaseWorkflowSetupSteps"`
	// Custom configuration used when creating changelog with standard-version package.
	//
	// Given values either append to default configuration or overwrite values in it.
	// Experimental.
	VersionrcOptions *map[string]interface{} `json:"versionrcOptions"`
	// Container image to use for GitHub workflows.
	// Experimental.
	WorkflowContainerImage *string `json:"workflowContainerImage"`
	// Github Runner selection labels.
	// Experimental.
	WorkflowRunsOn *[]*string `json:"workflowRunsOn"`
	// The name of the main release branch.
	// Experimental.
	DefaultReleaseBranch *string `json:"defaultReleaseBranch"`
	// A directory which will contain build artifacts.
	// Experimental.
	ArtifactsDirectory *string `json:"artifactsDirectory"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveProjenUpgrades *bool `json:"autoApproveProjenUpgrades"`
	// Automatically approve deps upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Experimental.
	AutoApproveUpgrades *bool `json:"autoApproveUpgrades"`
	// Define a GitHub workflow for building PRs.
	// Experimental.
	BuildWorkflow *bool `json:"buildWorkflow"`
	// Options for `Bundler`.
	// Experimental.
	BundlerOptions *BundlerOptions `json:"bundlerOptions"`
	// Define a GitHub workflow step for sending code coverage metrics to https://codecov.io/ Uses codecov/codecov-action@v1 A secret is required for private repos. Configured with @codeCovTokenSecret.
	// Experimental.
	CodeCov *bool `json:"codeCov"`
	// Define the secret name for a specified https://codecov.io/ token A secret is required to send coverage for private repositories.
	// Experimental.
	CodeCovTokenSecret *string `json:"codeCovTokenSecret"`
	// License copyright owner.
	// Experimental.
	CopyrightOwner *string `json:"copyrightOwner"`
	// The copyright years to put in the LICENSE file.
	// Experimental.
	CopyrightPeriod *string `json:"copyrightPeriod"`
	// Use dependabot to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `depsUpgrade`.
	// Experimental.
	Dependabot *bool `json:"dependabot"`
	// Options for dependabot.
	// Experimental.
	DependabotOptions *github.DependabotOptions `json:"dependabotOptions"`
	// Use github workflows to handle dependency upgrades.
	//
	// Cannot be used in conjunction with `dependabot`.
	// Experimental.
	DepsUpgrade *bool `json:"depsUpgrade"`
	// Options for depsUpgrade.
	// Experimental.
	DepsUpgradeOptions *UpgradeDependenciesOptions `json:"depsUpgradeOptions"`
	// Additional entries to .gitignore.
	// Experimental.
	Gitignore *[]*string `json:"gitignore"`
	// Setup jest unit tests.
	// Experimental.
	Jest *bool `json:"jest"`
	// Jest options.
	// Experimental.
	JestOptions *JestOptions `json:"jestOptions"`
	// Automatically update files modified during builds to pull-request branches.
	//
	// This means
	// that any files synthesized by projen or e.g. test snapshots will always be up-to-date
	// before a PR is merged.
	//
	// Implies that PR builds do not have anti-tamper checks.
	// Experimental.
	MutableBuild *bool `json:"mutableBuild"`
	// Additional entries to .npmignore.
	// Deprecated: - use `project.addPackageIgnore`
	Npmignore *[]*string `json:"npmignore"`
	// Defines an .npmignore file. Normally this is only needed for libraries that are packaged as tarballs.
	// Experimental.
	NpmignoreEnabled *bool `json:"npmignoreEnabled"`
	// Defines a `package` task that will produce an npm tarball under the artifacts directory (e.g. `dist`).
	// Experimental.
	Package *bool `json:"package"`
	// Indicates of "projen" should be installed as a devDependency.
	// Experimental.
	ProjenDevDependency *bool `json:"projenDevDependency"`
	// Generate (once) .projenrc.js (in JavaScript). Set to `false` in order to disable .projenrc.js generation.
	// Experimental.
	ProjenrcJs *bool `json:"projenrcJs"`
	// Options for .projenrc.js.
	// Experimental.
	ProjenrcJsOptions *ProjenrcOptions `json:"projenrcJsOptions"`
	// Automatically approve projen upgrade PRs, allowing them to be merged by mergify (if configued).
	//
	// Throw if set to true but `autoApproveOptions` are not defined.
	// Deprecated: use `autoApproveProjenUpgrades`.
	ProjenUpgradeAutoMerge *bool `json:"projenUpgradeAutoMerge"`
	// Customize the projenUpgrade schedule in cron expression.
	// Experimental.
	ProjenUpgradeSchedule *[]*string `json:"projenUpgradeSchedule"`
	// Periodically submits a pull request for projen upgrades (executes `yarn projen:upgrade`).
	//
	// This setting is a GitHub secret name which contains a GitHub Access Token
	// with `repo` and `workflow` permissions.
	//
	// This token is used to submit the upgrade pull request, which will likely
	// include workflow updates.
	//
	// To create a personal access token see https://github.com/settings/tokens
	// Experimental.
	ProjenUpgradeSecret *string `json:"projenUpgradeSecret"`
	// Version of projen to install.
	// Experimental.
	ProjenVersion *string `json:"projenVersion"`
	// Include a GitHub pull request template.
	// Experimental.
	PullRequestTemplate *bool `json:"pullRequestTemplate"`
	// The contents of the pull request template.
	// Experimental.
	PullRequestTemplateContents *[]*string `json:"pullRequestTemplateContents"`
	// Add release management to this project.
	// Experimental.
	Release *bool `json:"release"`
	// Automatically release to npm when new versions are introduced.
	// Experimental.
	ReleaseToNpm *bool `json:"releaseToNpm"`
	// DEPRECATED: renamed to `release`.
	// Deprecated: see `release`.
	ReleaseWorkflow *bool `json:"releaseWorkflow"`
	// Workflow steps to use in order to bootstrap this repo.
	// Experimental.
	WorkflowBootstrapSteps *[]interface{} `json:"workflowBootstrapSteps"`
	// The git identity to use in workflows.
	// Experimental.
	WorkflowGitIdentity *github.GitIdentity `json:"workflowGitIdentity"`
	// The node version to use in GitHub workflows.
	// Experimental.
	WorkflowNodeVersion *string `json:"workflowNodeVersion"`
}

// Experimental.
type NodeWorkflowSteps struct {
	// Experimental.
	Antitamper *[]interface{} `json:"antitamper"`
	// Experimental.
	Install *[]interface{} `json:"install"`
}

// Npm package access level.
// Experimental.
type NpmAccess string

const (
	NpmAccess_PUBLIC NpmAccess = "PUBLIC"
	NpmAccess_RESTRICTED NpmAccess = "RESTRICTED"
)

// File representing the local NPM config in .npmrc.
// Experimental.
type NpmConfig interface {
	projen.Component
	Project() projen.Project
	AddConfig(name *string, value *string)
	AddRegistry(url *string, scope *string)
	PostSynthesize()
	PreSynthesize()
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

// configure a generic property.
// Experimental.
func (n *jsiiProxy_NpmConfig) AddConfig(name *string, value *string) {
	_jsii_.InvokeVoid(
		n,
		"addConfig",
		[]interface{}{name, value},
	)
}

// configure a scoped registry.
// Experimental.
func (n *jsiiProxy_NpmConfig) AddRegistry(url *string, scope *string) {
	_jsii_.InvokeVoid(
		n,
		"addRegistry",
		[]interface{}{url, scope},
	)
}

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (n *jsiiProxy_NpmConfig) PostSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (n *jsiiProxy_NpmConfig) PreSynthesize() {
	_jsii_.InvokeVoid(
		n,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	// You can change this or add scoped registries using the addRegistry method
	// Experimental.
	Registry *string `json:"registry"`
}

// Experimental.
type PeerDependencyOptions struct {
	// Automatically add a pinned dev dependency.
	// Experimental.
	PinnedDevDependency *bool `json:"pinnedDevDependency"`
}

// Sets up a javascript project to use TypeScript for projenrc.
// Experimental.
type Projenrc interface {
	projen.Component
	Project() projen.Project
	PostSynthesize()
	PreSynthesize()
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (p *jsiiProxy_Projenrc) PostSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (p *jsiiProxy_Projenrc) PreSynthesize() {
	_jsii_.InvokeVoid(
		p,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Filename *string `json:"filename"`
}

// Experimental.
type TypeScriptCompilerOptions struct {
	// Allow JavaScript files to be compiled.
	// Experimental.
	AllowJs *bool `json:"allowJs"`
	// Allow default imports from modules with no default export.
	//
	// This does not affect code emit, just typechecking.
	// Experimental.
	AllowSyntheticDefaultImports *bool `json:"allowSyntheticDefaultImports"`
	// Ensures that your files are parsed in the ECMAScript strict mode, and emit “use strict” for each source file.
	// Experimental.
	AlwaysStrict *bool `json:"alwaysStrict"`
	// Lets you set a base directory to resolve non-absolute module names.
	//
	// You can define a root folder where you can do absolute file resolution.
	// Experimental.
	BaseUrl *string `json:"baseUrl"`
	// To be specified along with the above.
	// Experimental.
	Declaration *bool `json:"declaration"`
	// Offers a way to configure the root directory for where declaration files are emitted.
	// Experimental.
	DeclarationDir *string `json:"declarationDir"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	//
	// Decorators are a language feature which hasn’t yet been fully ratified into the JavaScript specification.
	// This means that the implementation version in TypeScript may differ from the implementation in JavaScript when it it decided by TC39.
	// You can find out more about decorator support in TypeScript in the handbook.
	// See: https://www.typescriptlang.org/docs/handbook/decorators.html
	//
	// Experimental.
	EmitDecoratorMetadata *bool `json:"emitDecoratorMetadata"`
	// Emit __importStar and __importDefault helpers for runtime babel ecosystem compatibility and enable --allowSyntheticDefaultImports for typesystem compatibility.
	// Experimental.
	EsModuleInterop *bool `json:"esModuleInterop"`
	// Enables experimental support for decorators, which is in stage 2 of the TC39 standardization process.
	// Experimental.
	ExperimentalDecorators *bool `json:"experimentalDecorators"`
	// Disallow inconsistently-cased references to the same file.
	// Experimental.
	ForceConsistentCasingInFileNames *bool `json:"forceConsistentCasingInFileNames"`
	// When set, instead of writing out a .js.map file to provide source maps, TypeScript will embed the source map content in the .js files.
	// Experimental.
	InlineSourceMap *bool `json:"inlineSourceMap"`
	// When set, TypeScript will include the original content of the .ts file as an embedded string in the source map. This is often useful in the same cases as inlineSourceMap.
	// Experimental.
	InlineSources *bool `json:"inlineSources"`
	// Perform additional checks to ensure that separate compilation (such as with transpileModule or @babel/plugin-transform-typescript) would be safe.
	// Experimental.
	IsolatedModules *bool `json:"isolatedModules"`
	// Support JSX in .tsx files: "react", "preserve", "react-native" etc.
	// Experimental.
	Jsx TypeScriptJsxMode `json:"jsx"`
	// Reference for type definitions / libraries to use (eg.
	//
	// ES2016, ES5, ES2018).
	// Experimental.
	Lib *[]*string `json:"lib"`
	// Sets the module system for the program.
	//
	// See https://www.typescriptlang.org/docs/handbook/modules.html#ambient-modules.
	// Experimental.
	Module *string `json:"module"`
	// Determine how modules get resolved.
	//
	// Either "Node" for Node.js/io.js style resolution, or "Classic".
	// Experimental.
	ModuleResolution TypeScriptModuleResolution `json:"moduleResolution"`
	// Do not emit outputs.
	// Experimental.
	NoEmit *bool `json:"noEmit"`
	// Do not emit compiler output files like JavaScript source code, source-maps or declarations if any errors were reported.
	// Experimental.
	NoEmitOnError *bool `json:"noEmitOnError"`
	// Report errors for fallthrough cases in switch statements.
	//
	// Ensures that any non-empty
	// case inside a switch statement includes either break or return. This means you won’t
	// accidentally ship a case fallthrough bug.
	// Experimental.
	NoFallthroughCasesInSwitch *bool `json:"noFallthroughCasesInSwitch"`
	// In some cases where no type annotations are present, TypeScript will fall back to a type of any for a variable when it cannot infer the type.
	// Experimental.
	NoImplicitAny *bool `json:"noImplicitAny"`
	// When enabled, TypeScript will check all code paths in a function to ensure they return a value.
	// Experimental.
	NoImplicitReturns *bool `json:"noImplicitReturns"`
	// Raise error on ‘this’ expressions with an implied ‘any’ type.
	// Experimental.
	NoImplicitThis *bool `json:"noImplicitThis"`
	// Raise error on use of the dot syntax to access fields which are not defined.
	// Experimental.
	NoPropertyAccessFromIndexSignature *bool `json:"noPropertyAccessFromIndexSignature"`
	// Raise error when accessing indexes on objects with unknown keys defined in index signatures.
	// Experimental.
	NoUncheckedIndexedAccess *bool `json:"noUncheckedIndexedAccess"`
	// Report errors on unused local variables.
	// Experimental.
	NoUnusedLocals *bool `json:"noUnusedLocals"`
	// Report errors on unused parameters in functions.
	// Experimental.
	NoUnusedParameters *bool `json:"noUnusedParameters"`
	// Output directory for the compiled files.
	// Experimental.
	OutDir *string `json:"outDir"`
	// A series of entries which re-map imports to lookup locations relative to the baseUrl, there is a larger coverage of paths in the handbook.
	//
	// paths lets you declare how TypeScript should resolve an import in your require/imports.
	// Experimental.
	Paths *map[string]*[]*string `json:"paths"`
	// Allows importing modules with a ‘.json’ extension, which is a common practice in node projects. This includes generating a type for the import based on the static JSON shape.
	// Experimental.
	ResolveJsonModule *bool `json:"resolveJsonModule"`
	// Specifies the root directory of input files.
	//
	// Only use to control the output directory structure with `outDir`.
	// Experimental.
	RootDir *string `json:"rootDir"`
	// Skip type checking of all declaration files (*.d.ts).
	// Experimental.
	SkipLibCheck *bool `json:"skipLibCheck"`
	// The strict flag enables a wide range of type checking behavior that results in stronger guarantees of program correctness.
	//
	// Turning this on is equivalent to enabling all of the strict mode family
	// options, which are outlined below. You can then turn off individual strict mode family checks as
	// needed.
	// Experimental.
	Strict *bool `json:"strict"`
	// When strictNullChecks is false, null and undefined are effectively ignored by the language.
	//
	// This can lead to unexpected errors at runtime.
	// When strictNullChecks is true, null and undefined have their own distinct types and you’ll
	// get a type error if you try to use them where a concrete value is expected.
	// Experimental.
	StrictNullChecks *bool `json:"strictNullChecks"`
	// When set to true, TypeScript will raise an error when a class property was declared but not set in the constructor.
	// Experimental.
	StrictPropertyInitialization *bool `json:"strictPropertyInitialization"`
	// Do not emit declarations for code that has an @internal annotation in it’s JSDoc comment.
	// Experimental.
	StripInternal *bool `json:"stripInternal"`
	// Modern browsers support all ES6 features, so ES6 is a good choice.
	//
	// You might choose to set
	// a lower target if your code is deployed to older environments, or a higher target if your
	// code is guaranteed to run in newer environments.
	// Experimental.
	Target *string `json:"target"`
}

// Determines how JSX should get transformed into valid JavaScript.
// See: https://www.typescriptlang.org/docs/handbook/jsx.html
//
// Experimental.
type TypeScriptJsxMode string

const (
	TypeScriptJsxMode_PRESERVE TypeScriptJsxMode = "PRESERVE"
	TypeScriptJsxMode_REACT TypeScriptJsxMode = "REACT"
	TypeScriptJsxMode_REACT_NATIVE TypeScriptJsxMode = "REACT_NATIVE"
	TypeScriptJsxMode_REACT_JSX TypeScriptJsxMode = "REACT_JSX"
	TypeScriptJsxMode_REACT_JSXDEV TypeScriptJsxMode = "REACT_JSXDEV"
)

// Determines how modules get resolved.
// See: https://www.typescriptlang.org/docs/handbook/module-resolution.html
//
// Experimental.
type TypeScriptModuleResolution string

const (
	TypeScriptModuleResolution_CLASSIC TypeScriptModuleResolution = "CLASSIC"
	TypeScriptModuleResolution_NODE TypeScriptModuleResolution = "NODE"
)

// Experimental.
type TypescriptConfig interface {
	CompilerOptions() *TypeScriptCompilerOptions
	Exclude() *[]*string
	File() projen.JsonFile
	FileName() *string
	Include() *[]*string
	AddExclude(pattern *string)
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

// Experimental.
func (t *jsiiProxy_TypescriptConfig) AddExclude(pattern *string) {
	_jsii_.InvokeVoid(
		t,
		"addExclude",
		[]interface{}{pattern},
	)
}

// Experimental.
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
	CompilerOptions *TypeScriptCompilerOptions `json:"compilerOptions"`
	// Filters results from the "include" option.
	// Experimental.
	Exclude *[]*string `json:"exclude"`
	// Experimental.
	FileName *string `json:"fileName"`
	// Specifies a list of glob patterns that match TypeScript files to be included in compilation.
	// Experimental.
	Include *[]*string `json:"include"`
}

// Upgrade node project dependencies.
// Experimental.
type UpgradeDependencies interface {
	projen.Component
	ContainerOptions() *workflows.ContainerOptions
	SetContainerOptions(val *workflows.ContainerOptions)
	IgnoresProjen() *bool
	Project() projen.Project
	Workflows() *[]github.GithubWorkflow
	AddPostBuildSteps(steps ...*workflows.JobStep)
	PostSynthesize()
	PreSynthesize()
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

func (j *jsiiProxy_UpgradeDependencies) Project() projen.Project {
	var returns projen.Project
	_jsii_.Get(
		j,
		"project",
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

// Add steps to execute a successful build.
// Experimental.
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

// Called after synthesis.
//
// Order is *not* guaranteed.
// Experimental.
func (u *jsiiProxy_UpgradeDependencies) PostSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"postSynthesize",
		nil, // no parameters
	)
}

// Called before synthesis.
// Experimental.
func (u *jsiiProxy_UpgradeDependencies) PreSynthesize() {
	_jsii_.InvokeVoid(
		u,
		"preSynthesize",
		nil, // no parameters
	)
}

// Synthesizes files to the project output directory.
// Experimental.
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
	Exclude *[]*string `json:"exclude"`
	// Whether or not to ignore projen upgrades.
	// Experimental.
	IgnoreProjen *bool `json:"ignoreProjen"`
	// List of package names to include during the upgrade.
	// Experimental.
	Include *[]*string `json:"include"`
	// Title of the pull request to use (should be all lower-case).
	// Experimental.
	PullRequestTitle *string `json:"pullRequestTitle"`
	// Add Signed-off-by line by the committer at the end of the commit log message.
	// Experimental.
	Signoff *bool `json:"signoff"`
	// The name of the task that will be created.
	//
	// This will also be the workflow name.
	// Experimental.
	TaskName *string `json:"taskName"`
	// Include a github workflow for creating PR's that upgrades the required dependencies, either by manual dispatch, or by a schedule.
	//
	// If this is `false`, only a local projen task is created, which can be executed manually to
	// upgrade the dependencies.
	// Experimental.
	Workflow *bool `json:"workflow"`
	// Options for the github workflow.
	//
	// Only applies if `workflow` is true.
	// Experimental.
	WorkflowOptions *UpgradeDependenciesWorkflowOptions `json:"workflowOptions"`
}

// How often to check for new versions and raise pull requests for version upgrades.
// Experimental.
type UpgradeDependenciesSchedule interface {
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
	// List of branches to create PR's for.
	// Experimental.
	Branches *[]*string `json:"branches"`
	// Job container options.
	// Experimental.
	Container *workflows.ContainerOptions `json:"container"`
	// The git identity to use for commits.
	// Experimental.
	GitIdentity *github.GitIdentity `json:"gitIdentity"`
	// Labels to apply on the PR.
	// Experimental.
	Labels *[]*string `json:"labels"`
	// Execute 'build' after the upgrade.
	//
	// When true, the workflow will run the project build task after the dependency upgrade.
	// This means that the PR created will include any changes caused by the `build` command,
	// (e.g project synth changes, test snapshots)
	//
	// This is necessary when using the default github token.
	// See: `secret` for more details.
	//
	// Experimental.
	Rebuild *bool `json:"rebuild"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `json:"runsOn"`
	// Schedule to run on.
	// Experimental.
	Schedule UpgradeDependenciesSchedule `json:"schedule"`
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
	Secret *string `json:"secret"`
}

