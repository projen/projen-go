package biomeconfig


// Linter options specific to the JavaScript linter.
// Experimental.
type JsLinterConfiguration struct {
	// Control the linter for JavaScript (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

