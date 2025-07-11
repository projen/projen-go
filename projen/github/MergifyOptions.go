package github


// Configure Mergify.
//
// This currently only offers a subset of options available.
// See: https://docs.mergify.com/configuration/file-format/
//
// Experimental.
type MergifyOptions struct {
	// The available merge queues.
	// Experimental.
	Queues *[]*MergifyQueue `field:"optional" json:"queues" yaml:"queues"`
	// Pull request automation rules.
	// Experimental.
	Rules *[]*MergifyRule `field:"optional" json:"rules" yaml:"rules"`
}

