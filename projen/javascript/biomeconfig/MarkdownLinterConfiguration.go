package biomeconfig


// Options that change how the Markdown linter behaves.
// Experimental.
type MarkdownLinterConfiguration struct {
	// Control the linter for Markdown files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

