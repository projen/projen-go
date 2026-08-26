package sonarqube


// Options for `sonar.scm.exclusions.*` properties.
// Experimental.
type SonarqubeScmExclusionsOptions struct {
	// Whether to disable files ignored by the SCM (e.g., files in .gitignore) from being excluded from analysis.
	//
	// Maps to `sonar.scm.exclusions.disabled`.
	// Default: false.
	//
	// Experimental.
	Disabled *bool `field:"optional" json:"disabled" yaml:"disabled"`
}

