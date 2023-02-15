package github


// Options for `Stale`.
// Experimental.
type StaleOptions struct {
	// How to handle stale issues.
	// Experimental.
	Issues *StaleBehavior `field:"optional" json:"issues" yaml:"issues"`
	// How to handle stale pull requests.
	// Experimental.
	PullRequest *StaleBehavior `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// Github Runner selection labels.
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
}

