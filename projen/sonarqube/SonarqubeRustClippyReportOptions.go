package sonarqube


// Options for `sonar.rust.clippyReport.*` properties.
// Experimental.
type SonarqubeRustClippyReportOptions struct {
	// Paths to Clippy JSON report files.
	//
	// Maps to `sonar.rust.clippyReport.reportPaths`.
	// Default: - not set.
	//
	// Experimental.
	ReportPaths *[]*string `field:"optional" json:"reportPaths" yaml:"reportPaths"`
}

