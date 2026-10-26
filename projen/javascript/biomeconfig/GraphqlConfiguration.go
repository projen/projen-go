package biomeconfig


// Options applied to GraphQL files.
//
// Language-specific settings take precedence over corresponding global settings. Global settings
// apply when their language-specific counterparts are omitted, unless stated otherwise.
// Experimental.
type GraphqlConfiguration struct {
	// GraphQL assist options.
	// Experimental.
	Assist *GraphqlAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// GraphQL formatter options.
	// Experimental.
	Formatter *GraphqlFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// GraphQL linter options.
	// Experimental.
	Linter *GraphqlLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
}

