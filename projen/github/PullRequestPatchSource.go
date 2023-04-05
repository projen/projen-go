package github


// Experimental.
type PullRequestPatchSource struct {
	// The id of the job that created the patch file.
	// Experimental.
	JobId *string `field:"required" json:"jobId" yaml:"jobId"`
	// The name of the output that indicates if a patch has been created.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The name of the workflow that will create the PR.
	// Experimental.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// The name of the artifact the patch is stored as.
	// Experimental.
	PatchFile *string `field:"optional" json:"patchFile" yaml:"patchFile"`
}

