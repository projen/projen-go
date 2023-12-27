package github


// Experimental.
type CommonWorkflowStepOptions struct {
	// A unique identifier for the step.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A condition to determine whether to run this step.
	// Experimental.
	If *string `field:"optional" json:"if" yaml:"if"`
	// The name of the step.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

