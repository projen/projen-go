package polaris


// Specifies how web application security analysis should be done.
// Experimental.
type CheckerConfigurationWebappSecurity struct {
	// Sets the web application checkers aggressiveness level.
	// Experimental.
	AggressivenessLevel CheckerConfigurationWebappSecurityAggressivenessLevel `field:"optional" json:"aggressivenessLevel" yaml:"aggressivenessLevel"`
	// Enables the checkers that are used for web application security analysis.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

