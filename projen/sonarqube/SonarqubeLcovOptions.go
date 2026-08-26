package sonarqube


// Options for lcov report paths (shared between languages).
// Experimental.
type SonarqubeLcovOptions struct {
	// Comma-separated paths to LCOV coverage report files.
	//
	// Maps to `sonar.<language>.lcov.reportPaths`.
	// Default: - not set.
	//
	// Experimental.
	ReportPaths *[]*string `field:"optional" json:"reportPaths" yaml:"reportPaths"`
}

