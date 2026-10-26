package biomeconfig


// Options applied to GritQL files.
//
// Language-specific settings take precedence over corresponding global settings. Global settings
// apply when their language-specific counterparts are omitted, unless stated otherwise.
// Experimental.
type GritConfiguration struct {
	// GritQL assist options.
	// Experimental.
	Assist *GritAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// GritQL formatter options.
	// Experimental.
	Formatter *GritFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// GritQL linter options.
	// Experimental.
	Linter *GritLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
}

