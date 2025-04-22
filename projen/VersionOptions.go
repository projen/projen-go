package projen


// Options for `Version`.
// Experimental.
type VersionOptions struct {
	// The name of the directory into which `changelog.md` and `version.txt` files are emitted.
	// Experimental.
	ArtifactsDirectory *string `field:"required" json:"artifactsDirectory" yaml:"artifactsDirectory"`
	// A name of a .json file to set the `version` field in after a bump.
	//
	// Example:
	//   "package.json"
	//
	// Experimental.
	VersionInputFile *string `field:"required" json:"versionInputFile" yaml:"versionInputFile"`
	// The `commit-and-tag-version` compatible package used to bump the package version, as a dependency string.
	//
	// This can be any compatible package version, including the deprecated `standard-version@9`.
	// Default: "commit-and-tag-version@12".
	//
	// Experimental.
	BumpPackage *string `field:"optional" json:"bumpPackage" yaml:"bumpPackage"`
	// A shell command to control the next version to release.
	//
	// If present, this shell command will be run before the bump is executed, and
	// it determines what version to release. It will be executed in the following
	// environment:
	//
	// - Working directory: the project directory.
	// - `$VERSION`: the current version. Looks like `1.2.3`.
	// - `$LATEST_TAG`: the most recent tag. Looks like `prefix-v1.2.3`, or may be unset.
	// - `$SUGGESTED_BUMP`: the suggested bump action based on commits. One of `major|minor|patch|none`.
	//
	// The command should print one of the following to `stdout`:
	//
	// - Nothing: the next version number will be determined based on commit history.
	// - `x.y.z`: the next version number will be `x.y.z`.
	// - `major|minor|patch`: the next version number will be the current version number
	//   with the indicated component bumped.
	// Default: - The next version will be determined based on the commit history and project settings.
	//
	// Experimental.
	NextVersionCommand *string `field:"optional" json:"nextVersionCommand" yaml:"nextVersionCommand"`
	// Find commits that should be considered releasable Used to decide if a release is required.
	// Default: ReleasableCommits.everyCommit()
	//
	// Experimental.
	ReleasableCommits ReleasableCommits `field:"optional" json:"releasableCommits" yaml:"releasableCommits"`
	// The tag prefix corresponding to this version.
	// Experimental.
	TagPrefix *string `field:"optional" json:"tagPrefix" yaml:"tagPrefix"`
	// Custom configuration for versionrc file used by standard-release.
	// Experimental.
	VersionrcOptions *map[string]interface{} `field:"optional" json:"versionrcOptions" yaml:"versionrcOptions"`
}

