package biomeconfig


// Assist options specific to the JavaScript assist.
// Experimental.
type JsAssistConfiguration struct {
	// Controls assist actions for JavaScript and languages that extend it.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

