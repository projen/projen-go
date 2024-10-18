package github


// The merge method used to add the PR to the merge queue.
//
// Behavior can be further configured in repository settings.
// Experimental.
type MergeMethod string

const (
	// Experimental.
	MergeMethod_SQUASH MergeMethod = "SQUASH"
	// Experimental.
	MergeMethod_MERGE MergeMethod = "MERGE"
	// Experimental.
	MergeMethod_REBASE MergeMethod = "REBASE"
)

