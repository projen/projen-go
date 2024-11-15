package projen


// Options to pass to `modifyBranchEnvironment`.
// Experimental.
type VersionBranchOptions struct {
	// The major versions released from this branch.
	// Experimental.
	MajorVersion *float64 `field:"optional" json:"majorVersion" yaml:"majorVersion"`
	// The minimum major version to release.
	// Experimental.
	MinMajorVersion *float64 `field:"optional" json:"minMajorVersion" yaml:"minMajorVersion"`
	// The minor versions released from this branch.
	// Experimental.
	MinorVersion *float64 `field:"optional" json:"minorVersion" yaml:"minorVersion"`
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
}

