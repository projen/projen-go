package circleci


// The matrix stanza allows you to run a parameterized job multiple times with different arguments.
// See: https://circleci.com/docs/2.0/configuration-reference/#matrix-requires-version-21
//
// Experimental.
type Matrix struct {
	// An alias for the matrix, usable from another job’s requires stanza.
	//
	// Defaults to the name of the job being executed.
	// Experimental.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A map of parameter names to every value the job should be called with.
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

