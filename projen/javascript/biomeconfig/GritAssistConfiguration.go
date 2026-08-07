package biomeconfig


// Options that change how GritQL assist behaves.
// Experimental.
type GritAssistConfiguration struct {
	// Controls assist actions for GritQL files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

