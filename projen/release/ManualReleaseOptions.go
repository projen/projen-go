package release


// Experimental.
type ManualReleaseOptions struct {
	// Maintain a project-level changelog.
	// Experimental.
	Changelog *bool `field:"optional" json:"changelog" yaml:"changelog"`
	// Project-level changelog file path.
	//
	// Ignored if `changelog` is false.
	// Experimental.
	ChangelogPath *string `field:"optional" json:"changelogPath" yaml:"changelogPath"`
	// Override git-push command.
	//
	// Set to an empty string to disable pushing.
	// Experimental.
	GitPushCommand *string `field:"optional" json:"gitPushCommand" yaml:"gitPushCommand"`
}

