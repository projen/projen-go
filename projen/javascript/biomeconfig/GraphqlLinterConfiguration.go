package biomeconfig


// Options that change how the GraphQL linter behaves.
// Experimental.
type GraphqlLinterConfiguration struct {
	// Controls the linter for GraphQL files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

