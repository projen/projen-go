package sonarqube


// Options for `sonar.qualitygate.*` properties.
// Experimental.
type SonarqubeQualityGateOptions struct {
	// The number of seconds that the scanner should wait for a report to be processed.
	//
	// Maps to `sonar.qualitygate.timeout`.
	// Default: 300.
	//
	// Experimental.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// Forces the analysis step to poll the server and wait for the Quality Gate status.
	//
	// Will fail the pipeline if the quality gate fails.
	//
	// Maps to `sonar.qualitygate.wait`.
	// Default: false.
	//
	// Experimental.
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
}

