package github


// Experimental.
type MergifyOptions struct {
	// Experimental.
	Queues *[]*MergifyQueue `field:"optional" json:"queues" yaml:"queues"`
	// Experimental.
	Rules *[]*MergifyRule `field:"optional" json:"rules" yaml:"rules"`
}

