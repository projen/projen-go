package biomeconfig


// Experimental.
type OverrideLinterConfiguration struct {
	// The lint-domain configuration for matched files.
	// Experimental.
	Domains *RuleDomains `field:"optional" json:"domains" yaml:"domains"`
	// Enables or disables the linter for matched files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The lint-rule configuration for matched files.
	// Experimental.
	Rules *Rules `field:"optional" json:"rules" yaml:"rules"`
}

