package biomeconfig


// Options that change how HTML assist behaves.
// Experimental.
type HtmlAssistConfiguration struct {
	// Controls assist actions for HTML and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

