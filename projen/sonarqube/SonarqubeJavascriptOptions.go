package sonarqube


// Options for `sonar.javascript.*` properties.
// Experimental.
type SonarqubeJavascriptOptions struct {
	// Options for `sonar.javascript.lcov.*`.
	// Default: - no LCOV configuration.
	//
	// Experimental.
	Lcov *SonarqubeLcovOptions `field:"optional" json:"lcov" yaml:"lcov"`
}

