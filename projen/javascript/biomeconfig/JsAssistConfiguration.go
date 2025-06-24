package biomeconfig


// Assist options specific to the JavaScript assist.
// Experimental.
type JsAssistConfiguration struct {
	// Control the assist for JavaScript (and its super languages) files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

