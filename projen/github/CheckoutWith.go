package github


// Options for `checkout`.
// Experimental.
type CheckoutWith struct {
	// Number of commits to fetch.
	//
	// 0 indicates all history for all branches and tags.
	// Default: 1.
	//
	// Experimental.
	FetchDepth *float64 `field:"optional" json:"fetchDepth" yaml:"fetchDepth"`
	// Whether LFS is enabled for the GitHub repository.
	// Default: false.
	//
	// Experimental.
	Lfs *bool `field:"optional" json:"lfs" yaml:"lfs"`
	// Relative path under $GITHUB_WORKSPACE to place the repository.
	// Default: - $GITHUB_WORKSPACE.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Branch or tag name.
	// Default: - the default branch is implicitly used.
	//
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// The repository (owner/repo) to use.
	// Default: - the default repository is implicitly used.
	//
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// A GitHub token to use when checking out the repository.
	//
	// If the intent is to push changes back to the branch, then you must use a
	// PAT with `repo` (and possibly `workflows`) permissions.
	// Default: - the default GITHUB_TOKEN is implicitly used.
	//
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

