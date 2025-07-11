package github


// Experimental.
type MergifyQueue struct {
	// Template to use as the commit message when using the merge or squash merge method.
	// Experimental.
	CommitMessageTemplate *string `field:"required" json:"commitMessageTemplate" yaml:"commitMessageTemplate"`
	// The name of the queue.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of conditions that needs to match to queue the pull request.
	// See: https://docs.mergify.com/configuration/file-format/#queue-rules
	//
	// Deprecated: use `queueConditions` instead.
	Conditions *[]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The list of conditions to match to get the queued pull request merged.
	//
	// This automatically includes the queueConditions.
	// In case of speculative merge pull request, the merge conditions are evaluated against the temporary pull request instead of the original one.
	// See: https://docs.mergify.com/conditions/#conditions
	//
	// Experimental.
	MergeConditions *[]interface{} `field:"optional" json:"mergeConditions" yaml:"mergeConditions"`
	// Merge method to use.
	//
	// Possible values are `merge`, `squash`, `rebase` or `fast-forward`.
	// `fast-forward` is not supported on queues with `speculative_checks` > 1, `batch_size` > 1, or with `allow_inplace_checks` set to false.
	// Default: "merge".
	//
	// Experimental.
	MergeMethod *string `field:"optional" json:"mergeMethod" yaml:"mergeMethod"`
	// The list of conditions that needs to match to queue the pull request.
	// See: https://docs.mergify.com/conditions/#conditions
	//
	// Experimental.
	QueueConditions *[]interface{} `field:"optional" json:"queueConditions" yaml:"queueConditions"`
	// Method to use to update the pull request with its base branch when the speculative check is done in-place.
	//
	// Possible values:
	//  - `merge` to merge the base branch into the pull request.
	//  - `rebase` to rebase the pull request against its base branch.
	//
	// Note that the `rebase` method has some drawbacks, see Mergify docs for details.
	// See: https://docs.mergify.com/actions/queue/#queue-rules
	//
	// Default: - `merge` for all merge methods except `fast-forward` where `rebase` is used.
	//
	// Experimental.
	UpdateMethod *string `field:"optional" json:"updateMethod" yaml:"updateMethod"`
}

