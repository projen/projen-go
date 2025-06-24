package biomeconfig


// Options that changes how the CSS linter behaves.
// Experimental.
type CssLinterConfiguration struct {
	// Control the linter for CSS files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

