package workflows


// Merge group options.
// Experimental.
type MergeGroupOptions struct {
	// When using the merge_group events, you can configure a workflow to run on specific base branches.
	//
	// If not specified, all branches will
	// trigger the workflow.
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
}

