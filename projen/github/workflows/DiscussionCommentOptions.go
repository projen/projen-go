package workflows


// Discussion comment options.
// Experimental.
type DiscussionCommentOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

