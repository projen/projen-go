package sonarqube


// Options for `sonar.rust.clippy.*` properties.
// Experimental.
type SonarqubeRustClippyOptions struct {
	// Whether Clippy analysis is enabled.
	//
	// Maps to `sonar.rust.clippy.enabled`.
	// Default: true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

