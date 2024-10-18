package github


// Options for 'MergeQueue'.
// Experimental.
type MergeQueueOptions struct {
	// Should pull requests be queued automatically to be merged once they pass required checks.
	// Default: true.
	//
	// Experimental.
	AutoQueue *bool `field:"optional" json:"autoQueue" yaml:"autoQueue"`
	// Configure auto-queue pull requests.
	// Default: - see AutoQueueOptions.
	//
	// Experimental.
	AutoQueueOptions *AutoQueueOptions `field:"optional" json:"autoQueueOptions" yaml:"autoQueueOptions"`
	// The branches that can be merged into using MergeQueue.
	// Default: - all branches.
	//
	// Experimental.
	TargetBranches *[]*string `field:"optional" json:"targetBranches" yaml:"targetBranches"`
}

