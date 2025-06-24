package biomeconfig


// Options that changes how the GraphQL linter behaves.
// Experimental.
type GraphqlAssistConfiguration struct {
	// Control the formatter for GraphQL files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

