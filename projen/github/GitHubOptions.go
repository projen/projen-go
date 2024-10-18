package github


// Experimental.
type GitHubOptions struct {
	// Download files in LFS in workflows.
	// Default: true if the associated project has `lfsPatterns`, `false` otherwise.
	//
	// Experimental.
	DownloadLfs *bool `field:"optional" json:"downloadLfs" yaml:"downloadLfs"`
	// Whether a merge queue should be used on this repository to merge pull requests.
	//
	// Requires additional configuration of the repositories branch protection rules.
	// Default: true.
	//
	// Experimental.
	MergeQueue *bool `field:"optional" json:"mergeQueue" yaml:"mergeQueue"`
	// Options for MergeQueue.
	// Default: - default options.
	//
	// Experimental.
	MergeQueueOptions *MergeQueueOptions `field:"optional" json:"mergeQueueOptions" yaml:"mergeQueueOptions"`
	// Whether mergify should be enabled on this repository or not.
	// Default: true.
	//
	// Experimental.
	Mergify *bool `field:"optional" json:"mergify" yaml:"mergify"`
	// Options for Mergify.
	// Default: - default options.
	//
	// Experimental.
	MergifyOptions *MergifyOptions `field:"optional" json:"mergifyOptions" yaml:"mergifyOptions"`
	// Choose a method of providing GitHub API access for projen workflows.
	// Default: - use a personal access token named PROJEN_GITHUB_TOKEN.
	//
	// Experimental.
	ProjenCredentials GithubCredentials `field:"optional" json:"projenCredentials" yaml:"projenCredentials"`
	// The name of a secret which includes a GitHub Personal Access Token to be used by projen workflows.
	//
	// This token needs to have the `repo`, `workflows`
	// and `packages` scope.
	// Default: "PROJEN_GITHUB_TOKEN".
	//
	// Deprecated: - use `projenCredentials`.
	ProjenTokenSecret *string `field:"optional" json:"projenTokenSecret" yaml:"projenTokenSecret"`
	// Add a workflow that allows backport of PRs to other branches using labels.
	//
	// When opening a new PR add a backport label to it,
	// and the PR will be backported to the target branches once the PR is merged.
	//
	// Should not be used together with mergify.
	// Default: false.
	//
	// Experimental.
	PullRequestBackport *bool `field:"optional" json:"pullRequestBackport" yaml:"pullRequestBackport"`
	// Options for configuring pull request backport.
	// Default: - see defaults in `PullRequestBackportOptions`.
	//
	// Experimental.
	PullRequestBackportOptions *PullRequestBackportOptions `field:"optional" json:"pullRequestBackportOptions" yaml:"pullRequestBackportOptions"`
	// Add a workflow that performs basic checks for pull requests, like validating that PRs follow Conventional Commits.
	// Default: true.
	//
	// Experimental.
	PullRequestLint *bool `field:"optional" json:"pullRequestLint" yaml:"pullRequestLint"`
	// Options for configuring a pull request linter.
	// Default: - see defaults in `PullRequestLintOptions`.
	//
	// Experimental.
	PullRequestLintOptions *PullRequestLintOptions `field:"optional" json:"pullRequestLintOptions" yaml:"pullRequestLintOptions"`
	// Enables GitHub workflows.
	//
	// If this is set to `false`, workflows will not be created.
	// Default: true.
	//
	// Experimental.
	Workflows *bool `field:"optional" json:"workflows" yaml:"workflows"`
}

