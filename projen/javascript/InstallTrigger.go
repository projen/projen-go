package javascript


// Describes why dependencies need to be installed.
// Experimental.
type InstallTrigger struct {
	// The reason for the install.
	// Experimental.
	Reason InstallReason `field:"required" json:"reason" yaml:"reason"`
	// A unified diff of the package.json changes. Only present when reason is `PACKAGE_JSON_CHANGED`.
	// Experimental.
	Diff *[]*string `field:"optional" json:"diff" yaml:"diff"`
	// Human-readable descriptions of resolved dependency version changes.
	//
	// Only present when reason is `DEPS_RESOLVED`.
	// Experimental.
	Resolutions *[]*string `field:"optional" json:"resolutions" yaml:"resolutions"`
}

