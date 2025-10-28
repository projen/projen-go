package biomeconfig


// Options that changes how the HTML assist behaves.
// Experimental.
type HtmlAssistConfiguration struct {
	// Control the assist for HTML (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

