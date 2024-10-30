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
	// Choose a method for authenticating with GitHub to enable auto-queue on pull requests.
	//
	// The workflow cannot use a default github token. Queuing a PR
	// with the default token will not trigger any merge queue workflows,
	// which results in the PR just not getting merged at all.
	// See: https://projen.io/docs/integrations/github/
	//
	// Default: - uses credentials from the GitHub component.
	//
	// Experimental.
	ProjenCredentials GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// Github Runner selection labels.
	// Default: ["ubuntu-latest"].
	//
	// Experimental.
	RunsOn *[]*string `field:"optional" json:"runsOn" yaml:"runsOn"`
}

