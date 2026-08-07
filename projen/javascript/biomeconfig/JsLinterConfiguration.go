package biomeconfig


// Linter options specific to the JavaScript linter.
// Experimental.
type JsLinterConfiguration struct {
	// Controls the linter for JavaScript and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

