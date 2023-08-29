package github


// Options for requiring a contributor statement on Pull Requests.
// Experimental.
type ContributorStatementOptions struct {
	// Pull requests with one of these labels are exempted from a contributor statement.
	// Default: - no labels are excluded.
	//
	// Experimental.
	ExemptLabels *[]*string `field:"optional" json:"exemptLabels" yaml:"exemptLabels"`
	// Pull requests from these GitHub users are exempted from a contributor statement.
	// Default: - no users are exempted.
	//
	// Experimental.
	ExemptUsers *[]*string `field:"optional" json:"exemptUsers" yaml:"exemptUsers"`
}

