package workflows


// An input definition for a `workflow_call` trigger.
// See: https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#onworkflow_callinputs
//
// Experimental.
type WorkflowCallInput struct {
	// The data type of the input.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The default value for the input when it is not provided by the caller.
	// Default: - no default.
	//
	// Experimental.
	Default interface{} `field:"optional" json:"default" yaml:"default"`
	// A description of the input parameter.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the input is required.
	//
	// If `true`, the caller must supply
	// this input when calling the workflow.
	// Default: false.
	//
	// Experimental.
	Required *bool `field:"optional" json:"required" yaml:"required"`
}

