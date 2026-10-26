package biomeconfig


// Options that change how CSS assist behaves.
// Experimental.
type CssAssistConfiguration struct {
	// Enables or disables assist actions for CSS.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

