package biomeconfig


// Options that change how the GritQL linter behaves.
// Experimental.
type GritLinterConfiguration struct {
	// Controls the linter for GritQL files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

