package biomeconfig


// Options that change how CSS assist behaves.
// Experimental.
type CssAssistConfiguration struct {
	// Controls assist actions for CSS files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

