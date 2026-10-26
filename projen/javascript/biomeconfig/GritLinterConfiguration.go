package biomeconfig


// Options that change how the GritQL linter behaves.
// Experimental.
type GritLinterConfiguration struct {
	// Enables or disables the linter for GritQL.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

