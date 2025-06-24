package biomeconfig


// Experimental.
type OverrideAssistConfiguration struct {
	// List of actions.
	// Experimental.
	Actions *Actions `field:"optional" json:"actions" yaml:"actions"`
	// if `false`, it disables the feature and the assist won't be executed.
	//
	// `true` by default.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

