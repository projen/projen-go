package github


// Experimental.
type DependabotOptions struct {
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
	// Map of package registries to use.
	// Experimental.
	Registries *map[string]*DependabotRegistry `field:"optional" json:"registries" yaml:"registries"`
	// How often to check for new versions and raise pull requests.
	// Experimental.
	ScheduleInterval DependabotScheduleInterval `field:"optional" json:"scheduleInterval" yaml:"scheduleInterval"`
	// The strategy to use when edits manifest and lock files.
	// Experimental.
	VersioningStrategy VersioningStrategy `field:"optional" json:"versioningStrategy" yaml:"versioningStrategy"`
}

