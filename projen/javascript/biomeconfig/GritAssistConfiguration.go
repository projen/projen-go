package biomeconfig


// Options that change how GritQL assist behaves.
// Experimental.
type GritAssistConfiguration struct {
	// Enables or disables assist actions for GritQL.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

