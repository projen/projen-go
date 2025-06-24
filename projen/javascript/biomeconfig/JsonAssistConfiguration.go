package biomeconfig


// Linter options specific to the JSON linter.
// Experimental.
type JsonAssistConfiguration struct {
	// Control the assist for JSON (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

