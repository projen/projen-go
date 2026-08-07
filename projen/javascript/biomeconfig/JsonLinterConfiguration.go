package biomeconfig


// Linter options specific to the JSON linter.
// Experimental.
type JsonLinterConfiguration struct {
	// Controls the linter for JSON and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

