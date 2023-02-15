package github


// Experimental.
type AutoMergeOptions struct {
	// Number of approved code reviews.
	// Experimental.
	ApprovedReviews *float64 `field:"optional" json:"approvedReviews" yaml:"approvedReviews"`
	// List of labels that will prevent auto-merging.
	// Experimental.
	BlockingLabels *[]*string `field:"optional" json:"blockingLabels" yaml:"blockingLabels"`
}

