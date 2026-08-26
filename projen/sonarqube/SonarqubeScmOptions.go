package sonarqube


// Options for `sonar.scm.*` properties.
// Experimental.
type SonarqubeScmOptions struct {
	// Options for `sonar.scm.exclusions.*`.
	// Default: - no exclusion overrides.
	//
	// Experimental.
	Exclusions *SonarqubeScmExclusionsOptions `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The SCM provider to use.
	//
	// Maps to `sonar.scm.provider`.
	// Default: - auto-detected.
	//
	// Experimental.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
}

