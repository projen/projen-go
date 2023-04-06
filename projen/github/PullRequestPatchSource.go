package github


// Experimental.
type PullRequestPatchSource struct {
	// Whether LFS is enabled for the GitHub repository.
	// Experimental.
	Lfs *bool `field:"optional" json:"lfs" yaml:"lfs"`
	// The name of the artifact the patch is stored as.
	// Experimental.
	PatchFile *string `field:"optional" json:"patchFile" yaml:"patchFile"`
	// Branch or tag name.
	// Experimental.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// The repository (owner/repo) to use.
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// A GitHub token to use when checking out the repository.
	//
	// If the intent is to push changes back to the branch, then you must use a
	// PAT with `repo` (and possibly `workflows`) permissions.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The id of the job that created the patch file.
	// Experimental.
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// The name of the output that indicates if a patch has been created.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
}

