package workflows


// The Workflow dispatch event options.
// Experimental.
type WorkflowDispatchOptions struct {
	// Inputs for the workflow_dispatch event.
	// Experimental.
	Inputs *map[string]*WorkflowDispatchInput `field:"optional" json:"inputs" yaml:"inputs"`
}

