package github


// Options for 'AutoQueue'.
// Experimental.
type AutoQueueOptions struct {
	// Only pull requests authored by these Github usernames will have auto-queue enabled.
	// Default: - pull requests from all users are eligible for auto-queuing.
	//
	// Experimental.
	AllowedUsernames *[]*string `field:"optional" json:"allowedUsernames" yaml:"allowedUsernames"`
	// Only pull requests with one of this labels will have auto-queue enabled.
	// Default: - all pull requests are eligible for auto-queueing.
	//
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// The method used to add the PR to the merge queue Any branch protection rules must allow this merge method.
	// Default: MergeMethod.SQUASH
	//
	// Experimental.
	MergeMethod MergeMethod `field:"optional" json:"mergeMethod" yaml:"mergeMethod"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
	// A GitHub secret name which contains a GitHub Access Token with write permissions for the `pull_request` scope.
	//
	// This token is used to enable auto-queue on pull requests.
	// Default: "GITHUB_TOKEN".
	//
	// Experimental.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

