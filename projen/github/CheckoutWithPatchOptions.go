package github


// Options for `checkoutWithPatch`.
// Experimental.
type CheckoutWithPatchOptions struct {
	// Whether LFS is enabled for the GitHub repository.
	// Default: false.
	//
	// Experimental.
	Lfs *bool `field:"optional" json:"lfs" yaml:"lfs"`
	// The name of the artifact the patch is stored as.
	// Default: ".repo.patch"
	//
	// Experimental.
	PatchFile *string `field:"optional" json:"patchFile" yaml:"patchFile"`
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

