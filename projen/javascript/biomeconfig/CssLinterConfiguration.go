package biomeconfig


// Options that change how the CSS linter behaves.
// Experimental.
type CssLinterConfiguration struct {
	// Enables or disables the linter for CSS.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

