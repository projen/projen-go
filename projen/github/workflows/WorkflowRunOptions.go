package workflows


// Workflow run options.
// Experimental.
type WorkflowRunOptions struct {
	// Which branches or branch-ignore to limit the trigger to.
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Which activity types to trigger on.
	// Experimental.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
	// Which workflow to trigger on.
	// Experimental.
	Workflows *[]*string `field:"optional" json:"workflows" yaml:"workflows"`
}

