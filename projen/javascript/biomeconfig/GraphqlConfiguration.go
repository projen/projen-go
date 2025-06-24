package biomeconfig


// Options applied to GraphQL files.
// Experimental.
type GraphqlConfiguration struct {
	// Assist options.
	// Experimental.
	Assist *GraphqlAssistConfiguration `field:"optional" json:"assist" yaml:"assist"`
	// GraphQL formatter options.
	// Experimental.
	Formatter *GraphqlFormatterConfiguration `field:"optional" json:"formatter" yaml:"formatter"`
	// Experimental.
	Linter *GraphqlLinterConfiguration `field:"optional" json:"linter" yaml:"linter"`
}

