package release


// Options for a release branch.
// Experimental.
type BranchOptions struct {
	// The major versions released from this branch.
	// Experimental.
	MajorVersion *float64 `field:"required" json:"majorVersion" yaml:"majorVersion"`
	// The minimum major version to release.
	// Experimental.
	MinMajorVersion *float64 `field:"optional" json:"minMajorVersion" yaml:"minMajorVersion"`
	// The npm distribution tag to use for this branch.
	// Experimental.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Bump the version as a pre-release tag.
	// Experimental.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Experimental.
	TagPrefix *string `field:"optional" json:"tagPrefix" yaml:"tagPrefix"`
	// The name of the release workflow.
	// Experimental.
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

