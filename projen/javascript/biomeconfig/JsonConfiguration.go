package biomeconfig


// Options applied to JSON, JSONC, and recognized JSON-based configuration files.
//
// Language-specific settings take precedence over corresponding global settings. Global settings
// apply when their language-specific counterparts are omitted, unless stated otherwise.
// Experimental.
type JsonConfiguration struct {
	// Assist options.
	// Experimental.
	Assist *JsonAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// Formatting options.
	// Experimental.
	Formatter *JsonFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Linting options.
	// Experimental.
	Linter *JsonLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
	// Parsing options.
	// Experimental.
	Parser *JsonParserConfiguration `field:"optional" json:"parser" yaml:"parser"`
}

