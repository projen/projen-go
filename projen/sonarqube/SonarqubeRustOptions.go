package sonarqube


// Options for `sonar.rust.*` properties.
// Experimental.
type SonarqubeRustOptions struct {
	// Options for `sonar.rust.clippy.*`.
	// Default: - no clippy configuration.
	//
	// Experimental.
	Clippy *SonarqubeRustClippyOptions `field:"optional" json:"clippy" yaml:"clippy"`
	// Options for `sonar.rust.clippyReport.*`.
	// Default: - no clippy report configuration.
	//
	// Experimental.
	ClippyReport *SonarqubeRustClippyReportOptions `field:"optional" json:"clippyReport" yaml:"clippyReport"`
	// Options for `sonar.rust.lcov.*`.
	// Default: - no Rust LCOV configuration.
	//
	// Experimental.
	Lcov *SonarqubeLcovOptions `field:"optional" json:"lcov" yaml:"lcov"`
}

