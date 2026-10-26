package biomeconfig


// Options that change how the GraphQL linter behaves.
// Experimental.
type GraphqlLinterConfiguration struct {
	// Enables or disables the linter for GraphQL.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

