package biomeconfig


// Options that changes how the HTML linter behaves.
// Experimental.
type HtmlLinterConfiguration struct {
	// Control the linter for HTML (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

