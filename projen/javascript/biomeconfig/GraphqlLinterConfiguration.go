package biomeconfig


// Options that change how the GraphQL linter behaves.
// Experimental.
type GraphqlLinterConfiguration struct {
	// Control the formatter for GraphQL files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

