package projen


// Experimental.
type GroupRunnerOptions struct {
	// Experimental.
	Group *string `field:"required" json:"group" yaml:"group"`
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

