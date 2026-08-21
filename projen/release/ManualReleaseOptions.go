package release


// Experimental.
type ManualReleaseOptions struct {
	// Maintain a project-level changelog.
	// Default: true.
	//
	// Experimental.
	Changelog *bool `field:"optional" json:"changelog" yaml:"changelog"`
	// Project-level changelog file path.
	//
	// Ignored if `changelog` is false.
	// Default: 'CHANGELOG.md'
	//
	// Experimental.
	ChangelogPath *string `field:"optional" json:"changelogPath" yaml:"changelogPath"`
	// Override git-push command.
	//
	// Runs as a shell command in the `publish:git` task - where release
	// credentials are in scope - replacing the default
	// `git push --follow-tags origin <branch>`. Shell syntax is interpreted, so
	// keep it a literal command.
	//
	// Set to an empty string to disable pushing.
	// Experimental.
	GitPushCommand *string `field:"optional" json:"gitPushCommand" yaml:"gitPushCommand"`
}

