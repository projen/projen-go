package sonarqube


// Options for `sonar.cpd.*` properties.
// Experimental.
type SonarqubeCpdOptions struct {
	// Comma-separated file path patterns to exclude from code duplication detection.
	//
	// Maps to `sonar.cpd.exclusions`.
	// Default: - no duplication exclusions.
	//
	// Experimental.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
}

