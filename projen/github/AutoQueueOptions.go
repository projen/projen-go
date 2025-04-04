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
	// The branch names that we should auto-queue for.
	//
	// This set of branches should be a subset of `MergeQueueOptions.targetBranches`.
	//
	// Be sure not to enable `autoQueue` for branches that don't have branch rules
	// with merge requirements set up, otherwise new PRs will be merged
	// immediately after creating without a chance for review.
	//
	// ## Automatically merging a set of Stacked PRs
	//
	// If you set this to `['main']` you can automatically merge a set of Stacked PRs
	// in the right order. It works like this:
	//
	// - Create PR #1 from branch `a`, targeting `main`.
	// - Create PR #2 from branch `b`, targeting branch `a`.
	// - Create PR #3 from branch `c`, targeting branch `b`.
	//
	// Initially, PR #1 will be set to auto-merge, PRs #2 and #3 will not.
	//
	// Once PR #1 passes all of its requirements it will merge. That will delete
	// branch `a` and change  the target branch of PR #2 change to `main`. At that
	// point, auto-queueing will switch on for PR #2 and it gets merged, etc.
	//
	// > [!IMPORTANT]
	// > This component will never disable AutoMerge, only enable it. So if a PR is
	// > initially targeted at one of the branches in this list, and then
	// > subsequently retargeted to another branch, *AutoMerge is not
	// > automatically turned off*.
	// Experimental.
	TargetBranches *[]*string `field:"optional" json:"targetBranches" yaml:"targetBranches"`
}

