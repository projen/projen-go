package biomeconfig


// Options applied to CSS files.
// Experimental.
type CssConfiguration struct {
	// CSS assist options.
	// Experimental.
	Assist *CssAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// CSS formatter options.
	// Experimental.
	Formatter *CssFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// CSS globals.
	// Experimental.
	Globals *[]*string `field:"optional" json:"globals" yaml:"globals"`
	// CSS linter options.
	// Experimental.
	Linter *CssLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// CSS parsing options.
	// Experimental.
	Parser *CssParserConfiguration `field:"optional" json:"parser" yaml:"parser"`
}

