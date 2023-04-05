package github


// Options for `uploadGitPatch`.
// Experimental.
type UploadGitPatchOptions struct {
	// The name of the output to emit.
	//
	// It will be set to `true` if there was a diff.
	// Experimental.
	OutputName *string `field:"required" json:"outputName" yaml:"outputName"`
	// The step ID which produces the output which indicates if a patch was created.
	// Experimental.
	StepId *string `field:"required" json:"stepId" yaml:"stepId"`
	// Fail if a mutation was found and print this error message.
	// Experimental.
	MutationError *string `field:"optional" json:"mutationError" yaml:"mutationError"`
	// The name of the artifact the patch is stored as.
	// Experimental.
	PatchFile *string `field:"optional" json:"patchFile" yaml:"patchFile"`
	// The name of the step.
	// Experimental.
	StepName *string `field:"optional" json:"stepName" yaml:"stepName"`
}

