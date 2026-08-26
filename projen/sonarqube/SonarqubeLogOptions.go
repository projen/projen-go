package sonarqube


// Options for `sonar.log.*` properties.
// Experimental.
type SonarqubeLogOptions struct {
	// Controls the quantity/level of logs produced during analysis.
	//
	// Maps to `sonar.log.level`.
	// Default: SonarqubeLogLevel.INFO
	//
	// Experimental.
	Level SonarqubeLogLevel `field:"optional" json:"level" yaml:"level"`
}

