package github


// Options for Kodiak `[update]`.
// Experimental.
type KodiakUpdateOptions struct {
	// Automatically update out-of-date pull requests rather than waiting until queued for merge.
	// Default: false.
	//
	// Experimental.
	Always *bool `field:"optional" json:"always" yaml:"always"`
	// Only update pull requests carrying this label.
	// Default: - no label-based auto-update (see `update.always`)
	//
	// Experimental.
	AutoupdateLabel *string `field:"optional" json:"autoupdateLabel" yaml:"autoupdateLabel"`
	// Usernames whose pull requests should not be automatically updated.
	// Default: - no usernames ignored.
	//
	// Experimental.
	IgnoredUsernames *[]*string `field:"optional" json:"ignoredUsernames" yaml:"ignoredUsernames"`
}

