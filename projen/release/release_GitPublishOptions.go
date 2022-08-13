package release


// Publishing options for Git releases.
// Experimental.
type GitPublishOptions struct {
	// The location of an .md file (relative to `dist/`) that includes the changelog for the release.
	//
	// Example:
	//   changelog.md
	//
	// Experimental.
	ChangelogFile *string `field:"required" json:"changelogFile" yaml:"changelogFile"`
	// The location of a text file (relative to `dist/`) that contains the release tag.
	//
	// Example:
	//   releasetag.txt
	//
	// Experimental.
	ReleaseTagFile *string `field:"required" json:"releaseTagFile" yaml:"releaseTagFile"`
	// The location of a text file (relative to `dist/`) that contains the version number.
	//
	// Example:
	//   version.txt
	//
	// Experimental.
	VersionFile *string `field:"required" json:"versionFile" yaml:"versionFile"`
	// Branch to push to.
	// Experimental.
	GitBranch *string `field:"optional" json:"gitBranch" yaml:"gitBranch"`
	// Override git-push command.
	//
	// Set to an empty string to disable pushing.
	// Experimental.
	GitPushCommand *string `field:"optional" json:"gitPushCommand" yaml:"gitPushCommand"`
	// The location of an .md file that includes the project-level changelog.
	// Experimental.
	ProjectChangelogFile *string `field:"optional" json:"projectChangelogFile" yaml:"projectChangelogFile"`
}

