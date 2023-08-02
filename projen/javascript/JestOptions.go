package javascript


// Experimental.
type JestOptions struct {
	// Path to JSON config file for Jest.
	// Default: - No separate config file, jest settings are stored in package.json
	//
	// Experimental.
	ConfigFilePath *string `field:"optional" json:"configFilePath" yaml:"configFilePath"`
	// Collect coverage.
	//
	// Deprecated.
	// Default: true.
	//
	// Deprecated: use jestConfig.collectCoverage
	Coverage *bool `field:"optional" json:"coverage" yaml:"coverage"`
	// Include the `text` coverage reporter, which means that coverage summary is printed at the end of the jest execution.
	// Default: true.
	//
	// Experimental.
	CoverageText *bool `field:"optional" json:"coverageText" yaml:"coverageText"`
	// Additional options to pass to the Jest CLI invocation.
	// Default: - no extra options.
	//
	// Experimental.
	ExtraCliOptions *[]*string `field:"optional" json:"extraCliOptions" yaml:"extraCliOptions"`
	// Defines `testPathIgnorePatterns` and `coveragePathIgnorePatterns`.
	// Default: ["/node_modules/"].
	//
	// Deprecated: use jestConfig.coveragePathIgnorePatterns or jestConfig.testPathIgnorePatterns respectively
	IgnorePatterns *[]*string `field:"optional" json:"ignorePatterns" yaml:"ignorePatterns"`
	// Jest configuration.
	// Default: - default jest configuration.
	//
	// Experimental.
	JestConfig *JestConfigOptions `field:"optional" json:"jestConfig" yaml:"jestConfig"`
	// The version of jest to use.
	//
	// Note that same version is used as version of `@types/jest` and `ts-jest` (if Typescript in use), so given version should work also for those.
	// Default: - installs the latest jest version.
	//
	// Experimental.
	JestVersion *string `field:"optional" json:"jestVersion" yaml:"jestVersion"`
	// Result processing with jest-junit.
	//
	// Output directory is `test-reports/`.
	// Default: true.
	//
	// Experimental.
	JunitReporting *bool `field:"optional" json:"junitReporting" yaml:"junitReporting"`
	// Preserve the default Jest reporter when additional reporters are added.
	// Default: true.
	//
	// Experimental.
	PreserveDefaultReporters *bool `field:"optional" json:"preserveDefaultReporters" yaml:"preserveDefaultReporters"`
	// Whether to update snapshots in task "test" (which is executed in task "build" and build workflows), or create a separate task "test:update" for updating snapshots.
	// Default: - ALWAYS.
	//
	// Experimental.
	UpdateSnapshot UpdateSnapshot `field:"optional" json:"updateSnapshot" yaml:"updateSnapshot"`
}

