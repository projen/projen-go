package workflows


// A secret definition for a `workflow_call` trigger.
// See: https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#onworkflow_callsecrets
//
// Experimental.
type WorkflowCallSecret struct {
	// A description of the secret parameter.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the secret is required.
	//
	// If `true`, the caller must supply
	// this secret when calling the workflow.
	// Default: false.
	//
	// Experimental.
	Required *bool `field:"optional" json:"required" yaml:"required"`
}

