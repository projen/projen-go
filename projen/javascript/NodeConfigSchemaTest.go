package javascript


// Experimental.
type NodeConfigSchemaTest struct {
	// enable code coverage in the test runner.
	// Experimental.
	ExperimentalTestCoverage *bool `field:"optional" json:"experimentalTestCoverage" yaml:"experimentalTestCoverage"`
	// enable module mocking in the test runner.
	// Experimental.
	ExperimentalTestModuleMocks *bool `field:"optional" json:"experimentalTestModuleMocks" yaml:"experimentalTestModuleMocks"`
	// launch test runner on startup.
	// Experimental.
	Test *bool `field:"optional" json:"test" yaml:"test"`
	// specify test runner concurrency.
	// Experimental.
	TestConcurrency *float64 `field:"optional" json:"testConcurrency" yaml:"testConcurrency"`
	// the branch coverage minimum threshold.
	// Experimental.
	TestCoverageBranches *float64 `field:"optional" json:"testCoverageBranches" yaml:"testCoverageBranches"`
	// exclude files from coverage report that match this glob pattern.
	// Experimental.
	TestCoverageExclude *[]*string `field:"optional" json:"testCoverageExclude" yaml:"testCoverageExclude"`
	// the function coverage minimum threshold.
	// Experimental.
	TestCoverageFunctions *float64 `field:"optional" json:"testCoverageFunctions" yaml:"testCoverageFunctions"`
	// include files in coverage report that match this glob pattern.
	// Experimental.
	TestCoverageInclude *[]*string `field:"optional" json:"testCoverageInclude" yaml:"testCoverageInclude"`
	// the line coverage minimum threshold.
	// Experimental.
	TestCoverageLines *float64 `field:"optional" json:"testCoverageLines" yaml:"testCoverageLines"`
	// force test runner to exit upon completion.
	// Experimental.
	TestForceExit *bool `field:"optional" json:"testForceExit" yaml:"testForceExit"`
	// specifies the path to the global setup file.
	// Experimental.
	TestGlobalSetup *string `field:"optional" json:"testGlobalSetup" yaml:"testGlobalSetup"`
	// configures the type of test isolation used in the test runner.
	// Experimental.
	TestIsolation *string `field:"optional" json:"testIsolation" yaml:"testIsolation"`
	// run tests whose name matches this regular expression.
	// Experimental.
	TestNamePattern *[]*string `field:"optional" json:"testNamePattern" yaml:"testNamePattern"`
	// run tests with 'only' option set.
	// Experimental.
	TestOnly *bool `field:"optional" json:"testOnly" yaml:"testOnly"`
	// run tests in a random order.
	// Experimental.
	TestRandomize *bool `field:"optional" json:"testRandomize" yaml:"testRandomize"`
	// seed used to randomize test execution order.
	// Experimental.
	TestRandomSeed *float64 `field:"optional" json:"testRandomSeed" yaml:"testRandomSeed"`
	// report test output using the given reporter.
	// Experimental.
	TestReporter *[]*string `field:"optional" json:"testReporter" yaml:"testReporter"`
	// report given reporter to the given destination.
	// Experimental.
	TestReporterDestination *[]*string `field:"optional" json:"testReporterDestination" yaml:"testReporterDestination"`
	// specifies the path to the rerun state file.
	// Experimental.
	TestRerunFailures *string `field:"optional" json:"testRerunFailures" yaml:"testRerunFailures"`
	// run test at specific shard.
	// Experimental.
	TestShard *string `field:"optional" json:"testShard" yaml:"testShard"`
	// run tests whose name do not match this regular expression.
	// Experimental.
	TestSkipPattern *[]*string `field:"optional" json:"testSkipPattern" yaml:"testSkipPattern"`
	// specify test runner timeout.
	// Experimental.
	TestTimeout *float64 `field:"optional" json:"testTimeout" yaml:"testTimeout"`
	// regenerate test snapshots.
	// Experimental.
	TestUpdateSnapshots *bool `field:"optional" json:"testUpdateSnapshots" yaml:"testUpdateSnapshots"`
}

