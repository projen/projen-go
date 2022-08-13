package github


// Experimental.
type MergifyQueue struct {
	// A list of Conditions string that must match against the pull request for the pull request to be added to the queue.
	// See: https://docs.mergify.com/conditions/#conditions
	//
	// Experimental.
	Conditions *[]interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The name of the queue.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

