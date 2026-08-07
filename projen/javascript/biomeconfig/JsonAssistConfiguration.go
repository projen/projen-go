package biomeconfig


// Assist options specific to the JSON linter.
// Experimental.
type JsonAssistConfiguration struct {
	// Controls assist actions for JSON and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

