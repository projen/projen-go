package workflows


// An output definition for a `workflow_call` trigger.
// See: https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#onworkflow_calloutputs
//
// Experimental.
type WorkflowCallOutput struct {
	// The value of the output.
	//
	// This must reference a job output from within
	// the reusable workflow, e.g. `${{ jobs.<job_id>.outputs.<output_name> }}`.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// A description of the output parameter.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

