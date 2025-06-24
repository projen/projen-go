package biomeconfig


// Experimental.
type OverrideLinterConfiguration struct {
	// List of rules.
	// Experimental.
	Domains *map[string]RuleDomainValue `field:"optional" json:"domains" yaml:"domains"`
	// if `false`, it disables the feature and the linter won't be executed.
	//
	// `true` by default.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// List of rules.
	// Experimental.
	Rules *Rules `field:"optional" json:"rules" yaml:"rules"`
}

