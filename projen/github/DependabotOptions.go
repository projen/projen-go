package github


// Experimental.
type DependabotOptions struct {
	// https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#allow.
	//
	// Use the allow option to customize which dependencies are updated. This
	// applies to both version and security updates.
	// Experimental.
	Allow *[]*DependabotAllow `field:"optional" json:"allow" yaml:"allow"`
	// Specify individual assignees or teams of assignees for all pull requests raised for a package manager.
	// Experimental.
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#groups.
	//
	// You can create groups to package dependency updates together into a single PR.
	// Experimental.
	Groups *map[string]*DependabotGroup `field:"optional" json:"groups" yaml:"groups"`
	// You can use the `ignore` option to customize which dependencies are updated.
	//
	// The ignore option supports the following options.
	// Experimental.
	Ignore *[]*DependabotIgnore `field:"optional" json:"ignore" yaml:"ignore"`
	// Ignores updates to `projen`.
	//
	// This is required since projen updates may cause changes in committed files
	// and anti-tamper checks will fail.
	//
	// Projen upgrades are covered through the `ProjenUpgrade` class.
	// Experimental.
	IgnoreProjen *bool `field:"optional" json:"ignoreProjen" yaml:"ignoreProjen"`
	// List of labels to apply to the created PR's.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Sets the maximum of pull requests Dependabot opens for version updates.
	//
	// Dependabot will not open any new requests until some of those open requests
	// are merged or closed.
	// Experimental.
	OpenPullRequestsLimit *float64 `field:"optional" json:"openPullRequestsLimit" yaml:"openPullRequestsLimit"`
	// Map of package registries to use.
	// Experimental.
	Registries *map[string]*DependabotRegistry `field:"optional" json:"registries" yaml:"registries"`
	// Specify individual reviewers or teams of reviewers for all pull requests raised for a package manager.
	// Experimental.
	Reviewers *[]*string `field:"optional" json:"reviewers" yaml:"reviewers"`
	// How often to check for new versions and raise pull requests.
	// Experimental.
	ScheduleInterval DependabotScheduleInterval `field:"optional" json:"scheduleInterval" yaml:"scheduleInterval"`
	// The strategy to use when edits manifest and lock files.
	// Experimental.
	VersioningStrategy VersioningStrategy `field:"optional" json:"versioningStrategy" yaml:"versioningStrategy"`
}

