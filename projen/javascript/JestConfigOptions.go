package javascript


// Experimental.
type JestConfigOptions struct {
	// Escape hatch to allow any value.
	// Experimental.
	AdditionalOptions *map[string]interface{} `field:"optional" json:"additionalOptions" yaml:"additionalOptions"`
	// This option tells Jest that all imported modules in your tests should be mocked automatically.
	//
	// All modules used in your tests will have a replacement implementation, keeping the API surface.
	// Default: - false.
	//
	// Experimental.
	Automock *bool `field:"optional" json:"automock" yaml:"automock"`
	// By default, Jest runs all tests and produces all errors into the console upon completion.
	//
	// The bail config option can be used here to have Jest stop running tests after n failures.
	// Setting bail to true is the same as setting bail to 1.
	// Default: - 0.
	//
	// Experimental.
	Bail interface{} `field:"optional" json:"bail" yaml:"bail"`
	// The directory where Jest should store its cached dependency information.
	// Default: - "/tmp/<path>".
	//
	// Experimental.
	CacheDirectory *string `field:"optional" json:"cacheDirectory" yaml:"cacheDirectory"`
	// Automatically clear mock calls and instances before every test.
	//
	// Equivalent to calling jest.clearAllMocks() before each test.
	// This does not remove any mock implementation that may have been provided.
	// Default: true.
	//
	// Experimental.
	ClearMocks *bool `field:"optional" json:"clearMocks" yaml:"clearMocks"`
	// Indicates whether the coverage information should be collected while executing the test.
	//
	// Because this retrofits all executed files with coverage collection statements,
	// it may significantly slow down your tests.
	// Default: true.
	//
	// Experimental.
	CollectCoverage *bool `field:"optional" json:"collectCoverage" yaml:"collectCoverage"`
	// An array of glob patterns indicating a set of files for which coverage information should be collected.
	// Default: - undefined.
	//
	// Experimental.
	CollectCoverageFrom *[]*string `field:"optional" json:"collectCoverageFrom" yaml:"collectCoverageFrom"`
	// The directory where Jest should output its coverage files.
	// Default: "coverage".
	//
	// Experimental.
	CoverageDirectory *string `field:"optional" json:"coverageDirectory" yaml:"coverageDirectory"`
	// An array of regexp pattern strings that are matched against all file paths before executing the test.
	//
	// If the file path matches any of the patterns, coverage information will be skipped.
	// Default: "/node_modules/".
	//
	// Experimental.
	CoveragePathIgnorePatterns *[]*string `field:"optional" json:"coveragePathIgnorePatterns" yaml:"coveragePathIgnorePatterns"`
	// Indicates which provider should be used to instrument code for coverage.
	//
	// Allowed values are v8 (default) or babel.
	// Default: - "v8".
	//
	// Experimental.
	CoverageProvider *string `field:"optional" json:"coverageProvider" yaml:"coverageProvider"`
	// A list of reporter names that Jest uses when writing coverage reports.
	//
	// Any istanbul reporter can be used.
	// Default: - ["json", "lcov", "text", "clover", "cobertura"].
	//
	// Experimental.
	CoverageReporters *[]*string `field:"optional" json:"coverageReporters" yaml:"coverageReporters"`
	// Specify the global coverage thresholds.
	//
	// This will be used to configure minimum threshold enforcement
	// for coverage results. Thresholds can be specified as global, as a glob, and as a directory or file path.
	// If thresholds aren't met, jest will fail.
	// Default: - undefined.
	//
	// Experimental.
	CoverageThreshold *CoverageThreshold `field:"optional" json:"coverageThreshold" yaml:"coverageThreshold"`
	// This option allows the use of a custom dependency extractor.
	//
	// It must be a node module that exports an object with an extract function.
	// Default: - undefined.
	//
	// Experimental.
	DependencyExtractor *string `field:"optional" json:"dependencyExtractor" yaml:"dependencyExtractor"`
	// Allows for a label to be printed alongside a test while it is running.
	// Default: - undefined.
	//
	// Experimental.
	DisplayName interface{} `field:"optional" json:"displayName" yaml:"displayName"`
	// Make calling deprecated APIs throw helpful error messages.
	//
	// Useful for easing the upgrade process.
	// Default: - false.
	//
	// Experimental.
	ErrorOnDeprecated *bool `field:"optional" json:"errorOnDeprecated" yaml:"errorOnDeprecated"`
	// Test files run inside a vm, which slows calls to global context properties (e.g. Math). With this option you can specify extra properties to be defined inside the vm for faster lookups.
	// Default: - undefined.
	//
	// Experimental.
	ExtraGlobals *[]*string `field:"optional" json:"extraGlobals" yaml:"extraGlobals"`
	// Test files are normally ignored from collecting code coverage.
	//
	// With this option, you can overwrite this behavior and include otherwise ignored files in code coverage.
	// Default: - [''].
	//
	// Experimental.
	ForceCoverageMatch *[]*string `field:"optional" json:"forceCoverageMatch" yaml:"forceCoverageMatch"`
	// A set of global variables that need to be available in all test environments.
	// Default: - {}.
	//
	// Experimental.
	Globals interface{} `field:"optional" json:"globals" yaml:"globals"`
	// This option allows the use of a custom global setup module which exports an async function that is triggered once before all test suites.
	//
	// This function gets Jest's globalConfig object as a parameter.
	// Default: - undefined.
	//
	// Experimental.
	GlobalSetup *string `field:"optional" json:"globalSetup" yaml:"globalSetup"`
	// This option allows the use of a custom global teardown module which exports an async function that is triggered once after all test suites.
	//
	// This function gets Jest's globalConfig object as a parameter.
	// Default: - undefined.
	//
	// Experimental.
	GlobalTeardown *string `field:"optional" json:"globalTeardown" yaml:"globalTeardown"`
	// This will be used to configure the behavior of jest-haste-map, Jest's internal file crawler/cache system.
	// Default: - {}.
	//
	// Experimental.
	Haste *HasteConfig `field:"optional" json:"haste" yaml:"haste"`
	// Insert Jest's globals (expect, test, describe, beforeEach etc.) into the global environment. If you set this to false, you should import from.
	// Default: - true.
	//
	// Experimental.
	InjectGlobals *bool `field:"optional" json:"injectGlobals" yaml:"injectGlobals"`
	// A number limiting the number of tests that are allowed to run at the same time when using test.concurrent. Any test above this limit will be queued and executed once a slot is released.
	// Default: - 5.
	//
	// Experimental.
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Specifies the maximum number of workers the worker-pool will spawn for running tests.
	//
	// In single run mode,
	// this defaults to the number of the cores available on your machine minus one for the main thread
	// In watch mode, this defaults to half of the available cores on your machine.
	// For environments with variable CPUs available, you can use percentage based configuration: "maxWorkers": "50%".
	// Default: - the number of the cores available on your machine minus one for the main thread.
	//
	// Experimental.
	MaxWorkers interface{} `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// An array of directory names to be searched recursively up from the requiring module's location.
	//
	// Setting this option will override the default, if you wish to still search node_modules for packages
	// include it along with any other options: ["node_modules", "bower_components"].
	// Default: - ["node_modules"].
	//
	// Experimental.
	ModuleDirectories *[]*string `field:"optional" json:"moduleDirectories" yaml:"moduleDirectories"`
	// An array of file extensions your modules use.
	//
	// If you require modules without specifying a file extension,
	// these are the extensions Jest will look for, in left-to-right order.
	// Default: - ["js", "json", "jsx", "ts", "tsx", "node"].
	//
	// Experimental.
	ModuleFileExtensions *[]*string `field:"optional" json:"moduleFileExtensions" yaml:"moduleFileExtensions"`
	// A map from regular expressions to module names or to arrays of module names that allow to stub out resources, like images or styles with a single module.
	// Default: - null.
	//
	// Experimental.
	ModuleNameMapper *map[string]interface{} `field:"optional" json:"moduleNameMapper" yaml:"moduleNameMapper"`
	// An array of regexp pattern strings that are matched against all module paths before those paths are to be considered 'visible' to the module loader.
	//
	// If a given module's path matches any of the patterns,
	// it will not be require()-able in the test environment.
	// Default: - [].
	//
	// Experimental.
	ModulePathIgnorePatterns *[]*string `field:"optional" json:"modulePathIgnorePatterns" yaml:"modulePathIgnorePatterns"`
	// An alternative API to setting the NODE_PATH env variable, modulePaths is an array of absolute paths to additional locations to search when resolving modules.
	//
	// Use the <rootDir> string token to include
	// the path to your project's root directory. Example: ["<rootDir>/app/"].
	// Default: - [].
	//
	// Experimental.
	ModulePaths *[]*string `field:"optional" json:"modulePaths" yaml:"modulePaths"`
	// Activates notifications for test results.
	// Default: - false.
	//
	// Experimental.
	Notify *bool `field:"optional" json:"notify" yaml:"notify"`
	// Specifies notification mode.
	//
	// Requires notify: true.
	// Default: - failure-change.
	//
	// Experimental.
	NotifyMode *string `field:"optional" json:"notifyMode" yaml:"notifyMode"`
	// A preset that is used as a base for Jest's configuration.
	//
	// A preset should point to an npm module
	// that has a jest-preset.json or jest-preset.js file at the root.
	// Default: - undefined.
	//
	// Experimental.
	Preset *string `field:"optional" json:"preset" yaml:"preset"`
	// Sets the path to the prettier node module used to update inline snapshots.
	// Default: - "prettier".
	//
	// Experimental.
	PrettierPath *string `field:"optional" json:"prettierPath" yaml:"prettierPath"`
	// When the projects configuration is provided with an array of paths or glob patterns, Jest will run tests in all of the specified projects at the same time.
	//
	// This is great for monorepos or
	// when working on multiple projects at the same time.
	// Default: - undefined.
	//
	// Experimental.
	Projects *[]interface{} `field:"optional" json:"projects" yaml:"projects"`
	// Use this configuration option to add custom reporters to Jest.
	//
	// A custom reporter is a class
	// that implements onRunStart, onTestStart, onTestResult, onRunComplete methods that will be
	// called when any of those events occurs.
	// Default: - undefined.
	//
	// Experimental.
	Reporters *[]JestReporter `field:"optional" json:"reporters" yaml:"reporters"`
	// Automatically reset mock state before every test.
	//
	// Equivalent to calling jest.resetAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed but
	// does not restore their initial implementation.
	// Default: - false.
	//
	// Experimental.
	ResetMocks *bool `field:"optional" json:"resetMocks" yaml:"resetMocks"`
	// By default, each test file gets its own independent module registry.
	//
	// Enabling resetModules
	// goes a step further and resets the module registry before running each individual test.
	// Default: - false.
	//
	// Experimental.
	ResetModules *bool `field:"optional" json:"resetModules" yaml:"resetModules"`
	// This option allows the use of a custom resolver.
	//
	// https://jestjs.io/docs/en/configuration#resolver-string
	// Default: - undefined.
	//
	// Experimental.
	Resolver *string `field:"optional" json:"resolver" yaml:"resolver"`
	// Automatically restore mock state before every test.
	//
	// Equivalent to calling jest.restoreAllMocks()
	// before each test. This will lead to any mocks having their fake implementations removed and
	// restores their initial implementation.
	// Default: - false.
	//
	// Experimental.
	RestoreMocks *bool `field:"optional" json:"restoreMocks" yaml:"restoreMocks"`
	// The root directory that Jest should scan for tests and modules within.
	//
	// If you put your Jest
	// config inside your package.json and want the root directory to be the root of your repo, the
	// value for this config param will default to the directory of the package.json.
	// Default: - directory of the package.json
	//
	// Experimental.
	RootDir *string `field:"optional" json:"rootDir" yaml:"rootDir"`
	// A list of paths to directories that Jest should use to search for files in.
	// Default: - ["<rootDir>"].
	//
	// Experimental.
	Roots *[]*string `field:"optional" json:"roots" yaml:"roots"`
	// This option allows you to use a custom runner instead of Jest's default test runner.
	// Default: - "jest-runner".
	//
	// Experimental.
	Runner *string `field:"optional" json:"runner" yaml:"runner"`
	// A list of paths to modules that run some code to configure or set up the testing environment.
	//
	// Each setupFile will be run once per test file. Since every test runs in its own environment,
	// these scripts will be executed in the testing environment immediately before executing the
	// test code itself.
	// Default: - [].
	//
	// Experimental.
	SetupFiles *[]*string `field:"optional" json:"setupFiles" yaml:"setupFiles"`
	// A list of paths to modules that run some code to configure or set up the testing framework before each test file in the suite is executed.
	//
	// Since setupFiles executes before the test
	// framework is installed in the environment, this script file presents you the opportunity of
	// running some code immediately after the test framework has been installed in the environment.
	// Default: - [].
	//
	// Experimental.
	SetupFilesAfterEnv *[]*string `field:"optional" json:"setupFilesAfterEnv" yaml:"setupFilesAfterEnv"`
	// The number of seconds after which a test is considered as slow and reported as such in the results.
	// Default: - 5.
	//
	// Experimental.
	SlowTestThreshold *float64 `field:"optional" json:"slowTestThreshold" yaml:"slowTestThreshold"`
	// The path to a module that can resolve test<->snapshot path.
	//
	// This config option lets you customize
	// where Jest stores snapshot files on disk.
	// Default: - undefined.
	//
	// Experimental.
	SnapshotResolver *string `field:"optional" json:"snapshotResolver" yaml:"snapshotResolver"`
	// A list of paths to snapshot serializer modules Jest should use for snapshot testing.
	// Default: = [].
	//
	// Experimental.
	SnapshotSerializers *[]*string `field:"optional" json:"snapshotSerializers" yaml:"snapshotSerializers"`
	// The test environment that will be used for testing.
	//
	// The default environment in Jest is a
	// browser-like environment through jsdom. If you are building a node service, you can use the node
	// option to use a node-like environment instead.
	// Default: - "jsdom".
	//
	// Experimental.
	TestEnvironment *string `field:"optional" json:"testEnvironment" yaml:"testEnvironment"`
	// Test environment options that will be passed to the testEnvironment.
	//
	// The relevant options depend on the environment.
	// Default: - {}.
	//
	// Experimental.
	TestEnvironmentOptions interface{} `field:"optional" json:"testEnvironmentOptions" yaml:"testEnvironmentOptions"`
	// The exit code Jest returns on test failure.
	// Default: - 1.
	//
	// Experimental.
	TestFailureExitCode *float64 `field:"optional" json:"testFailureExitCode" yaml:"testFailureExitCode"`
	// The glob patterns Jest uses to detect test files.
	//
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Default: ['**\/__tests__/**\/*.[jt]s?(x)', '**\/*(*.)@(spec|test).[tj]s?(x)']
	//
	// Experimental.
	TestMatch *[]*string `field:"optional" json:"testMatch" yaml:"testMatch"`
	// An array of regexp pattern strings that are matched against all test paths before executing the test.
	//
	// If the test path matches any of the patterns, it will be skipped.
	// Default: - ["/node_modules/"].
	//
	// Experimental.
	TestPathIgnorePatterns *[]*string `field:"optional" json:"testPathIgnorePatterns" yaml:"testPathIgnorePatterns"`
	// The pattern or patterns Jest uses to detect test files.
	//
	// By default it looks for .js, .jsx, .ts and .tsx
	// files inside of __tests__ folders, as well as any files with a suffix of .test or .spec
	// (e.g. Component.test.js or Component.spec.js). It will also find files called test.js or spec.js.
	// Default: - (/__tests__/.*|(\\.|/)(test|spec))\\.[jt]sx?$
	//
	// Experimental.
	TestRegex interface{} `field:"optional" json:"testRegex" yaml:"testRegex"`
	// This option allows the use of a custom results processor.
	// Default: - undefined.
	//
	// Experimental.
	TestResultsProcessor *string `field:"optional" json:"testResultsProcessor" yaml:"testResultsProcessor"`
	// This option allows the use of a custom test runner.
	//
	// The default is jasmine2. A custom test runner
	// can be provided by specifying a path to a test runner implementation.
	// Default: - "jasmine2".
	//
	// Experimental.
	TestRunner *string `field:"optional" json:"testRunner" yaml:"testRunner"`
	// This option allows you to use a custom sequencer instead of Jest's default.
	//
	// Sort may optionally return a Promise.
	// Default: - "@jest/test-sequencer".
	//
	// Experimental.
	TestSequencer *string `field:"optional" json:"testSequencer" yaml:"testSequencer"`
	// Default timeout of a test in milliseconds.
	// Default: - 5000.
	//
	// Experimental.
	TestTimeout *float64 `field:"optional" json:"testTimeout" yaml:"testTimeout"`
	// This option sets the URL for the jsdom environment.
	//
	// It is reflected in properties such as location.href.
	// Default: - "http://localhost".
	//
	// Experimental.
	TestURL *string `field:"optional" json:"testURL" yaml:"testURL"`
	// Setting this value to legacy or fake allows the use of fake timers for functions such as setTimeout.
	//
	// Fake timers are useful when a piece of code sets a long timeout that we don't want to wait for in a test.
	// Default: - "real".
	//
	// Experimental.
	Timers *string `field:"optional" json:"timers" yaml:"timers"`
	// A map from regular expressions to paths to transformers.
	//
	// A transformer is a module that provides a
	// synchronous function for transforming source files.
	// Default: - {"\\.[jt]sx?$": "babel-jest"}
	//
	// Experimental.
	Transform *map[string]Transform `field:"optional" json:"transform" yaml:"transform"`
	// An array of regexp pattern strings that are matched against all source file paths before transformation.
	//
	// If the test path matches any of the patterns, it will not be transformed.
	// Default: - ["/node_modules/", "\\.pnp\\.[^\\\/]+$"]
	//
	// Experimental.
	TransformIgnorePatterns *[]*string `field:"optional" json:"transformIgnorePatterns" yaml:"transformIgnorePatterns"`
	// An array of regexp pattern strings that are matched against all modules before the module loader will automatically return a mock for them.
	//
	// If a module's path matches any of the patterns in this list, it
	// will not be automatically mocked by the module loader.
	// Default: - [].
	//
	// Experimental.
	UnmockedModulePathPatterns *[]*string `field:"optional" json:"unmockedModulePathPatterns" yaml:"unmockedModulePathPatterns"`
	// Indicates whether each individual test should be reported during the run.
	//
	// All errors will also
	// still be shown on the bottom after execution. Note that if there is only one test file being run
	// it will default to true.
	// Default: - false.
	//
	// Experimental.
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
	// Whether to use watchman for file crawling.
	// Default: - true.
	//
	// Experimental.
	Watchman *bool `field:"optional" json:"watchman" yaml:"watchman"`
	// An array of RegExp patterns that are matched against all source file paths before re-running tests in watch mode.
	//
	// If the file path matches any of the patterns, when it is updated, it will not trigger
	// a re-run of tests.
	// Default: - [].
	//
	// Experimental.
	WatchPathIgnorePatterns *[]*string `field:"optional" json:"watchPathIgnorePatterns" yaml:"watchPathIgnorePatterns"`
	// Default: -.
	//
	// Experimental.
	WatchPlugins *[]WatchPlugin `field:"optional" json:"watchPlugins" yaml:"watchPlugins"`
}

