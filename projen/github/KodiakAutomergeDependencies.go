package github


// Options for Kodiak `[merge.automerge_dependencies]`.
// Experimental.
type KodiakAutomergeDependencies struct {
	// Restrict dependency auto-merging to specific bot or user logins.
	// Default: - no username filter.
	//
	// Experimental.
	Usernames *[]*string `field:"optional" json:"usernames" yaml:"usernames"`
	// Restrict auto-merging to specific semver upgrade levels.
	// Default: - no dependency upgrade levels are auto-merged.
	//
	// Experimental.
	Versions *[]KodiakDependencyVersion `field:"optional" json:"versions" yaml:"versions"`
}

