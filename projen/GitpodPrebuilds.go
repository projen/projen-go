package projen


// Configure the Gitpod App for prebuilds.
//
// Currently only GitHub is supported.
// See: https://www.gitpod.io/docs/prebuilds/
//
// Experimental.
type GitpodPrebuilds struct {
	// Add a "Review in Gitpod" button to the pull request's description.
	// Default: false.
	//
	// Experimental.
	AddBadge *bool `field:"optional" json:"addBadge" yaml:"addBadge"`
	// Add a check to pull requests.
	// Default: true.
	//
	// Experimental.
	AddCheck *bool `field:"optional" json:"addCheck" yaml:"addCheck"`
	// Add a "Review in Gitpod" button as a comment to pull requests.
	// Default: false.
	//
	// Experimental.
	AddComment *bool `field:"optional" json:"addComment" yaml:"addComment"`
	// Add a label once the prebuild is ready to pull requests.
	// Default: false.
	//
	// Experimental.
	AddLabel *bool `field:"optional" json:"addLabel" yaml:"addLabel"`
	// Enable for all branches in this repo.
	// Default: false.
	//
	// Experimental.
	Branches *bool `field:"optional" json:"branches" yaml:"branches"`
	// Enable for the master/default branch.
	// Default: true.
	//
	// Experimental.
	Master *bool `field:"optional" json:"master" yaml:"master"`
	// Enable for pull requests coming from this repo.
	// Default: true.
	//
	// Experimental.
	PullRequests *bool `field:"optional" json:"pullRequests" yaml:"pullRequests"`
	// Enable for pull requests coming from forks.
	// Default: false.
	//
	// Experimental.
	PullRequestsFromForks *bool `field:"optional" json:"pullRequestsFromForks" yaml:"pullRequestsFromForks"`
}

