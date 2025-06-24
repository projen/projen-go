package biomeconfig


// Linter options specific to the JSON linter.
// Experimental.
type JsonLinterConfiguration struct {
	// Control the linter for JSON (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

