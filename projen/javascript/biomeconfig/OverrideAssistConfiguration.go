package biomeconfig


// Experimental.
type OverrideAssistConfiguration struct {
	// The assist-action configuration for matched files.
	// Experimental.
	Actions *Actions `field:"optional" json:"actions" yaml:"actions"`
	// Enables or disables assist actions for matched files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

