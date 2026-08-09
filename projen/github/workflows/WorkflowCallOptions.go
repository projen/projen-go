package workflows


// Options for the `workflow_call` trigger event, used to define a reusable workflow.
// See: https://docs.github.com/en/actions/using-workflows/reusing-workflows
//
// Experimental.
type WorkflowCallOptions struct {
	// A map of inputs that are passed from the caller workflow.
	//
	// Inputs are available in the called workflow using the `inputs` context.
	// Default: - no inputs.
	//
	// Experimental.
	Inputs *map[string]*WorkflowCallInput `field:"optional" json:"inputs" yaml:"inputs"`
	// A map of outputs that the called workflow can set.
	//
	// Called workflow outputs are available to all downstream jobs in the caller workflow.
	// Default: - no outputs.
	//
	// Experimental.
	Outputs *map[string]*WorkflowCallOutput `field:"optional" json:"outputs" yaml:"outputs"`
	// A map of secrets that can be used in the called workflow.
	//
	// Secrets are available in the called workflow using the `secrets` context.
	// Default: - no secrets.
	//
	// Experimental.
	Secrets *map[string]*WorkflowCallSecret `field:"optional" json:"secrets" yaml:"secrets"`
}

