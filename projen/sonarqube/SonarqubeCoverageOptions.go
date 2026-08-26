package sonarqube


// Options for `sonar.coverage.*` properties.
// Experimental.
type SonarqubeCoverageOptions struct {
	// Comma-separated file path patterns to exclude from test coverage calculations.
	//
	// Maps to `sonar.coverage.exclusions`.
	// Default: - no coverage exclusions.
	//
	// Experimental.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
}

