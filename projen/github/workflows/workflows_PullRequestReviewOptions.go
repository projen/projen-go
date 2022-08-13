package workflows


// Pull request review options.
// Experimental.
type PullRequestReviewOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

