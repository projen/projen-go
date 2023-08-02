package github


// Experimental.
type MergifyQueue struct {
	// A list of Conditions string that must match against the pull request for the pull request to be added to the queue.
	// See: https://docs.mergify.com/conditions/#conditions
	//
	// Experimental.
	Conditions *[]interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// The name of the queue.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Merge method to use.
	//
	// Possible values are `merge`, `squash`, `rebase` or `fast-forward`.
	// `fast-forward` is not supported on queues with `speculative_checks` > 1, `batch_size` > 1, or with `allow_inplace_checks` set to false.
	// Default: "merge".
	//
	// Experimental.
	MergeMethod *string `field:"optional" json:"mergeMethod" yaml:"mergeMethod"`
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

