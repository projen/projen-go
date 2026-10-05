package github


// Options for `PullRequestSource.fromUsers`.
// Experimental.
type PullRequestSourceUsersOptions struct {
	// The GitHub usernames that may author the pull request.
	// Experimental.
	Logins *[]*string `field:"required" json:"logins" yaml:"logins"`
}

