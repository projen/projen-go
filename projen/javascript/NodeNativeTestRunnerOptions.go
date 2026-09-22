package javascript


// Options for Node.js' built-in test runner (`node --test`).
// Experimental.
type NodeNativeTestRunnerOptions struct {
	// Indicates whether the coverage information should be collected while executing the test, via `--experimental-test-coverage`.
	// Default: true.
	//
	// Experimental.
	CollectCoverage *bool `field:"optional" json:"collectCoverage" yaml:"collectCoverage"`
	// Path to the JSON configuration file for the test runner.
	// Default: "node.config.json"
	//
	// Experimental.
	ConfigFilePath *string `field:"optional" json:"configFilePath" yaml:"configFilePath"`
	// The directory where coverage files are output, if coverage collection is enabled.
	// Default: "coverage".
	//
	// Experimental.
	CoverageDirectory *string `field:"optional" json:"coverageDirectory" yaml:"coverageDirectory"`
	// An array of glob patterns that are matched against all file paths before executing coverage collection.
	//
	// If a file path matches any of the
	// patterns, coverage information will be skipped for it.
	// Default: ["**\/test/**", "**\/__tests__/**"].
	//
	// Experimental.
	CoveragePathIgnorePatterns *[]*string `field:"optional" json:"coveragePathIgnorePatterns" yaml:"coveragePathIgnorePatterns"`
	// Additional options to pass to the `node --test` CLI invocation.
	//
	// Each element is passed as a single argument, exactly as given: no shell
	// parses these, so a flag and its value need separate elements
	// (`["--foo", "bar"]`, not `["--foo bar"]`).
	// Default: - no extra options.
	//
	// Experimental.
	ExtraCliOptions *[]*string `field:"optional" json:"extraCliOptions" yaml:"extraCliOptions"`
	// This option allows the use of a custom global setup module which exports a function that is triggered once before all test suites.
	//
	// Written as `test-global-setup` in the generated Node.js configuration
	// file.
	// Default: - undefined.
	//
	// Experimental.
	GlobalSetup *string `field:"optional" json:"globalSetup" yaml:"globalSetup"`
	// Enable module mocking support via `--experimental-test-module-mocks`.
	// Default: false.
	//
	// Experimental.
	ModuleMocks *bool `field:"optional" json:"moduleMocks" yaml:"moduleMocks"`
	// Additional entries for the `nodeOptions` section of the generated configuration file (e.g. `enableSourceMaps`, `disableWarning`).
	// Default: - no additional node options.
	//
	// Experimental.
	NodeOptions *NodeConfigSchemaNodeOptions `field:"optional" json:"nodeOptions" yaml:"nodeOptions"`
	// Preserve the default reporters (`spec`, `lcov`, `junit`) when additional reporters are added.
	// Default: true.
	//
	// Experimental.
	PreserveDefaultReporters *bool `field:"optional" json:"preserveDefaultReporters" yaml:"preserveDefaultReporters"`
	// Additional reporters to configure (e.g. `{ name: "tap", destination: Destination.file("test-reports/tap.txt") }`).
	//
	// These are added on top of the default reporters (`spec`, `lcov`, `junit`),
	// which are controlled via `collectCoverage`. Use
	// `NodeNativeTestRunner.addReporter`/`removeReporter`/`listReporters` to
	// manage reporters after construction.
	// Default: - no additional reporters.
	//
	// Experimental.
	Reporters *[]*NodeReporter `field:"optional" json:"reporters" yaml:"reporters"`
	// Additional entries for the `test` section of the generated configuration file (e.g. `testConcurrency`, `testTimeout`).
	// Default: - no additional options.
	//
	// Experimental.
	TestConfig *NodeConfigSchemaTest `field:"optional" json:"testConfig" yaml:"testConfig"`
	// Glob patterns matching the files that contain tests.
	//
	// By default it
	// combines Node.js' own default test file discovery with Jest conventions.
	// Default: - combines Node.js' own default test file discovery with Jest conventions
	//
	// Experimental.
	TestMatch *[]*string `field:"optional" json:"testMatch" yaml:"testMatch"`
	// Whether to enable transformation of TypeScript-only syntax (e.g. enums, namespaces).
	//
	// Uses `amaro` (the TypeScript transformer used internally by Node.js) as an
	// external loader via `--import=amaro/transform`. Adds a dependency on the
	// `amaro` package and enables `--enable-source-maps` to preserve accurate
	// stack traces.
	// See: https://github.com/nodejs/amaro
	//
	// Default: false.
	//
	// Experimental.
	TransformTypes *bool `field:"optional" json:"transformTypes" yaml:"transformTypes"`
	// Whether to update snapshots in task "test" (which is executed in task "build" and build workflows), or create a separate task "test:update" for updating snapshots.
	// Default: - ALWAYS.
	//
	// Experimental.
	UpdateSnapshot NodeTestUpdateSnapshot `field:"optional" json:"updateSnapshot" yaml:"updateSnapshot"`
}

