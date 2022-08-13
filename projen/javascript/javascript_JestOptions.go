package javascript


// Experimental.
type JestOptions struct {
	// Path to JSON config file for Jest.
	// Experimental.
	ConfigFilePath *string `field:"optional" json:"configFilePath" yaml:"configFilePath"`
	// Collect coverage.
	//
	// Deprecated.
	// Deprecated: use jestConfig.collectCoverage
	Coverage *bool `field:"optional" json:"coverage" yaml:"coverage"`
	// Include the `text` coverage reporter, which means that coverage summary is printed at the end of the jest execution.
	// Experimental.
	CoverageText *bool `field:"optional" json:"coverageText" yaml:"coverageText"`
	// Additional options to pass to the Jest CLI invocation.
	// Experimental.
	ExtraCliOptions *[]*string `field:"optional" json:"extraCliOptions" yaml:"extraCliOptions"`
	// Defines `testPathIgnorePatterns` and `coveragePathIgnorePatterns`.
	// Deprecated: use jestConfig.coveragePathIgnorePatterns or jestConfig.testPathIgnorePatterns respectively
	IgnorePatterns *[]*string `field:"optional" json:"ignorePatterns" yaml:"ignorePatterns"`
	// Jest configuration.
	// Experimental.
	JestConfig *JestConfigOptions `field:"optional" json:"jestConfig" yaml:"jestConfig"`
	// The version of jest to use.
	//
	// Note that same version is used as version of `@types/jest` and `ts-jest` (if Typescript in use), so given version should work also for those.
	// Experimental.
	JestVersion *string `field:"optional" json:"jestVersion" yaml:"jestVersion"`
	// Result processing with jest-junit.
	//
	// Output directory is `test-reports/`.
	// Experimental.
	JunitReporting *bool `field:"optional" json:"junitReporting" yaml:"junitReporting"`
	// Preserve the default Jest reporter when additional reporters are added.
	// Experimental.
	PreserveDefaultReporters *bool `field:"optional" json:"preserveDefaultReporters" yaml:"preserveDefaultReporters"`
}

