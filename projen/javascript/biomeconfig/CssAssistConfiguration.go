package biomeconfig


// Options that changes how the CSS assist behaves.
// Experimental.
type CssAssistConfiguration struct {
	// Control the assist for CSS files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

