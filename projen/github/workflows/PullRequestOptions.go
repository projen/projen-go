package workflows


// Pull request options.
// Experimental.
type PullRequestOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

