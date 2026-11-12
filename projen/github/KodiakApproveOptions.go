package github


// Options for Kodiak `[approve]`.
// Experimental.
type KodiakApproveOptions struct {
	// Automatically approve pull requests opened by these usernames.
	// Default: - no usernames auto-approved.
	//
	// Experimental.
	AutoApproveUsernames *[]*string `field:"optional" json:"autoApproveUsernames" yaml:"autoApproveUsernames"`
}

