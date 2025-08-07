package release


// Options for a release branch.
// Experimental.
type BranchOptions struct {
	// The major versions released from this branch.
	// Experimental.
	MajorVersion *float64 `field:"required" json:"majorVersion" yaml:"majorVersion"`
	// The GitHub Actions environment used for the release.
	//
	// This can be used to add an explicit approval step to the release
	// or limit who can initiate a release through environment protection rules.
	//
	// When multiple artifacts are released, the environment can be overwritten
	// on a per artifact basis.
	// Default: - no environment used, unless set at the artifact level.
	//
	// Experimental.
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The minimum major version to release.
	// Experimental.
	MinMajorVersion *float64 `field:"optional" json:"minMajorVersion" yaml:"minMajorVersion"`
	// The minor versions released from this branch.
	// Experimental.
	MinorVersion *float64 `field:"optional" json:"minorVersion" yaml:"minorVersion"`
	// The npm distribution tag to use for this branch.
	// Default: "latest".
	//
	// Experimental.
	NpmDistTag *string `field:"optional" json:"npmDistTag" yaml:"npmDistTag"`
	// Bump the version as a pre-release tag.
	// Default: - normal releases.
	//
	// Experimental.
	Prerelease *string `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Automatically add the given prefix to release tags. Useful if you are releasing on multiple branches with overlapping version numbers.
	//
	// Note: this prefix is used to detect the latest tagged version
	// when bumping, so if you change this on a project with an existing version
	// history, you may need to manually tag your latest release
	// with the new prefix.
	// Default: - no prefix.
	//
	// Experimental.
	TagPrefix *string `field:"optional" json:"tagPrefix" yaml:"tagPrefix"`
	// The name of the release workflow.
	// Default: "release-BRANCH".
	//
	// Experimental.
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
}

