package workflows


// Milestone options.
// Experimental.
type MilestoneOptions struct {
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

