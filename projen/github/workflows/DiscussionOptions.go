package workflows


// Discussion options.
// Experimental.
type DiscussionOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

