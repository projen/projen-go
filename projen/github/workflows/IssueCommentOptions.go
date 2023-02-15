package workflows


// Issue comment options.
// Experimental.
type IssueCommentOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

