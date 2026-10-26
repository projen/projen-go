package biomeconfig


// Options that change how GraphQL assist behaves.
// Experimental.
type GraphqlAssistConfiguration struct {
	// Enables or disables assist actions for GraphQL.
	//
	// Defaults to `false`.
	// Default: false`.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

