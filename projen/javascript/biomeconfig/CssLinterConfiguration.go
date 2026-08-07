package biomeconfig


// Options that change how the CSS linter behaves.
// Experimental.
type CssLinterConfiguration struct {
	// Controls the linter for CSS files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

