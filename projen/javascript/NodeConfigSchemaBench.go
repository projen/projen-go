package javascript


// Experimental.
type NodeConfigSchemaBench struct {
	// launch benchmark runner on startup.
	// Experimental.
	Bench *bool `field:"optional" json:"bench" yaml:"bench"`
	// configures the type of benchmark isolation used in the benchmark runner.
	// Experimental.
	BenchIsolation *string `field:"optional" json:"benchIsolation" yaml:"benchIsolation"`
	// run benchmarks whose name matches this regular expression.
	// Experimental.
	BenchNamePattern *string `field:"optional" json:"benchNamePattern" yaml:"benchNamePattern"`
	// report benchmark output using the given reporter.
	// Experimental.
	BenchReporter *[]*string `field:"optional" json:"benchReporter" yaml:"benchReporter"`
	// report the given benchmark reporter to the given destination.
	// Experimental.
	BenchReporterDestination *[]*string `field:"optional" json:"benchReporterDestination" yaml:"benchReporterDestination"`
	// specify the number of measured benchmark samples.
	// Experimental.
	BenchSamples *float64 `field:"optional" json:"benchSamples" yaml:"benchSamples"`
	// specify the number of unreported benchmark warmup samples.
	// Experimental.
	BenchWarmup *float64 `field:"optional" json:"benchWarmup" yaml:"benchWarmup"`
}

