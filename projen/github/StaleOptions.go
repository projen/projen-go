package github

import (
	"github.com/projen/projen-go/projen"
)

// Options for `Stale`.
// Experimental.
type StaleOptions struct {
	// How to handle stale issues.
	// Default: - By default, stale issues with no activity will be marked as
	// stale after 60 days and closed within 7 days.
	//
	// Experimental.
	Issues *StaleBehavior `field:"optional" json:"issues" yaml:"issues"`
	// How to handle stale pull requests.
	// Default: - By default, pull requests with no activity will be marked as
	// stale after 14 days and closed within 2 days with relevant comments.
	//
	// Experimental.
	PullRequest *StaleBehavior `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// Github Runner Group selection options.
	// Experimental.
	RunsOnGroup *projen.GroupRunnerOptions `field:"optional" json:"runsOnGroup" yaml:"runsOnGroup"`
}

