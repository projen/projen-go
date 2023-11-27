package github


// Experimental.
type AutoMergeOptions struct {
	// Number of approved code reviews.
	// Default: 1.
	//
	// Experimental.
	ApprovedReviews *float64 `field:"optional" json:"approvedReviews" yaml:"approvedReviews"`
	// List of labels that will prevent auto-merging.
	// Default: ['do-not-merge'].
	//
	// Experimental.
	BlockingLabels *[]*string `field:"optional" json:"blockingLabels" yaml:"blockingLabels"`
	// Name of the mergify queue.
	// Default: 'default'.
	//
	// Experimental.
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// Name of the mergify rule.
	// Default: 'Automatic merge on approval and successful build'.
	//
	// Experimental.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
}

