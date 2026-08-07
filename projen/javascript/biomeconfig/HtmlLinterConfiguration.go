package biomeconfig


// Options that change how the HTML linter behaves.
// Experimental.
type HtmlLinterConfiguration struct {
	// Controls the linter for HTML and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

