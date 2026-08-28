package polaris


// Experimental.
type CodingStandardDeviation struct {
	// The name of the rule to deviate from.
	// Experimental.
	Deviation *string `field:"required" json:"deviation" yaml:"deviation"`
	// The reason that the rule is being deviated from.
	// Experimental.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
}

