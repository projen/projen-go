package github


// Options for `PullRequestSource.fromBranch`.
// Experimental.
type PullRequestSourceBranchOptions struct {
	// The GitHub usernames that may author the pull request.
	// Default: - any author.
	//
	// Experimental.
	Authors *[]*string `field:"optional" json:"authors" yaml:"authors"`
	// The exact branch the pull request comes from.
	// Default: - use `branchPrefix`.
	//
	// Experimental.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// A prefix the branch the pull request comes from must start with.
	//
	// Make sure to include a separator at the end like `/` or `-`.
	// Default: - use `branch`.
	//
	// Experimental.
	BranchPrefix *string `field:"optional" json:"branchPrefix" yaml:"branchPrefix"`
	// The branches the pull request may target.
	// Default: - any target branch.
	//
	// Experimental.
	TargetBranches *[]*string `field:"optional" json:"targetBranches" yaml:"targetBranches"`
}

