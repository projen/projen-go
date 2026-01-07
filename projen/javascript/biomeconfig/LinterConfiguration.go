package biomeconfig


// Experimental.
type LinterConfiguration struct {
	// An object where the keys are the names of the domains, and the values are `all`, `recommended`, or `none`.
	// Experimental.
	Domains *map[string]RuleDomainValue `field:"optional" json:"domains" yaml:"domains"`
	// if `false`, it disables the feature and the linter won't be executed.
	//
	// `true` by default.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of glob patterns.
	//
	// The analyzer will handle only those files/folders that will
	// match these patterns.
	// Experimental.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
	// List of rules.
	// Experimental.
	Rules *Rules `field:"optional" json:"rules" yaml:"rules"`
}

