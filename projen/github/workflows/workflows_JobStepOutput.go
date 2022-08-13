package workflows


// An output binding for a job.
// Experimental.
type JobStepOutput struct {
	// The name of the job output that is being bound.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The ID of the step that exposes the output.
	// Experimental.
	StepId *string `field:"required" json:"stepId" yaml:"stepId"`
}

