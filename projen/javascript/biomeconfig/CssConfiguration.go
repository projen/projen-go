package biomeconfig


// Options applied to CSS and languages that extend it.
//
// Language-specific settings take precedence over corresponding global settings. Global settings
// apply when their language-specific counterparts are omitted, unless stated otherwise.
// Experimental.
type CssConfiguration struct {
	// CSS assist options.
	// Experimental.
	Assist *CssAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// CSS formatter options.
	// Experimental.
	Formatter *CssFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Reserved for future CSS analyzer support.
	//
	// This option currently has no effect.
	// Experimental.
	Globals *[]*string `field:"optional" json:"globals" yaml:"globals"`
	// CSS linter options.
	// Experimental.
	Linter *CssLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// CSS parsing options.
	// Experimental.
	Parser *CssParserConfiguration `field:"optional" json:"parser" yaml:"parser"`
}

