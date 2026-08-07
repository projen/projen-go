package biomeconfig


// Options that change how GraphQL assist behaves.
// Experimental.
type GraphqlAssistConfiguration struct {
	// Controls assist actions for GraphQL files.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

