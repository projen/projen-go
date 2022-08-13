package workflows


// Pull request review comment options.
// Experimental.
type PullRequestReviewCommentOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

